package advice

import (
	"fmt"
	"github.com/pkg/errors"
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
	takenTime := time.Now().Sub(startTime).Seconds()

	if takenTime > a.limit {
		if index, _ := goaContext.Results().Find(
			func(_ int, arg *context.Arg) bool {
				return arg.IsError()
			}); index > -1 {
			goaContext.Results().
				SetAt(
					index,
					errors.New(
						fmt.Sprintf("function took %v seconds",
							takenTime)),
				)
		}
	}
}

func NewMyAdvice(limit float64) func() api.Around {
	return func() api.Around {
		return &MyAdvice{limit}
	}
}
