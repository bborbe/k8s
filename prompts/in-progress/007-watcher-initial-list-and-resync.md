---
status: failed
execution_id: k8s-watcher-resync-exec-007-watcher-initial-list-and-resync
dark-factory-version: v0.198.17
created: "2026-11-12T12:51:00Z"
queued: "2026-11-12T12:51:00Z"
started: "2026-11-12T13:03:27Z"
completed: "2026-11-12T13:03:29Z"
lastFailReason: 'execute prompt: docker run failed: wait command: exit status 1'
---

# Deliver existing objects on startup and resync watchers periodically

<summary>
- Both the Secret and the Pod watcher now deliver every object that already exists when they start, instead of only objects that change after they start.
- A periodic full resync re-delivers the current object set on a fixed interval, so a watch event that was missed or dropped repairs itself without a restart.
- The interval defaults to five minutes and can be set explicitly through a new constructor, so it is testable.
- Callers keep their existing call sites unchanged: the current constructors and the watcher interfaces are not altered.
- The existing return behaviour is preserved: the watcher still returns when the watch channel closes, still returns the event processor's error, and still returns the same error on a closed channel, so the retry wrapper reconnects as it does today — with the startup list, each reconnect re-delivers the full set.
- An error during the startup list is wrapped and returned, so the existing retry wrapper turns it into a retry.
- An interval that is zero or negative cannot panic the watcher: the startup list still runs and the periodic resync is skipped.
- The public doc comments are updated, because they currently describe the old contract.
- Consumers are told about the behaviour change through a new CHANGELOG entry.
</summary>

<objective>
Make `SecretWatcher` and `PodWatcher` reconcile from a full list at startup and on a fixed interval, so a consumer that starts after its objects already exist still receives them, and a dropped watch event repairs itself. Without this, a process restarted against existing objects copies or observes nothing until each object changes again.
</objective>

<context>
Read these files before making changes:

- `k8s_secret-watcher.go` — the `SecretWatcher` interface, `NewSecretWatcher`, and `secretStatusMonitoring`. Its `Watch` calls `Secrets(ns).Watch(ctx, metav1.ListOptions{})` with no initial list and no resync. The interface (`Watch(ctx context.Context) error`), `ErrResultChannelClosed`, and every existing return path must keep working.
- `k8s_pod-watcher.go` — the identical shape for `PodWatcher`, `NewPodWatcher`, `podStatusMonitoring`.
- `k8s_secret-watcher-retry.go` — the in-repo precedent for injecting timing: it takes `libtime.WaiterDuration` and `libtime.Duration` as constructor parameters and wraps errors with `errors.Wrapf`. Follow that signature style for the resync interval.
- `k8s_secret-watcher_test.go` and `k8s_pod-watcher_test.go` — the existing Ginkgo v2 + Gomega suites. The outer `BeforeEach` stubs `clientset.CoreV1Returns(coreV1)`, `coreV1.SecretsReturns(secretInterface)`, and `secretInterface.WatchReturns(fakeWatcher, nil)`; there is no `List` stub today. The suites call `Watch(ctx)` synchronously in a shared `JustBeforeEach` and rely on it returning — that must keep working.
- `docs/dod.md` — the project's Definition of Done. This repo has a root `CLAUDE.md` that forbids direct coding; this prompt is the sanctioned route.
- `CHANGELOG.md` — the topmost heading is `## v1.14.33`; there is currently no `## Unreleased` heading.

Relevant guides (read inside the container):

- `/home/node/.claude/plugins/marketplaces/coding/docs/go-context-cancellation-in-loops.md` — a non-blocking `ctx.Done()` check at the top of every loop that can outlive a few milliseconds, including the loop over a list result.
- `/home/node/.claude/plugins/marketplaces/coding/docs/go-testing-guide.md` — Ginkgo/Gomega conventions and counterfeiter mocks.
- `/home/node/.claude/plugins/marketplaces/coding/docs/go-time-injection.md` — the interval is injected, never a package constant read inside the loop body.

Current signatures that must not change:

```go
type SecretWatcher interface {
	Watch(ctx context.Context) error
}

func NewSecretWatcher(clientset kubernetes.Interface, secretManager SecretEventProcessor, namespace Namespace) SecretWatcher

type PodWatcher interface {
	Watch(ctx context.Context) error
}

func NewPodWatcher(clientset kubernetes.Interface, podManager PodEventProcessor, namespace Namespace) PodWatcher
```

The two event-processor interfaces consumers already implement:

```go
type SecretEventProcessor interface {
	OnUpdate(ctx context.Context, secret corev1.Secret) error
	OnDelete(ctx context.Context, secret corev1.Secret) error
}

type PodEventProcessor interface {
	OnUpdate(ctx context.Context, pod corev1.Pod) error
	OnDelete(ctx context.Context, pod corev1.Pod) error
}
```

Design decision, already settled — do not substitute a shared informer. A `SharedInformerFactory` would supply the initial list and the resync, but it makes `Watch` block until the context is cancelled and swallows event-processor errors (an informer handler has no error return). That would break this library's exported contract, its retry wrapper, and three existing test cases. Keep the hand-rolled watch loop and add the list and the ticker to it.
</context>

<requirements>

## 1. Add the startup list and the periodic resync to the Secret watcher

1. In `k8s_secret-watcher.go`, keep the `SecretWatcher` interface, `ErrResultChannelClosed`, `NewSecretWatcher`, and `secretStatusMonitoring` names. Declare the default interval once:

   ```go
   // defaultResyncInterval is the periodic full-resync interval applied by
   // NewSecretWatcher and NewPodWatcher when no explicit interval is given.
   const defaultResyncInterval = libtime.Duration(5 * time.Minute)
   ```

2. Add a second constructor and make `NewSecretWatcher` delegate to it. Do not change `NewSecretWatcher`'s parameter list — consumers call it as-is.

   ```go
   // NewSecretWatcherWithResync is NewSecretWatcher with an explicit resync interval.
   // A non-positive interval disables the periodic resync; the startup list still runs.
   func NewSecretWatcherWithResync(clientset kubernetes.Interface, secretManager SecretEventProcessor, namespace Namespace, resyncInterval libtime.Duration) SecretWatcher
   ```

   Add `resyncInterval libtime.Duration` to the `secretStatusMonitoring` struct and set it from both constructors. The non-positive branch exists so `time.NewTicker` is never constructed with a zero or negative duration — `time.NewTicker` panics on those. No consumer is expected to use it today; it is a guard, not a feature.

3. Add an unexported method on `*secretStatusMonitoring` that performs one full reconcile:

   ```go
   func (s *secretStatusMonitoring) reconcile(ctx context.Context) error
   ```

   It calls `s.clientset.CoreV1().Secrets(s.namespace.String()).List(ctx, metav1.ListOptions{})`, wraps a non-nil error with `errors.Wrap(ctx, err, "list secrets failed")` and returns it, then iterates the returned `SecretList.Items` and calls `s.secretManager.OnUpdate(ctx, item)` for each, wrapping a non-nil error with `errors.Wrap(ctx, err, "on update failed")`. Put a non-blocking `ctx.Done()` check at the top of the item loop and return `ctx.Err()` if the context has ended.

   Do not pass a resource version — an empty `metav1.ListOptions{}` is a true full list, which is what a resync needs.

4. Rewrite `(*secretStatusMonitoring).Watch` as:

   1. the existing non-blocking `ctx.Done()` pre-check first, returning `ctx.Err()` with no API call — this keeps `WatchCallCount()` at `0` for the already-cancelled case;
   2. `s.reconcile(ctx)` — the startup list; return its error directly;
   3. the existing `Secrets(ns).Watch(ctx, metav1.ListOptions{})` call, unchanged, including its `defer watchInf.Stop()` and its `errors.Wrap(ctx, err, "watch failed")`;
   4. a ticker built only when `s.resyncInterval > 0`, with `defer ticker.Stop()` and `time.NewTicker(s.resyncInterval.Duration())` — `libtime.Duration` is a defined type over `time.Duration`, so convert with `.Duration()` rather than casting; when the interval is non-positive, use a nil channel so the `select` simply never takes that branch;
   5. the existing `select` over `ctx.Done()`, `resultChan`, and now the ticker channel. Every existing branch keeps its current behaviour and its current error text. On a tick, call `s.reconcile(ctx)` and return its error on failure.

   Do NOT use `time.Tick` or `time.After` in the loop.

5. Update the two doc comments. `NewSecretWatcher`'s currently says "This watcher returns when the watch connection closes. Use NewSecretWatcherRetry for automatic retries." — after this change it must say that the watcher delivers the objects that already exist at startup, re-delivers the full set every `defaultResyncInterval`, and still returns when the watch channel closes so the retry wrapper reconnects — and that delivery is **at-least-once**, because each reconnect re-delivers the full set. Same for `NewPodWatcher`.

6. Keep the existing `glog.V(2)` lifecycle lines and add one that names the effective resync interval, so an operator can read it from the log.

7. Do NOT delete `ErrResultChannelClosed` — it is exported, it is declared in `k8s_errors.go`, and the closed-channel branch still returns it.

## 2. Apply the identical change to the Pod watcher

8. In `k8s_pod-watcher.go`, mirror every step above: `NewPodWatcherWithResync`, the same `defaultResyncInterval` (declared once, in `k8s_secret-watcher.go`), a `reconcile` method using `Pods(...).List(...)` and `PodEventProcessor`, and the same ticker branch. Mirror all six new test cases from requirement 10 in `k8s_pod-watcher_test.go`, using the same case names with `secret` → `pod` (`existing-pod`, `podInterface.ListReturns(&corev1.PodList{}, nil)`, `podEventProcessor`, `podWatcher = k8s.NewPodWatcherWithResync(...)`, and the non-positive case). The pod List error must be wrapped with `errors.Wrap(ctx, err, "list pods failed")` — do not copy the secret wording. This is what covers `podStatusMonitoring.reconcile` and the pod ticker branch — `docs/dod.md` requires new functions to have tests. The two implementation files are deliberately near-identical, matching their current shape — do not extract a shared generic helper.

## 3. Update the existing test suites

9. In both `k8s_secret-watcher_test.go` and `k8s_pod-watcher_test.go`, add a `List` stub to the outer `BeforeEach` so the startup list is deterministic:

   ```go
   secretInterface.ListReturns(&corev1.SecretList{}, nil)   // pod suite: podInterface.ListReturns(&corev1.PodList{}, nil)
   ```

   Do not modify or delete any existing assertion, `Context`, or `BeforeEach` other than adding this line. Every existing case must still pass — including `Watch` returning `ErrResultChannelClosed` on a closed channel and returning the processor's error. If a case fails, the implementation has broken the contract; fix the implementation, not the test.

10. Add these new `Context` blocks inside the existing `Describe("Watch", ...)`, after the last existing case. They all run through the shared `JustBeforeEach`, which calls `Watch` synchronously and returns only when the watch channel closes — so every case whose `Watch` reaches the watch loop (not just the early-return error cases) must stop the fake watcher from its own `BeforeEach`, or the spec hangs. `Stop()` only closes the channel; it sends no event.

    - **"when objects already exist before the watch starts"** — `List` returns a `SecretList` with one item named `existing-secret`. In the case's `BeforeEach`, start a goroutine that stops the fake watcher so the shared `JustBeforeEach`'s `Watch` call can return. Assert `OnUpdateCallCount()` is `1` and the received secret's name is `existing-secret`, with no watch event sent.
    - **"when the initial list fails"** — `ListReturns(nil, errors.New("list failed"))`. Assert `Watch` returns an error whose message contains `list secrets failed`, and `WatchCallCount()` is `0` (the watch never starts). No stop is needed — this case returns before the watch loop.
    - **"when the event processor fails during the initial list"** — `List` returns one item and `OnUpdateReturns(errors.New("update failed"))`. Assert `Watch` returns an error whose message contains `on update failed`. No stop is needed — this case returns before the watch loop.
    - **"when the resync interval elapses"** — in the case's `BeforeEach`, assign `secretWatcher = k8s.NewSecretWatcherWithResync(clientset, secretEventProcessor, namespace, libtime.Duration(10*time.Millisecond))` (import `libtime "github.com/bborbe/time"` and `"time"`); the assignment is required, because the shared `JustBeforeEach` calls `Watch` on the outer `secretWatcher` variable. `List` returns one item. Do NOT call `Watch` yourself — the shared `JustBeforeEach` already does. In the case's `BeforeEach`, start a goroutine that waits roughly 200 ms and then stops the fake watcher, so the shared `Watch` call returns after several ticks. Assert `OnUpdateCallCount()` is at least `2` — the startup list plus at least one resync. This is the case that proves the interval reaches the ticker rather than being silently ignored.
    - **"when the resync interval is non-positive"** — assign `secretWatcher = k8s.NewSecretWatcherWithResync(clientset, secretEventProcessor, namespace, libtime.Duration(0))`, `List` returns one item. In the case's `BeforeEach`, start a goroutine that stops the fake watcher so the shared `JustBeforeEach`'s `Watch` call can return. Assert `OnUpdateCallCount()` is exactly `1` — the startup list ran and no resync followed.
    - **"when the context is canceled during the initial list"** — `List` returns two items, and `OnUpdateStub` calls `cancel()` and returns `nil`. Assert `Watch` returns `context.Canceled` and `OnUpdateCallCount()` is `1` — the item loop's `ctx.Done()` check stops the second item. No stop is needed — this case returns before the watch loop.

    Assert exact counts only where the count is deterministic; use `BeNumerically(">=", n)` where timing could add an extra pass.

11. Keep every new test deterministic: no `time.Sleep` inside an assertion, no wall-clock reads. Where the fake watcher must be stopped after a delay, do it in a goroutine and let the assertion run after `Watch` returns.

## 4. Record the change in `CHANGELOG.md`

12. Add an `## Unreleased` section directly above `## v1.14.33`, matching the file's existing heading and bullet style, with a single bullet stating that both watchers now deliver pre-existing objects on startup and resync periodically at an injected interval (default five minutes), that delivery is at-least-once (a reconnect re-delivers the full set), and that the interfaces and existing constructor signatures are unchanged.

13. Do NOT modify any existing released version section.

## 5. Finish

14. Run `ROOTDIR=/workspace make precommit` and fix anything it reports. This prompt runs from a worktree with `.git` masked (`hideGit: true`), so `ROOTDIR` is empty unless passed explicitly and `make` dies before any target runs — always pass `ROOTDIR=/workspace` on every `make` call. `make format` reformats Go files to a 100-column limit via golines plus gofmt — let it rewrite line wrapping rather than hand-wrapping to match.

15. If `ROOTDIR=/workspace make precommit` reports a lint finding on the new code, fix it with the minimal code change; do not weaken `.golangci.yml`. One exception is pre-authorised: if `dupl` fires on the two mirrored watcher bodies, add `//nolint:dupl` with a one-line reason stating that the two files are intentionally parallel — `docs/dod.md` permits `//nolint` with an explanation, and extracting a shared generic helper is explicitly forbidden here.

16. Before finishing, re-run every command in `<verification>` and confirm each meets its stated expectation, then walk each requirement above against the change and name the test or grep that satisfies it.
</requirements>

<constraints>
- The exported API must stay source-compatible: `NewSecretWatcher` and `NewPodWatcher` keep their exact parameter lists, and the `SecretWatcher` / `PodWatcher` interfaces keep their single `Watch(ctx context.Context) error` method. Consumers upgrade by bumping the module version only.
- The existing return contract is preserved: `Watch` returns when the watch channel closes (with `ErrResultChannelClosed`), returns the event processor's wrapped error, and returns `ctx.Err()` on cancellation. `NewSecretWatcherRetry` / `NewPodWatcherRetry` keep working unchanged.
- Error wrapping stays on `github.com/bborbe/errors` — no `fmt.Errorf`.
- The resync interval is injected through the constructor. Do NOT read a package-level constant inside `Watch`, and do NOT call `time.Now()` in the implementation.
- Do NOT switch to a `SharedInformerFactory`, an informer cache, or a `Lister`. The design decision in `<context>` explains why; the exported contract is an event-processor callback, not a cache.
- Do NOT use `time.Tick` or `time.After` in a loop.
- Tests stay Ginkgo v2 + Gomega, following the existing structure. No `t.Run` tables and no hand-written fakes — use the counterfeiter mocks in `mocks/`.
- Do NOT regenerate or hand-edit anything in `mocks/` — the interfaces are unchanged. (`ROOTDIR=/workspace make precommit` runs `make generate`; that is expected and fine.)
- Only these files may be modified: `k8s_secret-watcher.go`, `k8s_pod-watcher.go`, `k8s_secret-watcher_test.go`, `k8s_pod-watcher_test.go`, `CHANGELOG.md`. `k8s_service-watcher.go` has the same startup gap but is out of scope for this change.
- Do NOT commit — dark-factory handles git.
- Existing tests must still pass.
</constraints>

<verification>
Run `ROOTDIR=/workspace make precommit` — must exit 0 (format + generate + test + lint + security checks). `ROOTDIR=/workspace` is mandatory here: `.git` is masked in this run, so `make` cannot derive the repo root on its own.

Then run each of these and check the stated expectation:

- `go test -mod=mod -run TestSuite -count=1 ./...` — must pass (the same suite `make test` runs; use it for a fast re-check while iterating).
- `grep -c 'NewSecretWatcherWithResync\|NewPodWatcherWithResync' k8s_secret-watcher.go k8s_pod-watcher.go` — **both** counts must be at least `2` (a doc comment in one file must not be able to carry the total alone).
- `grep -q 'const defaultResyncInterval' k8s_secret-watcher.go` — must exit 0.
- `grep -c 'func (s \*secretStatusMonitoring) reconcile' k8s_secret-watcher.go` — must print `1`, and the same for `podStatusMonitoring` in `k8s_pod-watcher.go`.
- `grep -c 'List(ctx' k8s_secret-watcher.go k8s_pod-watcher.go` — both counts must be at least `1`, proving both watchers perform the startup list. Use this exact pattern: the call chain breaks after `Secrets(...)` / `Pods(...)`, so `List` sits at the start of its own line with **no** leading dot, and `golines` may wrap the argument list — both `\.List(` and `List(ctx, metav1.ListOptions{})` return `0` against correct code.
- `grep -c 'time.NewTicker' k8s_secret-watcher.go k8s_pod-watcher.go` — both counts must be at least `1`, proving the ticker branch exists in both.
- `! grep -q 'time\.Tick(' k8s_secret-watcher.go k8s_pod-watcher.go` — must exit 0 (no per-iteration timer). Written as `! grep -q` rather than a bare `grep`, because a bare `grep` exits 1 when it finds nothing, i.e. exactly when this assertion passes. Note this and the next check also pass on an untouched repo — they are regression guards against a rejected approach, not change-detectors; the `time.NewTicker` check above is the change-detector.
- `! grep -q 'SharedInformerFactory' k8s_secret-watcher.go k8s_pod-watcher.go` — must exit 0, proving the informer route was not substituted.
- `grep -n '^## Unreleased' CHANGELOG.md` — must return a line number **smaller** than the one `grep -n '^## v1.14.33' CHANGELOG.md` returns, proving the new section sits above the released one rather than merely existing.
- `grep -c 'ListReturns' k8s_secret-watcher_test.go k8s_pod-watcher_test.go` — both counts must be at least `1`, proving both suites stub the startup list.
- `grep -c 'NewSecretWatcherWithResync' k8s_secret-watcher_test.go` — must be at least `1`, proving the positive-interval resync case exists.
- `grep -c 'NewPodWatcherWithResync' k8s_pod-watcher_test.go` — must be at least `1`, proving the pod resync case exists.
- `grep -c 'existing-pod\|list pods failed\|non-positive' k8s_pod-watcher_test.go` — must be at least `3`, proving the pod suite carries the mirrored cases rather than only the resync case.
</verification>
