# go-events

[![GoDoc](https://godoc.org/github.com/ronelliott/go-events?status.png)](https://godoc.org/github.com/ronelliott/go-events)
[![Build Status](https://travis-ci.org/ronelliott/go-events.svg?branch=master)](https://travis-ci.org/ronelliott/go-events)
[![Coverage Status](https://coveralls.io/repos/github/ronelliott/go-events/badge.svg?branch=master)](https://coveralls.io/github/ronelliott/go-events?branch=master)

A go library for writing command line interfaces. Only supports go versions newer than, or equal to, 1.5

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
