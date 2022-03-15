package command

import "os"

type iAction interface {
	Prepare() error
	Run()
	End()
}

type action struct {
	actions map[string]iAction
	desc    map[string]string
}

func (this *action) init() *action {
	this.actions = make(map[string]iAction)
	this.desc = make(map[string]string)
	return this
}

func (this *action) register(name string, action iAction, desc ...string) {

	this.actions[name] = action
	if len(desc) > 0 {
		this.desc[name] = desc[0]
	}
}

func (this *action) parse() (name string, action iAction) {

	if len(os.Args) > 1 {
		name = os.Args[1]
		if action = this.actions[name]; action != nil {
			return
		}
		return
	}
	return
}
