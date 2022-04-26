package core

import (
	"context"
	"github.com/rfyc/frame/route"
	"github.com/rfyc/frame/utils/structs"
)

type Controller struct {
	ctx    context.Context
	input  *route.Input
	output *route.Output
}

func (this *Controller) Ctx(ctx ...context.Context) context.Context {
	if len(ctx) > 0 {
		this.ctx = ctx[0]
	}
	return this.ctx
}

func (this *Controller) Init() {
	this.ctx = context.Background()
	this.input = &route.Input{}
	this.output = &route.Output{}
}

func (this *Controller) Prepare(action route.IAction) error {
	if err := structs.Set(action, this.In().Request); err != nil {
		return err
	}
	return nil
}

func (this *Controller) Run(action route.IAction) {

}

func (this *Controller) In() *route.Input {

	return this.input
}

func (this *Controller) Out() *route.Output {
	return this.output
}
