# go-observing development

The following instructions are useful during development.

**Note:** This has been tested on Linux and Darwin/macOS.
It has not been tested on Windows.

## Prerequisites for development

:thinking: The following tasks need to be complete before proceeding.
These are "one-time tasks" which may already have been completed.

1. The following software programs need to be installed:
    1. [git]
    1. [make]
    1. [docker]
    1. [go]

## Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing-garage
    export GIT_REPOSITORY=go-observing
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow
   steps in [clone-repository] to install the Git repository.

## Dependencies

1. A one-time command to install dependencies needed for `make` targets.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make dependencies-for-development

    ```

1. Install dependencies needed for [Go] code.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make dependencies

    ```

## Lint

1. Run linting.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make lint

    ```

## Developing gRPC client and server SDK

The following instructions were used to create a
[go module] and other [example generated source code].

1. Follow the [Go Quick start] tutorial to prepare an environment.

1. [Generating client and server code].
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make generate

    ```

    1. In `${GIT_REPOSITORY_DIR}/observerpb`, files *with* `_grpc.` in the filename contain the following:
        - Interface types (or stubs) for clients to call with the methods defined in the services.
        - Interface types for servers to implement, also with the methods defined in the services.
        - In other words, its the "gRPC" code that handles the network traffic, not the message content.
    1. In `${GIT_REPOSITORY_DIR}/observerpb`, files *without* `_grpc.` in the filename contain the following:
        - protocol buffer code to populate, serialize, and retrieve request and response message types.
        - In other words, it manages message content, not the network traffic.
1. References:
    1. [gRPC Documents for Go]
        1. [Go Quick start]
    1. [Thread safety]

## Test

1. Run tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean setup test

    ```

## Coverage

Create a code coverage map.

1. Run Go tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean setup coverage

    ```

   A web-browser will show the results of the coverage.
   The goal is to have over 80% coverage.
   Anything less needs to be reflected in [testcoverage.yaml].

## Documentation

1. View documentation.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean documentation

    ```

1. If a web page doesn't appear, visit [localhost:6060].
1. Senzing documentation will be in the "Third party" section.
   `github.com` > `senzing-garage` > `go-observing`

1. When a versioned release is published with a `v0.0.0` format tag,
the reference can be found by clicking on the following badge at the top of the README.md page.
Example:

    [![Go Reference Badge]][Go Reference]

1. To stop the `godoc` server, run

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

## References

[clone-repository]: https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/clone-repository.md
[docker]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/docker.md
[example generated source code]: example_generated_source_code
[Generating client and server code]: https://grpc.io/docs/languages/go/basics/#generating-client-and-server-code
[git]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/git.md
[go module]: ../grpc
[Go Quick start]: https://grpc.io/docs/languages/go/quickstart/
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/template-go.svg
[Go Reference]: https://pkg.go.dev/github.com/senzing-garage/template-go
[go]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/go.md
[gRPC Documents for Go]: https://grpc.io/docs/languages/go/
[localhost:6060]: http://localhost:6060/pkg/github.com/senzing-garage/template-go/
[make]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/make.md
[testcoverage.yaml]: ../.github/coverage/testcoverage.yaml
[Thread safety]: https://grpc.io/docs/languages/go/generated-code/
