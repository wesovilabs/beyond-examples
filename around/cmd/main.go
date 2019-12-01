package main

import (
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond-examples/around/advice"
	"github.com/wesovilabs/beyond-examples/around/greeting"
)

func Beyond() *api.Beyond {
	return api.New().
		WithAround(advice.NewTimerAdvice(advice.Microseconds), "greeting.Hello(string)...").
		WithAround(advice.NewTimerAdvice(advice.Nanoseconds), "greeting.Bye(string)...")
}

func main() {
	greeting.Greetings("Hello", "John")
	greeting.Greetings("Bye", "John")
}
