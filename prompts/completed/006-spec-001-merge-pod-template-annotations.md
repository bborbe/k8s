---
status: completed
spec: ["001"]
summary: StatefulSetDeployer now merges pod-template annotations on update instead of replacing them, preserving annotations written by other tools; added 7 Ginkgo test contexts and a CHANGELOG entry, with make precommit exiting 0.
execution_id: k8s-statefulset-annotations-exec-006-spec-001-merge-pod-template-annotations
dark-factory-version: dev
created: "2026-09-10T12:00:00Z"
queued: "2026-09-10T12:52:50Z"
started: "2026-09-10T12:52:51Z"
completed: "2026-09-10T12:56:59Z"
---

# Merge pod template annotations on StatefulSet update

<summary>
- The StatefulSet deployer no longer wipes pod-template annotations written by other tools when it updates a live StatefulSet.
- Annotation keys that exist on the live pod template but are not set by the caller now survive an update.
- The caller still wins on any annotation key it does set, and caller-only keys are still applied.
- Everything else about the pod template (labels, containers, volumes, all other fields) keeps the previous replace-everything behaviour.
- Neither the caller's object nor the fetched live object has its annotation map modified in place — the merge allocates a fresh map.
- A nil annotation map on either side is handled as empty, so no input combination panics.
- Creating a StatefulSet that does not exist yet is completely unchanged.
- No new configuration, flag, or opt-in: every consumer gets the merge on upgrade.
- Consumers are told about the behaviour change through a new CHANGELOG entry.
- Accepted trade-off documented for consumers: dropping an annotation from the built template no longer removes it from the live object.
</summary>

<objective>
Fix the StatefulSet deployer so its update path merges pod-template annotations instead of replacing them wholesale. After the change, an annotation another tool wrote onto the live pod template survives `Deploy`, which stops the endless pod-template-hash churn and pod recycling described in spec 001.
</objective>

<context>
Read these files before making changes:

- `k8s_statefulset-deployer.go` — the single file under change. The bug is the line `existing.Spec.Template = statefulSet.Spec.Template` inside `(*statefulSetDeployer).Deploy`, directly under the comment `// Update path: merge only the mutable spec fields into the live object.`
- `k8s_statefulset-deployer_test.go` — the existing Ginkgo v2 + Gomega suite. New test cases must follow its structure exactly: a top-level `Describe("StatefulSet Deployer")`, a shared `statefulSet appsv1.StatefulSet` built in the outer `BeforeEach`, per-case `Context(...)` blocks that stub the counterfeiter fakes in their own `BeforeEach`, and a `JustBeforeEach` that calls `statefulSetDeployer.Deploy(ctx, statefulSet)`.
- `k8s_service-deployer.go` (function `mergeService`) and `k8s_configmap-deployer.go` (function `mergeConfigMap`) — the project's existing convention for package-level unexported merge helpers placed at the bottom of a deployer file. Follow that shape.
- `docs/dod.md` — the project's Definition of Done (this repo has no root `CLAUDE.md`).
- `CHANGELOG.md` — note there is currently **no** `## Unreleased` heading; the topmost version heading is `## v1.14.19`.
- `README.md` — project overview.

Relevant guides (read inside the container):

- `/home/node/.claude/plugins/marketplaces/coding/docs/go-testing-guide.md` — Ginkgo/Gomega + counterfeiter conventions.
- `/home/node/.claude/plugins/marketplaces/coding/docs/changelog-guide.md` — CHANGELOG entry format.
- `/home/node/.claude/plugins/marketplaces/coding/docs/go-error-wrapping-guide.md` — `github.com/bborbe/errors` usage (no new error paths are introduced here, but existing wrapping must stay intact).

Verbatim current state of the buggy block in `(*statefulSetDeployer).Deploy`:

```go
	// Update path: merge only the mutable spec fields into the live object.
	// Immutable fields (selector, serviceName, volumeClaimTemplates) must not be
	// sent on update — the API server rejects any change to them.
	existing.Spec.Template = statefulSet.Spec.Template
	existing.Spec.Replicas = statefulSet.Spec.Replicas
	existing.Spec.UpdateStrategy = statefulSet.Spec.UpdateStrategy
	existing.Spec.RevisionHistoryLimit = statefulSet.Spec.RevisionHistoryLimit
	existing.Spec.MinReadySeconds = statefulSet.Spec.MinReadySeconds
```

Current signatures that must not change:

```go
func NewStatefulSetDeployer(clientset k8s_kubernetes.Interface) StatefulSetDeployer

type StatefulSetDeployer interface {
	Deploy(ctx context.Context, statefulSet appsv1.StatefulSet) error
	Undeploy(ctx context.Context, namespace Namespace, name Name) error
}
```
</context>

<requirements>

## 1. Implement the annotation merge in `k8s_statefulset-deployer.go`

1. In `(*statefulSetDeployer).Deploy`, on the update path only, capture the live pod template's annotation map **before** the template is overwritten, then reassign the merged map afterwards. Replace this line:

   ```go
   existing.Spec.Template = statefulSet.Spec.Template
   ```

   with:

   ```go
   livePodTemplateAnnotations := existing.Spec.Template.ObjectMeta.Annotations
   existing.Spec.Template = statefulSet.Spec.Template
   existing.Spec.Template.ObjectMeta.Annotations = mergePodTemplateAnnotations(
   	livePodTemplateAnnotations,
   	statefulSet.Spec.Template.ObjectMeta.Annotations,
   )
   ```

   Capturing the live map first is load-bearing: after `existing.Spec.Template = statefulSet.Spec.Template`, `existing.Spec.Template.ObjectMeta.Annotations` aliases the **caller's** map, so reading it at that point would read the wrong side.

2. Add the merge helper as a package-level unexported function at the **bottom** of `k8s_statefulset-deployer.go`, below `Undeploy`, matching the placement convention of `mergeService` in `k8s_service-deployer.go`:

   ```go
   // mergePodTemplateAnnotations returns the union of the live pod template's
   // annotations and the desired ones. Keys present in desired win. A nil map on
   // either side is treated as empty. When either side is non-empty the result is
   // a newly allocated map; when both are empty the desired map is returned
   // unchanged (nil stays nil). Neither input map is ever mutated.
   func mergePodTemplateAnnotations(live map[string]string, desired map[string]string) map[string]string {
   	if len(live) == 0 && len(desired) == 0 {
   		return desired
   	}
   	result := make(map[string]string, len(live)+len(desired))
   	for key, value := range live {
   		result[key] = value
   	}
   	for key, value := range desired {
   		result[key] = value
   	}
   	return result
   }
   ```

   Do NOT use `maps.Copy` or any new import — the function above needs no imports beyond what the file already has.

3. Update the update-path comment block so it documents the new annotation semantics. Replace:

   ```go
   	// Update path: merge only the mutable spec fields into the live object.
   	// Immutable fields (selector, serviceName, volumeClaimTemplates) must not be
   	// sent on update — the API server rejects any change to them.
   ```

   with:

   ```go
   	// Update path: merge only the mutable spec fields into the live object.
   	// Immutable fields (selector, serviceName, volumeClaimTemplates) must not be
   	// sent on update — the API server rejects any change to them.
   	// The pod template is taken from the desired object except for its
   	// annotations, which are merged so keys written by other tools survive.
   	// Consequence: dropping an annotation from the desired template no longer
   	// removes it from the live object.
   ```

4. Do NOT touch the create path. The line `Create(ctx, &statefulSet, metav1.CreateOptions{})` must remain exactly as it is and must appear exactly once in the file.

5. Do NOT touch `Undeploy`, the `StatefulSetDeployer` interface, `NewStatefulSetDeployer`, or the `statefulSetDeployer` struct.

6. Do NOT add any tool-specific annotation key or key prefix anywhere. The merge is key-agnostic; the strings `keel`, `keel.sh`, and any other vendor prefix must not appear in the implementation.

## 2. Add unit tests to `k8s_statefulset-deployer_test.go`

Add the following new `Context` blocks **inside the existing `Describe("Deploy", ...)` block**, after the existing `Context("when statefulSet exists with a different template", ...)` block. Do not modify or delete any existing assertion, `BeforeEach`, or `Context`. All new cases are additive.

Build the live object with `statefulSet.DeepCopy()` (never the shallow `existingStatefulSet := statefulSet` form used by some older cases) so the live and desired annotation maps are genuinely distinct objects — a shallow copy shares the map and would make the mutation test vacuous.

7. **Context "when the live pod template has an annotation the desired template does not set"** — covers acceptance criterion 1 (foreign annotation survives).

   ```go
   Context("when the live pod template has an annotation the desired template does not set", func() {
   	BeforeEach(func() {
   		existingStatefulSet := statefulSet.DeepCopy()
   		existingStatefulSet.ResourceVersion = "123"
   		existingStatefulSet.Spec.Template.ObjectMeta.Annotations = map[string]string{
   			"example.com/marker": "live-value",
   		}
   		statefulSetInterface.GetReturns(existingStatefulSet, nil)
   		statefulSetInterface.UpdateReturns(existingStatefulSet, nil)
   	})

   	It("returns no error", func() {
   		Expect(err).To(BeNil())
   	})

   	It("preserves the annotation only present on the live pod template", func() {
   		Expect(statefulSetInterface.UpdateCallCount()).To(Equal(1))
   		_, updatedStatefulSet, _ := statefulSetInterface.UpdateArgsForCall(0)
   		Expect(updatedStatefulSet.Spec.Template.ObjectMeta.Annotations).To(
   			HaveKeyWithValue("example.com/marker", "live-value"),
   		)
   	})
   })
   ```

8. **Context "when live and desired pod templates set the same annotation key"** — covers acceptance criterion 2 (caller wins on conflict). Live gets `map[string]string{"shared": "old"}`, the desired template gets `statefulSet.Spec.Template.ObjectMeta.Annotations = map[string]string{"shared": "new"}` (set in the same `BeforeEach`, after the `DeepCopy`). Assert the object from `UpdateArgsForCall(0)` has `HaveKeyWithValue("shared", "new")`.

9. **Context "when the desired pod template adds an annotation the live one lacks"** — covers acceptance criterion 3. This case does NOT test the nil-live-map path (that is requirement 10). Give live `map[string]string{"existing": "kept"}` and the desired template `map[string]string{"added": "yes"}` — both non-nil. Assert the updated object has BOTH `HaveKeyWithValue("added", "yes")` and `HaveKeyWithValue("existing", "kept")`.

10. **Context "when the live pod template has a nil annotation map"** — covers acceptance criterion 4, first half. Leave the live template's annotations nil (the outer `BeforeEach` already builds a template without annotations, so a `DeepCopy` has nil there — do not set it), and set the desired template's annotations to `map[string]string{"added": "yes"}`. Assert:
    - `Expect(err).To(BeNil())`
    - `Expect(updatedStatefulSet.Spec.Template.ObjectMeta.Annotations).To(Equal(map[string]string{"added": "yes"}))`

11. **Context "when the desired pod template has a nil annotation map"** — covers acceptance criterion 4, second half. Give live `map[string]string{"example.com/marker": "live-value"}` and leave the desired template's annotations nil. Assert:
    - `Expect(err).To(BeNil())`
    - `Expect(updatedStatefulSet.Spec.Template.ObjectMeta.Annotations).To(Equal(map[string]string{"example.com/marker": "live-value"}))`

12. **Context "when merging pod template annotations"** with an `It("does not mutate the annotations of the object passed in by the caller")` — covers acceptance criterion 5 and the "merge writes into the caller's map" failure mode. Setup: live gets `map[string]string{"example.com/marker": "live-value"}`, the desired template gets `map[string]string{"added": "yes"}` (set the desired template's annotations after the `DeepCopy`, as in requirement 8). Assertions:
    - The updated object carries both keys.
    - `Expect(statefulSet.Spec.Template.ObjectMeta.Annotations).To(Equal(map[string]string{"added": "yes"}))` — the caller's own map is unchanged and must NOT have gained `example.com/marker`.

    Read `statefulSet` (the outer suite variable that was passed to `Deploy`) directly inside the `It`; because `Deploy` takes the struct by value but the map by reference, this assertion is the negative evidence that the merge allocated a new map.

13. **Context "when the live pod template has a label and a container the desired template does not"** — covers acceptance criterion 6 (nothing outside annotations is merged). Setup on the `DeepCopy`d live object:
    - `existingStatefulSet.Spec.Template.ObjectMeta.Labels = map[string]string{"app": "test-app", "live-only": "yes"}`
    - `existingStatefulSet.Spec.Template.Spec.Containers[0].Image = "old-image:1.0"`

    Assertions on `UpdateArgsForCall(0)`:
    - `Expect(updatedStatefulSet.Spec.Template.ObjectMeta.Labels).NotTo(HaveKey("live-only"))` — labels keep replace semantics.
    - `Expect(updatedStatefulSet.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("app", "test-app"))`
    - `Expect(updatedStatefulSet.Spec.Template.Spec.Containers[0].Image).To(Equal("test-image:latest"))` — the container list is fully replaced by the desired value.

14. Every new `Context` must stub both `statefulSetInterface.GetReturns(existingStatefulSet, nil)` and `statefulSetInterface.UpdateReturns(existingStatefulSet, nil)` so `Deploy` takes the update path and returns no error.

15. Do NOT add tests that assert on the source text of the implementation (no assertion about the absence of `existing.Spec.Template = statefulSet.Spec.Template`). The spec deliberately excludes implementation-shape assertions.

## 3. Record the change in `CHANGELOG.md`

16. `CHANGELOG.md` currently has no `## Unreleased` heading. Create one directly above `## v1.14.19`, separated by blank lines, matching the file's existing heading and bullet style:

    ```markdown
    ## Unreleased

    - fix: StatefulSetDeployer merges pod template annotations on update instead of replacing them, so annotations written by other tools are no longer removed. Note: an annotation dropped from the desired template is no longer removed from the live object — remove it directly on the live object instead.

    ## v1.14.19
    ```

17. Do NOT modify any existing released version section in `CHANGELOG.md`.

## 4. Finish

18. Run `make precommit` and fix anything it reports. `make format` (part of `precommit`) reformats Go files to a 100-column limit via golines plus gofmt — let it rewrite line wrapping rather than hand-wrapping to match.

19. If `make precommit` reports any lint finding on the new code, fix it with the minimal code change; do not weaken `.golangci.yml`.

20. Decided semantics for the both-empty case (the spec does not state it): return the desired map unchanged, so nil stays nil. This keeps pre-fix observable behaviour and avoids introducing a nil -> empty-map transition on the live object. Do NOT "simplify" this to always allocate.

21. Before finishing, re-run every command in `<verification>` and confirm each one meets its stated expectation, then walk each of spec 001's nine acceptance criteria against the change and name the test or grep that satisfies it.
</requirements>

<constraints>
- The exported API must not change: `NewStatefulSetDeployer(clientset k8s_kubernetes.Interface) StatefulSetDeployer` keeps its current signature, and the `StatefulSetDeployer` interface keeps its current two methods. Consumers upgrade by bumping the module version only.
- Error wrapping stays on `github.com/bborbe/errors` — no `fmt.Errorf`.
- Tests stay Ginkgo v2 + Gomega, following the existing structure of `k8s_statefulset-deployer_test.go`.
- No pre-existing assertion in `k8s_statefulset-deployer_test.go` may be modified or deleted; new cases are additive.
- The deployer must not gain knowledge of any specific external tool — no tool-specific annotation keys or key prefixes anywhere in the implementation.
- `DeploymentDeployer`, `CronJobDeployer`, `JobDeployer`, `ServiceDeployer`, `IngressDeployer` and `ConfigMapDeployer` are out of this change's reach and must not be edited. The only Go files you may modify are `k8s_statefulset-deployer.go` and `k8s_statefulset-deployer_test.go`.
- The merge is unconditional: do NOT add a config field, option, opt-in flag, functional option, or constructor parameter to enable or disable it.
- Only pod-template **annotations** gain merge semantics. Labels and every other pod-template field keep replace semantics.
- Do NOT regenerate or hand-edit anything in `mocks/` — the interface is unchanged. (`make precommit` runs `make generate`; that is expected and fine.)
- Do NOT commit — dark-factory handles git.
- Existing tests must still pass.
</constraints>

<verification>
Run `make precommit` — must exit 0 (format + generate + test + lint + security checks).

Then run each of these and check the stated expectation:

- `grep -nE 'Template\.(ObjectMeta\.)?Annotations' k8s_statefulset-deployer.go` — must return at least 1 line, proving the annotation handling landed in the deployer and not only in the tests.
- `grep -c 'Create(ctx, &statefulSet, metav1.CreateOptions{})' k8s_statefulset-deployer.go` — must print exactly `1`, proving the create path is intact.
- `grep -n 'func mergePodTemplateAnnotations' k8s_statefulset-deployer.go` — must return 1 line.
- `grep -A5 '## Unreleased' CHANGELOG.md` — must show the new fix bullet.
- `! grep -qiE 'keel' k8s_statefulset-deployer.go` — must exit 0 (no tool-specific knowledge in the implementation). Written as `! grep -q` rather than a bare `grep`, because a bare `grep` exits 1 when it finds nothing, i.e. exactly when this assertion passes.
- `go test -mod=mod -run TestSuite -count=1 ./...` — must pass (this is the same suite `make test` runs; use it for a fast re-check while iterating).

Note: the runtime replay against the live cluster (annotation survives ≥3 deployer cycles) is an operator-side step recorded in the spec's Verification section — it cannot run here and is not part of this prompt.
</verification>
