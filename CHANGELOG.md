# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], [markdownlint],
and this project adheres to [Semantic Versioning].

## [Unreleased]

-

## [0.3.6] - 2025-06-25

### Changed in 0.3.6

- Improved concurrency

## [0.3.5] - 2025-05-20

### Changed in 0.3.5

- Improved error messages

## [0.3.4] - 2025-04-23

### Changed in 0.3.4

- Updated dependencies

## [0.3.3] - 2024-09-03

### Added in 0.3.3

- Refactor to [template-go]

## [0.3.2] - 2024-06-03

### Added in 0.3.2

- grpcserver.Interface `GracefulStop()`

### Changed in 0.3.2

- Changed `GrpcServerImpl` to `SimpleGrpcServer`

## [0.3.1] - 2024-04-19

### Changed in 0.3.1

- Update dependencies
  - github.com/stretchr/testify v1.9.0
  - google.golang.org/grpc v1.63.2
  - google.golang.org/protobuf v1.33.0

## [0.3.0] - 2023-12-29

### Changed in 0.3.0

- Renamed module to `github.com/senzing-garage/go-observing`
- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - google.golang.org/grpc v1.60.1
  - google.golang.org/protobuf v1.32.0

## [0.2.8] - 2023-10-13

### Changed in 0.2.8

- Refactored to `template-go`
- Update dependencies
  - google.golang.org/grpc v1.58.3

## [0.2.7] - 2023-08-04

### Changed in 0.2.7

- Refactored to `template-go`
- Update dependencies
  - google.golang.org/grpc v1.57.0
  - google.golang.org/protobuf v1.31.0

## [0.2.6] - 2023-06-16

### Changed in 0.2.6

- Update dependencies
  - github.com/stretchr/testify v1.8.4
  - google.golang.org/grpc v1.56.0

## [0.2.5] - 2023-05-17

### Added in 0.2.5

- Notify() uses context.Background to avoid canceling Observation

## [0.2.4] - 2023-05-16

### Added in 0.2.4

- Support for gRPC Server options

## [0.2.3] - 2023-05-12

### Added in 0.2.3

- Fix bug (nil observer)

## [0.2.2] - 2023-05-10

### Added in 0.2.2

- gRPC service

### Changed in 0.2.2

- Method signature for notifier.Notify(...) adds "origin" parameter

## [0.2.1] - 2023-04-19

### Changed in 0.2.1

- Change time format to `time.Now().UTC().Format(time.RFC3339Nano)`

### Fixed in 0.2.1

- Fixed concurrency issue with unregistering observer

## [0.2.0] - 2023-03-03

### Added in 0.2.0

- `ObserverWhiteList` example observer

## [0.1.3] - 2023-03-01

### Added in 0.1.3

- `subject.GetObservers()`
- `notifier.Notify()`

## [0.1.2] - 2023-02-23

### Added in 0.1.2

- IsSilent feature to suppress output

## [0.1.1] - 2023-02-10

### Changed to 0.1.1

- Fixed proxy-pull

## [0.1.0] - 2023-02-02

### Added to 0.1.0

- Initial functionality

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[markdownlint]: https://dlaa.me/markdownlint/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
[template-go]: https://github.com/senzing-garage/template-go
