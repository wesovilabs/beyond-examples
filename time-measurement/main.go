package main

import (
	"github.com/wesovilabs/goaexamples/timemeasurement/measurement"
)

func main() {
	measurement.MyFunction("very-slow", false)
	measurement.MyFunction("slow", false)
	measurement.MyFunction("medium", false)
	measurement.MyFunction("fast", false)
	measurement.MyFunction("very-fast", false)
}
