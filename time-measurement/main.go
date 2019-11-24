package main

import "github.com/wesovilabs/goaexamples/timemeasurement/measurement"

func main() {
	measurement.MyFunction("very-slow", false)
	measurement.MyFunction("slow", false)
	measurement.MyFunction("normal", false)
	measurement.MyFunction("fast", false)
	measurement.MyFunction("very-fast", false)
	measurement.MyFunction("Stop testing on Animals", false)
}
