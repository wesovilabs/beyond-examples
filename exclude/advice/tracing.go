package advice

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
	"strings"
)

type TracingAdvice struct {
	prefix string
}

func (a *TracingAdvice) Before(ctx *context.BeyondContext) {
	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	printTrace(ctx,a.prefix,params)
}

func NewTracingAdvice() api.Before {
	return &TracingAdvice{}
}

func NewTracingAdviceWithPrefix(prefix string) func() api.Before {
	return func() api.Before {
		return &TracingAdvice{
			prefix: prefix,
		}
	}
}

func printTrace(ctx *context.BeyondContext, prefix string, params []string){
	if prefix == "" {
		fmt.Printf("%s.%s(%s)\n", ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
		return
	}
	fmt.Printf("%s %s.%s(%s)\n", prefix, ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
}
