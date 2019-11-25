package main

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goaexamples/timemeasurement/advice"
	"github.com/wesovilabs/goaexamples/timemeasurement/measurement"
)

func Goa() *api.Goa {
	return api.New().WithAround(advice.NewMyAdvice(float64(1.1)), "*.*(...)...")
}

func main() {
	if err:=measurement.MyFunction("very-slow", false);err!=nil{
		panic(err.Error())
	}
	if err:=measurement.MyFunction("slow", false);err!=nil{
		panic(err.Error())
	}
	if err:=measurement.MyFunction("medium", false);err!=nil{
		panic(err.Error())
	}
	if err:=measurement.MyFunction("fast", false);err!=nil{
		panic(err.Error())
	}
	if err:=measurement.MyFunction("very-fast", false);err!=nil{
		panic(err.Error())
	}
}
