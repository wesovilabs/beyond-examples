package measurement

import (
	"errors"
	"fmt"
	"time"
)

var times = map[string]int32{
	"very-slow": 5,
	"slow":      4,
	"medium":    3,
	"fast":      2,
	"very-fast": 1,
}

func MyFunction(msg string, returnError bool) error {
	fmt.Printf("[%s] %s\n",time.Now().Format("15:04:05",),msg)
	time.Sleep(time.Duration(times[msg]) * time.Second)
	if returnError {
		return errors.New("unexpected error")
	}
	return nil
}
