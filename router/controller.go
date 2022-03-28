package router

import "github.com/rfyc/frame/ctx"

type IController interface {
	Init()
	Ctx() *ctx.Ctx
	Prepare() error
	RunAction(runFunc func() (error, interface{})) (int, []byte)
	End()
}
