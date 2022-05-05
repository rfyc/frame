package route

import "context"

type IController interface {
	Init(context.Context, *Input, IAction)
	Ctx() context.Context
	Prepare() error
	Out(err error, content interface{}) *Output
}
