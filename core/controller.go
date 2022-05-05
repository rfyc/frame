package core

import (
	"context"
	"github.com/rfyc/frame/route"
	"github.com/rfyc/frame/utils/structs"
)

type Controller struct {
	ctx    context.Context
	action route.IAction
	input  *route.Input
}

func (this *Controller) Ctx() context.Context {
	return this.ctx
}

func (this *Controller) Init(ctx context.Context, in *route.Input, action route.IAction) {
	this.ctx = ctx
	this.input = in
	this.action = action
}

func (this *Controller) In() *route.Input {

	return this.input
}

func (this *Controller) Action() route.IAction {
	return this.action
}

func (this *Controller) Prepare() error {

	if err := structs.Set(this.action, this.In().Request); err != nil {
		return err
	}

	if err := this.action.Prepare(); err != nil {
		return err
	}

	return nil
}

func (this *Controller) Out(err error, content interface{}) *route.Output {
	return &route.Output{}
}
