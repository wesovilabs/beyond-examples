package main

import (
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond-examples/before/advice"
	"github.com/wesovilabs/beyond-examples/before/greeting"
)

func Beyond() *api.Beyond {
	return api.New().
		WithBefore(advice.NewTracingAdvice, "greeting.Hello(...)...").
		WithBefore(advice.NewTracingAdviceWithPrefix("[beyond]"), "greeting.Bye(...)...")
}

func main() {
	greeting.Hello("John")
	greeting.Bye("John")
}
