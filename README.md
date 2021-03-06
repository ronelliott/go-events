# go-events

[![GoDoc](https://godoc.org/github.com/ronelliott/go-events?status.png)](https://godoc.org/github.com/ronelliott/go-events)
[![Build Status](https://travis-ci.org/ronelliott/go-events.svg?branch=master)](https://travis-ci.org/ronelliott/go-events)
[![Coverage Status](https://coveralls.io/repos/github/ronelliott/go-events/badge.svg?branch=master)](https://coveralls.io/github/ronelliott/go-events?branch=master)

A basic event emitter written in go.

## Installation
  $ go get github.com/ronelliott/go-events

## Examples

### Basic:

```go
package main

import (
  "fmt"
  "github.com/ronelliott/go-events"
)

func main() {
  bus := events.NewBus()

  bus.On("something-awesome", func(data ...interface{}) {
      fmt.Println("Do something awesome with your data!", data)
  })

  bus.Emit("something-awesome", "foo", "bar")
}
```

## Server:
  *Server is a WIP*

## Benchmarks:

```
BenchmarkBus_Emit_One-12             2000000         656 ns/op
BenchmarkBus_Emit_Ten-12             2000000         649 ns/op
BenchmarkBus_Emit_Twenty-12          2000000         653 ns/op
BenchmarkBus_Emit_Thirty-12          2000000         646 ns/op
BenchmarkBus_Emit_Forty-12           2000000         652 ns/op
BenchmarkBus_Emit_OneHundred-12      2000000         648 ns/op
```
