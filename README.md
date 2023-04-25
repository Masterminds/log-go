# Log Package

The log package provides a common interface for logging that can be used in
applications and libraries along with reference implementations for logrus, zap,
the Go standard library package, and for a CLI (console app).

[![Go Report Card](https://goreportcard.com/badge/github.com/Masterminds/log-go)](https://goreportcard.com/report/github.com/Masterminds/log-go)
[![](https://github.com/Masterminds/semver/workflows/Tests/badge.svg)](https://github.com/Masterminds/log-go/actions)
[![GoDoc](https://img.shields.io/static/v1?label=go.dev&message=reference&color=blue)](https://pkg.go.dev/github.com/Masterminds/log-go)

## Why A Log Interface

In many programming languages there is a logging interface. Libraries will use
use the interface for logging and then the application will choose the logging
library of their choice for the application as a whole. Since everything follows
the interface it works.

Go does not have a detailed interface. The logging package included in the
standard library has been insufficient so many logging libraries have been
created.

Various libraries/packages use the a logging implementation of their choice.
When those are pulled into a larger application is will end up having multiple
different logging systems. The main application needs to wire them all up or it
will miss some logs. If you take a look at applications like Kubernetes, Nomad,
and many others you'll discover they import many different logging implementations.

Using interfaces provides a way to move away from multiple implementation,
simplify codebases, reduce binary size, and reduce threat surface area.

## What This Package Provides

This library includes several things including:

- A Go interface for leveled logging. Those levels include - Fatal, Panic, Error,
  Warn, Info, Debug, Trace
- When logging messages the interface can do a message, a message with formatting
  string and arguments, and a message with fields (key/value pairs)
- Package level logging functions whose implementation can be changed/set
- Reference implementations for logrus, zap, the standard library, and a CLI
- A simple default implementation so it just works for library testing and
  simple situations

## Usage

The usage documentation is broken down into 3 types of usage depending on your
situation. These examples are for library/package authors, applications, and
logger implementation developers.

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

For details on exactly which functions are on the package or on the `Logger`
interface, please see the [package docs](https://pkg.go.dev/github.com/Masterminds/log-go).

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
