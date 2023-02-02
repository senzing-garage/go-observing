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

### Developing

### Testing

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=go-observing
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test

    ```
