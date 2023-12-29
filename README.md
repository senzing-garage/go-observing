# go-observing

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## Synopsis

Implements the
[Observer](https://en.wikipedia.org/wiki/Observer_pattern)
software design pattern.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/go-observing.svg)](https://pkg.go.dev/github.com/senzing-garage/go-observing)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/go-observing)](https://goreportcard.com/report/github.com/senzing-garage/go-observing)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/go-observing/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/go-observing/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/go-observing/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/go-observing/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/go-observing/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/go-observing/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/go-observing/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/go-observing/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/go-observing/actions/workflows/go-test-windows.yaml)

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

## Use

### Observer gRPC Server

In addition to local, "in-process", observation,
the `go-observer` repository also supports a gRPC-based aggregator of observer messages.

The following image shows flow of messages.

![Image of architecture](docs/img/repeater.png)

The Subject notifies local Observers.  One of the Observers "repeats"
the message by sending it via gRPC to a GrpcServer that embeds a Subject.
That Subject notifies remote Observers.

To create a GrpcServer, a Subject is created with Observers and wrapped with a GrpcServer.
Example:

```go

package main

import (
    "context"

    "github.com/senzing-garage/go-observing/grpcserver"
    "github.com/senzing-garage/go-observing/observer"
    "github.com/senzing-garage/go-observing/subject"
)

func main() {
    ctx := context.TODO()

    // Create a Subject.

    aSubject := &subject.SubjectImpl{}

    // Register Observer(s) with the Subject.

    anObserver1 := &observer.ObserverNull{
        Id: "Observer 1",
    }
    err = aSubject.RegisterObserver(ctx, anObserver1)
    if err != nil {
        fmt.Print(err)
    }

    // Start gRPC service with an embedded Subject.

    aGrpcServer := &grpcserver.GrpcServerImpl{
        Port:    8260,
        Subject: aSubject,
    }
    aGrpcServer.Serve(ctx)
```

A working GrpcServer can be seen in
[main.go](main.go).

The
[ObserverGrpc](observer/observer_grpc.go)
is an Observer that sends messages to a GrpcServer.

## References

1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
1. [Package reference](https://pkg.go.dev/github.com/senzing-garage/go-observing)
