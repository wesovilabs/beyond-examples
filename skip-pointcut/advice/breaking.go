package advice

import (
	"errors"
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
)

type BreakingAdvice struct {
	prefix string
}

func (a *BreakingAdvice) Before(ctx *context.BeyondContext) {
	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	if ctx.Function()=="Hello"{
		fmt.Println("Leaving the flow")
		ctx.Results().SetAt(0,errors.New("this is awesome!"))
		ctx.Exit()
	}
}


func (a *BreakingAdvice) Returning(ctx *context.BeyondContext) {
	fmt.Println("Returning")
}

func NewBreakingAdvice() api.Around {
	return &BreakingAdvice{}
}

