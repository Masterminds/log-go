# Log Package

The log package provides a common interface that can be used in applications and
libraries along with reference implementation wrappers for logrus, zap, the Go
standard library package, and for a CLI.

[![Go Report Card](https://goreportcard.com/badge/github.com/Masterminds/log-go)](https://goreportcard.com/report/github.com/Masterminds/log-go)
[![](https://github.com/Masterminds/semver/workflows/Tests/badge.svg)](https://github.com/Masterminds/log-go/actions)
[![GoDoc](https://img.shields.io/static/v1?label=go.dev&message=reference&color=blue)](https://pkg.go.dev/github.com/Masterminds/log-go)

## The Problem

The problem is that there are many logging packages where each has its own
interface. This makes it difficult to mix and match libraries and code that
depend on different interfaces. It also makes it difficult to change logging
libraries if the one you are using becomes deprecated or you would like to
switch to a different one (e.g., you want to tie into a logging service).

The log package aims to provide a solution to this problem.

## Use Cases

This library was born out of a specific problem that also presents itself as a
solution for a second use case.

### Use Case 1: Console and Logging

I was working on an application that needs to write output to the console
(sometimes with color) and to logrus. logrus is in the maintenance stage of
development (it's even recommending other logging packages) so I can see it
someday being replaced in the system I'll have to deal with.

I needed a solution that let me write messages that in some situations are sent
to console output and in other situations are sent to logs. In the first
situation it is as part of a CLI application. In the second case the code is
imported as a library in a service.

The solution is to use an interface with varying implementations.

### Use Case 2: Logging In Go Is Diverse Which Makes It Hard

Rust, PHP, and many other languages have APIs for logging that different
implementations can implement. This makes logging pluggable and improves library
interoperability.

This is a secondary use case and benefit from having an interface.

## Usage

The usage documentation is broken down into 3 types of usage depending on your
situation.

### Library / Package Authors

If you are a library or package author there are two ways you can use this log
package.

First, you can import the package and use the package level logging options. For
example:

```go
import(
    "github.com/Masterminds/log-go"
)

log.Info("Send Some Info")
```

You can use this for logging within your package.

Second, if you want to pass a logger around your package you can use the
interface provided by this package. For example:

```go
import "github.com/Masterminds/log-go"

func NewConstructorExample(logger log.Logger) {
    return &Example{
        logger: logger,
    }
}

func (e *Example) Foo() {
    e.logger.Info("Send Some Info")
}

```

In your packages testing you can check log messages if you need to see that they
are working and contain what you are looking for. A simple example of doing this
is in the `_examples` directory.

### Application Developers

If you are developing an application that will be writing logs you will want to
setup and configure logging the way you want. This is where this interface
based package empowers you. You can pick your logging implementation or write
your own.

For example, if you want to use a standard logrus logger you can setup logging
like so:

```go
import(
    "github.com/Masterminds/log-go"
    "github.com/Masterminds/log-go/impl/logrus"
)

log.Current = logrus.NewStandard()
```

In this example a standard logrus logger is created, wrapped in a struct
instance that conforms to the interface, and is set as the global logger to use.

The `impl` directory has several reference implementations and they have
configurable setups.

Once you setup the global logger the way you want all the packages will use this
same logger.

### Logger Developers

There are many loggers and many ways to record logs. They can be written to a
file, sent to stdout/stderr, sent to a logging service, and more. Each of these
is possible with this package.

If you have logger you want to use you can write one that conforms to the
`Logger` interface found in `log.go`. That logger can then be configured as
documented in the previous section.

The `impl` directory has reference implementations you can look at for further
examples.

## Log Levels

The following log levels are currently implemented across the interface and all
the reference implementations.

- Fatal
- Panic
- Error
- Warn
- Info
- Debug
- Trace
