# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.12.0

- update go and deps

## v1.11.0

- Add PodWatcher interface and implementation for watching Kubernetes pod changes
- Add PodEventProcessor interface for handling pod update and delete events
- Add PodEventProcessorSkipError wrapper for gracefully skipping errors in pod event processing
- Add PodInterface mock for testing
- Add NewPodWatcherRetry wrapper for automatic retry with configurable wait duration
- Add ErrResultChannelClosed error constant for watch channel closure detection
- Add ErrUnknownEventType error constant for unknown K8s event types
- **Improve error semantics**: Watchers now return ErrResultChannelClosed instead of nil when watch connection closes
- Update ServiceWatcher and SecretWatcher to return ErrResultChannelClosed on channel closure
- Add comprehensive test coverage for PodWatcher with all event types and error scenarios

## v1.10.0

- Add SecretWatcher interface and implementation for watching Kubernetes secret changes
- Add SecretEventProcessor interface for handling secret update and delete events
- Add SecretEventProcessorSkipError wrapper for gracefully skipping errors in secret event processing
- Add SecretInterface mock for testing
- **Refactor watchers with decorator pattern**: Separate retry logic from base watchers
- Add NewServiceWatcherRetry wrapper for automatic retry with configurable wait duration
- Add NewSecretWatcherRetry wrapper for automatic retry with configurable wait duration
- Add godoc comments to exported constructors (NewServiceWatcher, NewSecretWatcher)
- **BREAKING CHANGE**: NewServiceWatcher and NewSecretWatcher constructors no longer accept waiterDuration parameter - use retry wrappers instead
- Update dependencies (go.mod and go.sum)

## v1.9.3

- Update Go version to 1.25.2 in CI workflow
- Upgrade osv-scanner to v2 with improved config file support
- Update dependencies (go.mod and go.sum)

## v1.9.2

- rename NewServiceManagerSkipError -> NewServiceEventProcessorSkipError

## v1.9.1

- Add ServiceWatcher interface and implementation for watching Kubernetes service changes
- Add ServiceEventProcessor interface for handling service update and delete events
- Add ServiceEventProcessorSkipError wrapper for gracefully skipping errors in service event processing
- Add comprehensive test coverage for ServiceWatcher with all event types and error scenarios

## v1.9.0

- Add comprehensive package-level documentation (doc.go)
- Add godoc comments to all exported types and interfaces
- Improve README.md with installation guide, features, quick start examples, and API documentation
- Clean up CLAUDE.md to reference coding guidelines instead of duplicating content
- Update mocks and regenerate after adding documentation
- **WARNING**: Breaking changes in CronJobBuilder - SetParallelism, SetBackoffLimit, and SetCompletions now accept int32 instead of int

## v1.8.9

- go mod update

## v1.8.8

- add ApiextensionsClientset 

## v1.8.7

- add github workflow

## v1.8.6

- go mod update

## v1.8.5

- add tests
- go mod update
  `
## v1.8.4

- add JobAlreadyExistsError in JobDeployer  

## v1.8.3

- add MarshalAsYaml for testing

## v1.8.2

- remove undeploy from job deployer

## v1.8.1

- add CronScheduleExpression type with validation

## v1.8.0

- add CronJobBuilder and CronJobDeployer 

## v1.7.4

- add JobBuilder setters for TTLSecondsAfterFinished, CompletionMode and PodReplacementPolicy

## v1.7.3

- increase test coverage for all deployer components and missing builder components

## v1.7.2

- add Job restart policy validation and tests

## v1.7.1

- add ObjectMetaBuilder.SetLabels and ObjectMetaBuilder.SetAnnotations

## v1.7.0

- add HasBuild interfaces and HasBuildFunc funcs

## v1.6.3

- fix job builder app label

## v1.6.2

- fix job api version

## v1.6.1
 
- allow set PriorityClassName in PodSpec
- allow set parallelism, completions and backoffLimit in JobBuilder
- go mod update

## v1.6.0

- remove vendor
- go mod update

## v1.5.2

- add SetImagePullSecrets to PodSpecBuilder 
- go mod update

## v1.5.1

- fix DeploymentBuilder and StatefulSetBuilder

## v1.5.0

- allow define imagePullSecrets
- go mod update

## v1.4.0

- add ResourceEventHandler
- add EventHandler

## v1.3.8

- improve error message
- go mod update

## v1.3.7

- add k8s_ prefix to all go files

## v1.3.6

- allow BuildName with number at the end
- go mod update
- remove deprecated golint

## v1.3.5

- add name from pod function
- go mod update

## v1.3.4

- go mod update

## v1.3.3

- set some defaults for jobs

## v1.3.2

- allow set affinity in podSpec

## v1.3.1

- allow set restartPolicy on containerBuilder
- go mod update

## v1.3.0

- update to k8s v0.31.0
- go mod update

## v1.2.0

- add Name type
- use Name in services
- move validation

## v1.1.0

- skip ingress update if equal
- go mod update

## v1.0.0

- Initial Version
