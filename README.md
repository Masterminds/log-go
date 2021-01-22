# Log Package

The log package provides a common interface that can be used in applications and
libraries along with reference implementation wrappers for logrus, zap, the Go
standard library package, and for a CLI.

The problem is that there are many logging packages where each has its own
interface. This makes it difficult to mix and match libraries and code that
depend on different interfaces. It also makes it difficult to change logging
libraries if the one you are using becomes deprecated or you would like to
switch to a different one (e.g., you want to tie into a logging service).

The log package aims to provide a solution to this problem.

## Usage

There are a couple parts to using the log package.

### Using The Interface

The core of the package is the logging interface meant to be used in applications.
It can be used when you need to have a logger passed round or when want to call
a global logger.

```go
import "github.com/mattfarina/log"

func NewConstructorExample(*log.Logger) {
    ...
}

func (e *Example) Foo() {
    e.logger.Info("Send Some Info")
}

```

In this example a logger is passed into a constructor and used in a method. What
is the logger implementation? It could be any logger that either provides an
implementation of the interface or has a wrapper written around it to do so. In
the _impl_ directory are 4 reference implementations of loggers that work with
this interface.

Need to change the logger out? All you need to do is change the instance you use
in the setup. This makes it easy to change logger later or based on your case.
For example, you could have a single library and for a CLI application send the
output to the console and in a service log the output elsewhere.

There is also a global logger that you can use.

```go
import(
    "github.com/mattfarina/log"
)

log.Info("Send Some Info")
```

Before you can use the global logger a logger needs to be setup. For example,
if you are using logrus it could look like...

```go
import(
    "github.com/mattfarina/log"
    "github.com/mattfarina/log/impl/logrus"
)

// This step only needs to happen once in an application.
log.Current = logrus.NewStandard()

// All through the application the global/singleton calls will now work.
log.Info("Send Some Info")
```

The current reference loggers have multiple configurations. For example, you can
provide fully custom configuration to logrus.

### Providing A Logger

What does it take to create a new logger or provide a logger? It just takes
creating a struct that conforms to the `Logger` interface.

## Log Levels

The following log levels are currently implemented across the interface and all
the reference implementations.

- Fatal
- Error
- Warn
- Info
- Debug
