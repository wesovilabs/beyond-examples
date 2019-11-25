package main

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goaexamples/around/advice"
	"github.com/wesovilabs/goaexamples/around/greeting"
)

func Goa() *api.Goa {
	return api.New().
		WithAround(advice.NewTimerAdvice(advice.Microseconds), "greeting.Hello(string)...").
		WithAround(advice.NewTimerAdvice(advice.Nanoseconds), "greeting.Bye(string)...")
}

func main() {
	greeting.Greetings("Hello", "John")
	greeting.Greetings("Bye", "John")
}
