package route

import "context"

type IController interface {
	Init()
	Ctx(ctx ...context.Context) context.Context
	Prepare() error
	Run(IAction)
	In() *Input
	Out() *Output
}
