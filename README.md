# go-observing

## Synopsis

Implements the
[Observer](https://en.wikipedia.org/wiki/Observer_pattern)
software design pattern.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-observing.svg)](https://pkg.go.dev/github.com/senzing/go-observing)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/go-observing)](https://goreportcard.com/report/github.com/senzing/go-observing)
[![go-test.yaml](https://github.com/Senzing/go-observing/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/go-observing/actions/workflows/go-test.yaml)

## Overview

Examples of using `go-observing` can be seen in
[main.go](main.go).

As an example, a null object,
[ObserverNull](observer/observer_null.go),
shows how an observer is used in the Observer software design pattern.
It is an example only.
Any code wishing to observe, would write their own code
which adheres to the
[Observer](observer/main.go)
interface.

## References

- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
- [Package reference](https://pkg.go.dev/github.com/senzing/go-observing)
