package advice

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
	"time"
)

type MyAdvice struct {
	limit float64 // In seconds
}

func (a *MyAdvice) Before(context *context.GoaContext) {
	context.Set("advice.start", time.Now())
}

func (a *MyAdvice) Returning(goaContext *context.GoaContext) {
	startTime := goaContext.Get("advice.start").(time.Time)
	spentTime := time.Now().Sub(startTime).Seconds()
	fmt.Printf(" -> took %v seconds\n", spentTime)
}

func NewMyAdvice(limit float64) func() api.Around {
	return func() api.Around {
		return &MyAdvice{limit}
	}
}
