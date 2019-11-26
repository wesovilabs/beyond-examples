package main

import (
	"github.com/wesovilabs/goaexamples/timemeasurement/measurement"
)



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
