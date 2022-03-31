package route

import "context"

type IController interface {
	Init()
	Ctx(ctx ...context.Context) context.Context
	Prepare() error
	RunAction(runFunc func() (error, interface{})) (int, []byte)
	Input() *Input
	Output() *Output
	End()
}
