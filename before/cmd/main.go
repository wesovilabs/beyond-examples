package main

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goaexamples/before/advice"
	"github.com/wesovilabs/goaexamples/before/greeting"
)

func Goa() *api.Goa {
	return api.New().
		WithBefore(advice.NewTracingAdvice, "greeting.Hello(...)...").
		WithBefore(advice.NewTracingAdviceWithPrefix("[goa]"), "greeting.Bye(...)...")
}

func main() {
	greeting.Hello("John")
	greeting.Bye("John")
}
