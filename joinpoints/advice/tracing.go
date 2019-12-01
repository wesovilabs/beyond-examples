package advice

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
	"reflect"
)

type SimpleTracingAdvice struct {
}

func (advice *SimpleTracingAdvice) Before(ctx *context.BeyondContext) {
	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	if ctx.Type() != nil {
		fmt.Printf("%s.%s.%v\n", ctx.Pkg(), reflect.TypeOf(ctx.Type()), ctx.Function())
		return
	}
	fmt.Printf("%s.%s\n", ctx.Pkg(), ctx.Function())
}

func NewSimpleTracingAdvice() api.Before { return &SimpleTracingAdvice{} }
