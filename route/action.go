package route

import "context"

type IAction interface {
	Init()
	Ctx(ctx ...context.Context) context.Context
	Prepare() error
	Run() (error, interface{})
	End()
}
