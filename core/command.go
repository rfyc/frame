package core

import (
	"os"
	"strings"
)

var Command = &command{}

type iApp interface {
	Prepare() error
	Start() error
	Stop() error
}
type iAction interface {
}
type iArgs interface {
	Prepare(app iApp) error
}
type command struct {
	actions    map[string]iAction
	commonArgs map[string]iArgs
	cmds       map[string]func()
	desc       map[string]string
}

func (this *command) Run(app iApp) {

}

func (this *command) parseArgs() map[string]string {

	args := map[string]string{}

	return args
}

func (this *command) parseAction() (name string, action iAction) {
	if len(os.Args) > 1 {
		if false == strings.Contains(os.Args[1], "-") {
			name = os.Args[1]
		}
	}
	return
}

func (this *command) RegisterArgs(name string, args iArgs, desc ...string) {
	if this.commonArgs == nil {
		this.commonArgs = map[string]iArgs{}
	}
	this.commonArgs[name] = args
	if len(desc) > 0 {
		this.desc["args_"+name] = desc[0]
	}
}

func (this *command) RegisterAction(name string, action iAction, desc ...string) {
	if this.actions == nil {
		this.actions = map[string]iAction{}
	}
	this.actions[name] = action
	if len(desc) > 0 {
		this.desc["action_"+name] = desc[0]
	}
}

func (this *command) RegisterCmd(name string, runCmd func(), desc ...string) {
	if this.cmds == nil {
		this.cmds = map[string]func(){}
	}
	this.cmds[name] = runCmd
	if len(desc) > 0 {
		this.desc["cmd_"+name] = desc[0]
	}
}
