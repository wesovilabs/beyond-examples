package main

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond-examples/skipping/advice"
	"github.com/wesovilabs/beyond-examples/skipping/greeting"
)

func Beyond() *api.Beyond {
	return api.New().
		WithAround(advice.NewBreakingAdvice, "greeting.*(...)...")
}

func main() {
	if err:=greeting.Hello("John");err!=nil{
		fmt.Println(err.Error())
	}
	if err:=greeting.Bye("John");err!=nil{
		fmt.Println(err.Error())
	}
}
