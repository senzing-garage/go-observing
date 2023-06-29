# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

-

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
