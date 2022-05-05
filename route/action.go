package route

import "context"

type IAction interface {
	Init(context.Context)
	Ctx() context.Context
	Prepare() error
	Run() (err error, content interface{})
}
