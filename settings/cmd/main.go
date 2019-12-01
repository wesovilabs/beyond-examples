package main

import (
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond-examples/settings/advice"
	"github.com/wesovilabs/beyond-examples/settings/greeting"
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
