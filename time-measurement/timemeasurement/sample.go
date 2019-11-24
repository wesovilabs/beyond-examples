package main

import (
	"errors"
	"fmt"
	"time"
)

var times = map[string]float32{
	"very-slow": 4,
	"slow":      3,
	"medium":    2,
	"fast":      1,
	"very-fast": 0.5,
}

func myFunction(msg string, returnError bool) error {
	fmt.Println(msg)
	time.Sleep(time.Duration(times[msg]) * time.Second)
	if returnError {
		return errors.New("unexpected error")
	}
	return nil
}

type myType struct {
}

func (t *myType) myMethod(msg string, returnError bool) error {
	fmt.Println(msg)
	time.Sleep(time.Duration(times[msg]) * time.Second)
	if returnError {
		return errors.New("unexpected error")
	}
	return nil
}

func main() {
	myFunction("very-slow", false)
	myFunction("slow", false)
	myFunction("normal", false)
	myFunction("fast", false)
	myFunction("very-fast", false)
	myFunction("Stop testing on Animals", false)
}
