# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-observing is a Go library implementing the Observer software design pattern with gRPC support for distributed observation. It allows local in-process observation and remote observation via gRPC.

## Build and Development Commands

```bash
# Run tests (requires setup first for generated protobuf code)
make clean setup test

# Run a single test
go test -v -run TestFunctionName ./path/to/package

# Run linting (golangci-lint + govulncheck + cspell)
make lint

# Generate protobuf code (regenerates observerpb/ from observer.proto)
make generate

# Update Go dependencies
make dependencies

# Auto-fix lint issues
make fix

# Run code coverage
make coverage

# Check coverage thresholds
make check-coverage
```

## Architecture

The library implements the Observer pattern with four core packages:

**observer/** - The Observer interface and implementations:

- `Observer` interface: `GetObserverID()` and `UpdateObserver(ctx, message)`
- `NullObserver`: No-op implementation for testing
- `GrpcObserver`: Forwards messages to a remote gRPC server
- `RawObserver`: Receives raw message strings
- `WhiteListObserver`: Filters messages based on a whitelist

**subject/** - The Subject that manages observers:

- `Subject` interface: `RegisterObserver()`, `UnregisterObserver()`, `NotifyObservers()`, `HasObservers()`, `GetObservers()`
- `SimpleSubject`: Thread-safe implementation using sync.RWMutex
- Notifications are sent asynchronously via goroutines

**notifier/** - Utility for structured JSON notification messages:

- `Notify()` creates JSON messages with origin, subjectId, messageId, messageTime, and custom details

**grpcserver/** - gRPC server for remote observation:

- `SimpleGrpcServer`: Receives messages via gRPC and forwards to local observers
- Enables chaining: local Subject → GrpcObserver → remote GrpcServer → remote Observers

**observerpb/** - Generated protobuf code (do not edit manually):

- Generated from `observer.proto` via `make generate`

## Code Style

- Max line length: 120 characters
- Formatter: gofumpt
- Linting: golangci-lint with extensive linter set (see `.github/linters/.golangci.yaml`)
- Test framework: testify (github.com/stretchr/testify)

## gRPC Development

When modifying the gRPC interface:

1. Edit `observer.proto`
2. Run `make generate` to regenerate `observerpb/` files
3. Requires `protoc` compiler installed
