# Testing with Std Lib Example

When using a logger it is common to write tests to verify correct handling in
an application. One way to do that is to use a buffer to collect the log
messages and then check the content of the buffer.

The code in this example showcases using the Go standard library logger in
a testing setup with a buffer to collect the messages.