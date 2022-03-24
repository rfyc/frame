package command

import "github.com/rfyc/frame/ext/validator"

type iArgs interface {
	Prepare(app interface{}) error
	Rules() validator.IRules
	String() string
}

type args struct {
	maps map[string]iArgs
	desc map[string]string
}

func (this *args) init() *args {
	this.maps = make(map[string]iArgs)
	this.desc = make(map[string]string)
	return this
}

func (this *args) register(name string, argv iArgs, desc ...string) {
	this.maps[name] = argv
	if len(desc) > 0 {
		this.desc["args_"+name] = desc[0]
	}
}
