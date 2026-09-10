---
status: prompted
approved: "2026-09-10T11:54:03Z"
generating: "2026-09-10T12:11:55Z"
prompted: "2026-09-10T12:11:55Z"
branch: dark-factory/bug-statefulset-deployer-strips-foreign-pod-template-annotations
---

## Summary

- The StatefulSet deployer replaces the whole pod template on every update instead of merging its annotations.
- Annotations another tool wrote onto the live pod template are silently deleted, which changes the pod-template hash and rolls every pod.
- The other tool re-adds its annotation and the deployer deletes it again, so the two fight indefinitely and the workload never settles.
- Beyond the wasted rollouts, the recycling resets container restart counters, so restart-based health checks report healthy workloads that are actually being recreated.
- After the fix, pod-template annotations the deployer does not set survive an update untouched.

## Problem

`statefulSetDeployer.Deploy` documents its update path as *"merge only the mutable spec fields into the live object"* and then assigns `Spec.Template` wholesale from the caller's freshly built object.

Callers build that template from scratch and have no view of the live object, so any annotation added to the live pod template by a different writer is lost on the next reconcile. Because pod-template annotations participate in the StatefulSet's revision hash, each loss creates a new ControllerRevision and rolls every pod. When the other writer is a controller that re-applies its annotation on a poll loop, the two overwrite each other forever.

The other fields on that code path are scalars or pointers the caller fully owns (`Replicas`, `UpdateStrategy`, `RevisionHistoryLimit`, `MinReadySeconds`, and the nil-guarded `PersistentVolumeClaimRetentionPolicy` / `Ordinals`). The pod template is different in kind: it is a composite whose annotation map has other legitimate writers, so replacing it wholesale destroys state the caller never owned.

This was found in production: a per-account tick StatefulSet reached generation 7 with 7 ControllerRevisions in 4 days 18 hours without the object ever being recreated, and the diff between consecutive revisions was a single removed annotation.

## Goal

`Deploy` merges the pod template's annotations instead of replacing them: annotation keys present on the live pod template but absent from the caller's template survive the update. The deployer keeps full authority over every key the caller does set and over every other part of the template. A second writer's pod-template annotation is no longer a cause of pod churn.

**Accepted consequence (library-wide, all consumers):** omitting a pod-template annotation from the built template no longer removes it from the live object — such a key becomes sticky. Removal becomes an explicit operator action (edit or recreate the live object). This is the deliberate price of the fix and applies to every consumer on upgrade; it is accepted because these callers set only static pod-template annotations, and because the alternative — per-consumer opt-in configuration — pushes the same defect onto every consumer that forgets to opt in.

## Acceptance Criteria

- [ ] An annotation present only on the live pod template survives `Deploy` — evidence: a unit test builds a live StatefulSet whose `Spec.Template.ObjectMeta.Annotations` contains `example.com/marker: "live-value"`, calls `Deploy` with a template that does not set that key, and asserts the object passed to `UpdateArgsForCall(0)` still carries `example.com/marker: "live-value"`.
- [ ] The caller wins on a conflicting key — evidence: a unit test where live has `shared: "old"` and the built template sets `shared: "new"` asserts the updated object carries `shared: "new"`.
- [ ] Caller-only annotations are still applied — evidence: a unit test where the built template sets `added: "yes"` and live has no such key asserts the updated object carries `added: "yes"`.
- [ ] Merging works when either side has a nil annotation map — evidence: two unit tests. With `live.Spec.Template.ObjectMeta.Annotations == nil` and the built template setting `added: "yes"`, `Deploy` returns no error and the updated object's pod-template annotations are exactly `{added: "yes"}`. With the built template's map nil and live carrying `example.com/marker: "live-value"`, `Deploy` returns no error and the updated object's pod-template annotations are exactly `{example.com/marker: "live-value"}`.
- [ ] The caller's own object is not mutated — evidence: a unit test asserts that after `Deploy` returns, the `statefulSet` value the caller passed in still does NOT contain `example.com/marker` in `Spec.Template.ObjectMeta.Annotations` (negative evidence: the merge builds a new map rather than writing into the caller's).
- [ ] Nothing outside the pod template's annotations is merged — evidence: a unit test asserts that a live pod-template **label** absent from the built template is NOT preserved (labels keep replace semantics), and that live `Spec.Template.Spec.Containers` is fully replaced by the built value.
- [ ] The create path is unchanged — evidence: `grep -n 'Create(ctx, &statefulSet, metav1.CreateOptions{})' k8s_statefulset-deployer.go` returns exactly 1 line, and the existing create-path tests pass unmodified.
- [ ] Existing behaviour is not regressed — evidence: `make precommit` exits 0 with no edits to any pre-existing assertion in `k8s_statefulset-deployer_test.go`.
- [ ] The change is recorded for consumers — evidence: `grep -A5 '## Unreleased' CHANGELOG.md` shows a bullet describing the pod-template annotation merge (the heading does not currently exist and must be created above `## v1.14.19`).

## Verification

### Container-executable (runs inside the YOLO container at prompt time)

- `make precommit` — exits 0 (format + lint + test + security checks)
- `grep -nE 'Template\.(ObjectMeta\.)?Annotations' k8s_statefulset-deployer.go` — returns ≥1 line, proving the annotation handling landed in the deployer rather than only in tests
- `grep -n 'Create(ctx, &statefulSet, metav1.CreateOptions{})' k8s_statefulset-deployer.go` — returns exactly 1 line, proving the create path is intact
- `grep -A5 '## Unreleased' CHANGELOG.md` — returns the new bullet

The unit tests in the acceptance criteria are what prove the wholesale-replace semantics are gone. Deliberately NOT asserted: the absence of the literal line `existing.Spec.Template = statefulSet.Spec.Template`. Copying the template and then overwriting its annotation map is a legitimate implementation shape, so a negative grep on that line would fail a correct fix and turn a behavioral spec into an implementation-shape mandate.

### Operator-executable (runs on the host, after release + consumer bump)

This is a library, so the runtime replay cannot happen at prompt time — it requires a tag, a consumer `go.mod` bump, and a redeploy. Replay owner: the operator, tracked on the vault task *Stop Keel and mt5-tick-controller Fighting Over Per-Account StatefulSet Pod Templates*.

**Completion gate:** every acceptance criterion above is a unit test or a grep, so they can all pass while the bug is still live in production. Per `bug-workflow.md` § Verification ("tests passing ≠ bug fixed"), this spec must NOT be marked `completed` on the container rung alone — the two probes below must be replayed and observed first.

- With the consumer running a build that includes this fix, and `keel.sh/update-time` present on the live pod template of `mt5-tick-requester-ftmo-1514532957`:
  - `kubectlnukedev -n dev get statefulset mt5-tick-requester-ftmo-1514532957 -o jsonpath='{.spec.template.metadata.annotations.keel\.sh/update-time}'` returns the same timestamp after ≥3 deployer cycles (≥45 min at the observed 15-min interval). Pre-fix this value disappears within one cycle, so a no-op cannot satisfy it.
  - `kubectlnukedev -n dev get statefulset mt5-tick-requester-ftmo-1514532957 -o jsonpath='{.metadata.generation}'` is unchanged across that same window, and pod age exceeds 45 min at the final check.

## Desired Behavior

1. On the update path, the deployer computes the pod template's annotations as the union of the live template's annotations and the caller's, with the caller's value winning on any key present in both.
2. Every other field of `Spec.Template` continues to be taken from the caller's object exactly as before, including labels, so only annotations gain merge semantics.
3. The merge writes into a newly allocated map, treating a nil map on either side as empty — neither the caller's object nor the fetched live object's original map is mutated in place, and no input combination panics.
4. The create path is untouched — a StatefulSet that does not yet exist is created verbatim from the caller's object.
5. The merge is unconditional and requires no configuration: every consumer of the library gets it on upgrade.

## Constraints

- The exported API must not change: `NewStatefulSetDeployer(clientset k8s_kubernetes.Interface) StatefulSetDeployer` keeps its current signature, and the `StatefulSetDeployer` interface keeps its current two methods. Consumers upgrade by bumping the module version only.
- Error wrapping stays on `github.com/bborbe/errors` — no `fmt.Errorf`.
- Tests stay Ginkgo v2 + Gomega, following the existing structure of `k8s_statefulset-deployer_test.go`.
- No pre-existing assertion in `k8s_statefulset-deployer_test.go` may be modified or deleted; new cases are additive.
- The deployer must not gain knowledge of any specific external tool — no tool-specific annotation keys or key prefixes anywhere in the implementation.
- `DeploymentDeployer`, `CronJobDeployer`, `JobDeployer`, `ServiceDeployer`, `IngressDeployer` and `ConfigMapDeployer` are out of this change's reach and must not be edited.

## Failure Modes

| Trigger | Expected behavior | Recovery | Detection |
|---|---|---|---|
| Live pod template has a nil annotation map | Merge treats it as empty; caller's annotations applied as-is | None needed | Unit test with nil live map passes |
| Caller's pod template has a nil annotation map | Live annotations are preserved wholesale; no panic | None needed | Unit test with nil built map passes |
| Caller intends to REMOVE a pod-template annotation it previously set | Removal no longer takes effect — the key is sticky and survives | Operator deletes the annotation directly on the live object, or deletes and recreates the StatefulSet | Annotation still present after a deploy that dropped it from the builder |
| Merge writes into the caller's map instead of a new one | Caller's rebuilt template accumulates live keys across reconciles, corrupting its own state | Fix the merge to allocate; no runtime recovery | Acceptance criterion asserting the input object is unmutated fails |
| Two writers set the SAME annotation key with different values | Caller wins every reconcile; the other writer's value is overwritten as before | Change one writer to use a distinct key | Churn continues on that key despite the fix |
| Live object fetched, then modified by another writer before Update | `Update` fails with a conflict error, wrapped and returned as today | Caller's next reconcile cycle retries | `update statefulSet failed` in the returned error |

## Reproduction

Observed on `mt5-tick-requester-ftmo-1514532957` (namespace `dev`), consuming `github.com/bborbe/k8s v1.14.17`.

Setup — a StatefulSet whose pod template carries an annotation written by a second tool, and a controller that calls `Deploy` on a fixed interval (15 min in the observed case) with a template rebuilt from scratch.

Observed evidence, verbatim:

- The StatefulSet is at `generation: 7` with 7 ControllerRevisions spanning 4 days 18 hours, while the StatefulSet object itself was never recreated.
- Annotation presence alternates across revisions — revision 4 carries `keel.sh/update-time: 2026-09-06 14:57:23 UTC`, revision 6 carries `keel.sh/update-time: 2026-09-09 09:21:23 UTC`, revisions 1-3 and 7 carry none.
- A diff of the revision 6 and revision 7 ControllerRevision objects shows exactly one changed line: the removal of the `keel.sh/update-time` annotation. Nothing else in the template differs.
- On the live object, `.metadata.annotations` still carries `kubernetes.io/change-cause: "keel automated update version dev -> dev [2026-09-09T09:21:23Z]"` while `.spec.template.metadata.annotations` carries only the four `prometheus.io/*` keys — the second writer's annotation is already gone.

The `change-cause` timestamp matches ControllerRevision 6 exactly, establishing the order: the second tool wrote the annotation at 09:21:23, the deployer stripped it on a later reconcile, producing revision 7 and a full pod roll.

## Expected vs Actual

**Expected** — per the method's own contract comment in `k8s_statefulset-deployer.go` (*"Update path: merge only the mutable spec fields into the live object"*), state on the live object that the caller does not set is left alone.

**Actual** — `existing.Spec.Template = statefulSet.Spec.Template` replaces the entire pod template, so every annotation the caller does not set is deleted, regardless of who wrote it.

## Why this is a bug

The file states the merge contract in a comment directly above the offending line. The other assignments on that path are scalars and pointers the caller genuinely owns, so replacing them is correct; the pod template is a composite whose annotation map is a documented extension point for other tools, and replacing it wholesale destroys state the caller never owned. The user-visible consequence — unbounded pod churn whenever a second tool annotates the pod template — is a defect by any reading.

## Workaround

Stop the second writer from targeting the pod template (for example, move its annotation to the object's own `.metadata.annotations`, which the deployer does not touch). This avoids the churn but gives up whatever behaviour required the annotation to sit on the pod template.

## Do-Nothing Option

Every consumer of this library that shares a StatefulSet's pod template with another writer keeps rolling its pods indefinitely, and the restart-counter reset keeps hiding real instability from restart-based health checks — a monitoring blind spot that has already invalidated one multi-day stability measurement.
