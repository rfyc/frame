package core

import "context"

type Action struct {
	ctx context.Context
}

func (this *Action) Init(ctx context.Context) {
	this.ctx = ctx
}

func (this *Action) Ctx() context.Context {
	return this.ctx
}

func (this *Action) Prepare() error {
	return nil
}

func (this *Action) Run() (err error, content interface{}) {
	return nil, nil
}
