package command

type iArgs interface {
	Prepare(app iApp) error
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

func (this *args) register(name string, args iArgs, desc ...string) {
	this.maps[name] = args
	if len(desc) > 0 {
		this.desc["args_"+name] = desc[0]
	}
}
