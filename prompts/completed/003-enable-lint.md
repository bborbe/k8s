---
status: completed
summary: Enabled golangci-lint in the check target and fixed all lint violations including depguard, dupl, prealloc, revive, and staticcheck issues across multiple files.
container: k8s-003-enable-lint
dark-factory-version: v0.59.5-dirty
created: "2026-03-21T12:00:00Z"
queued: "2026-03-21T10:12:35Z"
started: "2026-03-21T10:19:25Z"
completed: "2026-03-21T10:27:29Z"
---

<summary>
- The `check` Makefile target includes `lint` as a dependency, running golangci-lint on every precommit
- The TODO comment about enabling lint is removed from the Makefile
- All golangci-lint violations are fixed so that `make lint` passes cleanly
- The `.golangci.yml` config is updated to match the standard v2 config used across the ecosystem
- The full `make precommit` pipeline passes end-to-end
</summary>

<objective>
Enable the golangci-lint linter in the k8s project's `check` target and fix all lint violations so the project passes `make precommit` with lint enabled. The linter has been configured but deliberately excluded from the check pipeline — time to activate it and bring the codebase into compliance.
</objective>

<context>
Read CLAUDE.md for project conventions and build commands.

Read these files before making changes:
- `Makefile` — the check target to modify
- `.golangci.yml` — the current lint config (needs updates)

Reference standard `.golangci.yml` config to match:
```yaml
version: "2"

run:
  timeout: 5m
  tests: true

linters:
  enable:
    - govet
    - errcheck
    - staticcheck
    - unused
    - revive
    - gosec
    - gocyclo
    - depguard
    - dupl
    - nestif
    - errname
    - unparam
    - bodyclose
    - forcetypeassert
    - asasalint
    - prealloc
  settings:
    depguard:
      rules:
        Main:
          deny:
            - pkg: "github.com/pkg/errors"
              desc: "use github.com/bborbe/errors instead"
            - pkg: "github.com/bborbe/argument"
              desc: "use github.com/bborbe/argument/v2 instead"
            - pkg: "golang.org/x/net/context"
              desc: "use context from standard library instead"
            - pkg: "golang.org/x/lint/golint"
              desc: "deprecated, use revive or staticcheck instead"
            - pkg: "io/ioutil"
              desc: "deprecated since Go 1.16, use io and os packages instead"
    funlen:
      lines: 80
      statements: 50
    gocognit:
      min-complexity: 20
    nestif:
      min-complexity: 4
    maintidx:
      min-maintainability-index: 20
  exclusions:
    presets:
      - comments
      - std-error-handling
      - common-false-positives
    rules:
      - linters:
          - staticcheck
        text: "SA1019"
      - linters:
          - revive
        path: "_test\\.go$"
        text: "dot-imports"
      - linters:
          - revive
        text: "unused-parameter"
      - linters:
          - revive
        text: "exported"
      - linters:
          - dupl
        path: "_test\\.go$"
      - linters:
          - unparam
        path: "_test\\.go$"
      - linters:
          - dupl
        path: "-test-suite\\.go$"
      - linters:
          - revive
        path: "-test-suite\\.go$"
        text: "dot-imports"

formatters:
  enable:
    - gofmt
    - goimports
```
</context>

<requirements>
1. **Update the Makefile `check` target** in `Makefile`:
   - Remove the two comment lines (the `# TODO: enable lint` comment and the `# check: lint vet errcheck vulncheck osv-scanner gosec trivy` commented-out target)
   - Change the `check:` dependency line from `check: vet errcheck vulncheck osv-scanner gosec trivy` to `check: lint vet errcheck vulncheck osv-scanner gosec trivy`

   Before:
   ```makefile
   # TODO: enable lint
   # check: lint vet errcheck vulncheck osv-scanner gosec trivy
   .PHONY: check
   check: vet errcheck vulncheck osv-scanner gosec trivy
   ```

   After:
   ```makefile
   .PHONY: check
   check: lint vet errcheck vulncheck osv-scanner gosec trivy
   ```

2. **Update `.golangci.yml`** in `.golangci.yml` to match the standard config from the reference standard config. Key differences to apply:
   - Add missing linters to the `enable` list: `nestif`, `errname`, `unparam`, `bodyclose`, `forcetypeassert`, `asasalint`, `prealloc`
   - Fix the typo in depguard deny rules: `"github.com/pkg/erros"` should be `"github.com/pkg/errors"` and `"use github.com/bborbe/erros instead"` should be `"use github.com/bborbe/errors instead"`
   - Add the additional depguard deny entries from the kv config (argument/v2, golang.org/x/net/context, golang.org/x/lint/golint, io/ioutil)
   - Add the `settings` blocks for `funlen`, `gocognit`, `nestif`, `maintidx`
   - Add additional exclusion rules from kv config (errname, unparam in tests, dupl/revive in test-suite files)
   - Keep any k8s-specific exclusion rules that exist (e.g. for k8s-specific error types if any)

3. **Run `make lint`** to identify all violations

4. **Fix all lint violations** in Go source files. Common fixes include:
   - Adding error checks for unchecked errors
   - Fixing unused parameters
   - Simplifying nested if statements
   - Removing duplicate code blocks
   - Adding type assertion safety checks (replace bare `x.(Type)` with checked `x, ok := x.(Type)`)
   - Closing HTTP response bodies
   - Pre-allocating slices where the capacity is known

   When fixing violations:
   - Prefer the minimal fix that satisfies the linter
   - Do NOT refactor or restructure code beyond what the linter requires
   - If a specific violation is a false positive in the k8s context (e.g. Kubernetes API patterns that look unusual to generic linters), add a targeted exclusion rule in `.golangci.yml` rather than changing the code
   - For `errname` violations on existing exported error types, add an exclusion rule in `.golangci.yml` rather than renaming (renaming would break consumers)

5. **Run `make lint` again** to confirm zero violations

6. **Run `make precommit`** to confirm the full pipeline passes
</requirements>

<constraints>
- Do NOT commit — dark-factory handles git
- Do NOT refactor code unrelated to fixing lint violations
- Do NOT add new features
- Do NOT change the `lint` target itself in the Makefile — it already has the correct command
- Do NOT remove or weaken linters — only add exclusion rules for genuine false positives
- Minimize code changes — fix the root cause of each violation with the smallest possible change
- Existing tests must still pass
</constraints>

<verification>
Run `make lint` — must pass with exit code 0.
Run `make precommit` — must pass with exit code 0.
Verify the Makefile no longer contains `# TODO: enable lint` or the commented-out check target.
Verify `.golangci.yml` includes all linters from the reference config (reference standard config).
</verification>
