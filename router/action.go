package router

import "github.com/rfyc/frame/ctx"

type IAction interface {
	Init()
	Ctx() *ctx.Ctx
	Prepare() error
	Run() (error, interface{})
	End()
}
