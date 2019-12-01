package main

import (
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond-examples/exclude/advice"
	"github.com/wesovilabs/beyond-examples/exclude/greeting"
)

func Beyond() *api.Beyond {
	return api.New().
		WithBefore(advice.NewTracingAdvice, "*.*(...)...").
		WithBefore(advice.NewTracingAdviceWithPrefix("[beyond]"), "greeting.Bye(...)...").
		Exclude("advice.*(...)...")
}

func main() {
	greeting.Hello("John")
	greeting.Bye("John")
}
