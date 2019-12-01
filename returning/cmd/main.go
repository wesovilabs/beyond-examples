package main

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond-examples/returning/advice"
	"github.com/wesovilabs/beyond-examples/returning/greeting"
)

func Beyond() *api.Beyond {
	return api.New().
		WithReturning(advice.NewErrorsEnrichAdviceAdvice, "*.*(...)error")
}

func main() {
	checkError(greeting.Greetings("Hello", ""))
	checkError(greeting.Greetings("Bye", ""))
	checkError(greeting.Greetings("--", "John"))
	checkError(greeting.Greetings("Hello", "Sally"))
}

func checkError(err error){
	if err!=nil{
		fmt.Println(err.Error())
	}
}
