package measurement

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

func MyFunction(msg string, returnError bool) error {
	fmt.Println(msg)
	time.Sleep(time.Duration(times[msg]) * time.Second)
	if returnError {
		return errors.New("unexpected error")
	}
	return nil
}
