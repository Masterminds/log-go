package main

import (
	"io"

	"github.com/mattfarina/log-go"
	logio "github.com/mattfarina/log-go/io"
)

// Sometimes you have an io.Writer interface and you want to have it write to
// the logs. Since loggers do things a variety of ways it is difficult to ask
// each of them to expose an underlying io.Writer. It is easier to be able to
// generate one that leverages the log.Logger api. The following is an example
// that does just that.

func main() {
	// Get an io.Writer compliant instance that will write messages at the info
	// level. This uses the Current configured logger.
	w := logio.NewCurrentWriter(log.InfoLevel)

	// io.WriteString accepts an io.Writer as its first argument. This shows
	// writing the string "foo" to the logger at the info level.
	_, _ = io.WriteString(w, "foo")
}
