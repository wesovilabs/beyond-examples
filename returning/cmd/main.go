package main

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goaexamples/returning/advice"
	"github.com/wesovilabs/goaexamples/returning/greeting"
)

func Goa() *api.Goa {
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
