package route

import "context"

type IController interface {
	Init()
	Ctx(ctx ...context.Context) context.Context
	Prepare() error
	Run(action IAction) (int, []byte)
	In() *Input
	Out() *Output
	End()
}
