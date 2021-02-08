# Log Package

The log package provides a common interface that can be used in applications and
libraries along with reference implementation wrappers for logrus, zap, the Go
standard library package, and for a CLI.

*_Note: this codebase is under active development. Consider it pre-alpha code._*

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

There are a couple parts to using the log package.

### Using The Interface

The easiest way to get started is to import the package and start logging.

```go
import(
    "github.com/mattfarina/log"
)

log.Info("Send Some Info")
```

This uses the global interface which is one of the two ways to use the interface.

The real power is in the ability to change the logger. In the setup to your
application you can set the logger like so...

```go
import(
    "github.com/mattfarina/log"
    "github.com/mattfarina/log/impl/logrus"
)

log.Current = logrus.NewStandard()
```

Any calls to log following this will be sent to logrus. The `logrus.NewStandard()`
function provides a default logrus client. There is another constructor that
lets you provide a custom logrus configuration.

The loggers provided by this library are reference implementations. The real
power is in the interface where you can create you own implementation.

There is a second way to use the interface in your applications and packages.

```go
import "github.com/mattfarina/log"

func NewConstructorExample(logger log.Logger) {
    return &Example{
        logger: logger,
    }
}

func (e *Example) Foo() {
    e.logger.Info("Send Some Info")
}

```

In this setup the interface `log.Logger` is is used in the constructor. Any of
the implementations can be passed in an instance. When the function `Foo` is
called the logger in use will do the logging. It could be any logger that works
for the interface.

### Providing A Logger

What does it take to create a new logger or provide a logger? It just takes
creating a struct that conforms to the `log.Logger` interface. There are several
reference implementations in the _impl_ directory to help you understand the
structure.

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
