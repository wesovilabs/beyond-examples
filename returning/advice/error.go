package advice

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
	"strings"
)

type ErrorsEnrichAdvice struct {
}

func (a *ErrorsEnrichAdvice) Returning(ctx *context.BeyondContext) {
	if index, result := ctx.Results().Find(isError);index>=0{
		ctx.Results().SetAt(index, &CustomError{err: result.Value().(error), pkg: ctx.Pkg(), function: ctx.Function(), params: ctx.Params()})
	}
}

func isError(_ int, arg *context.Arg) bool{
	return arg.IsError()
}


type CustomError struct {
	err      error
	pkg      string
	function string
	params   *context.Args
}

func (e *CustomError) Error() string {
	params := make([]string, e.params.Count())
	e.params.ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	errDetail := fmt.Sprintf("%s.%s(%s)", e.pkg, e.function, strings.Join(params, ","))
	return fmt.Sprintf("[%s] => %s", errDetail, e.err.Error())
}

func NewErrorsEnrichAdviceAdvice() api.Returning {
	return &ErrorsEnrichAdvice{}
}
