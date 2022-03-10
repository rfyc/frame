package core

import (
	"encoding/json"
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
	actions map[string]iAction
	args    map[string]iArgs
	input   map[string]string
	cmds    map[string]func()
	desc    map[string]string
}

func (this *command) init() {

	if this.actions == nil {
		this.actions = map[string]iAction{}
	}
	if this.args == nil {
		this.args = map[string]iArgs{}
	}
	if this.input == nil {
		this.input = map[string]string{}
	}
	if this.cmds == nil {
		this.cmds = map[string]func(){}
	}
	if this.desc == nil {
		this.desc = map[string]string{}
	}
}

func (this *command) RegisterArgs(name string, args iArgs, desc ...string) {
	this.init()
	this.args[name] = args
	if len(desc) > 0 {
		this.desc["args_"+name] = desc[0]
	}
}

func (this *command) RegisterAction(name string, action iAction, desc ...string) {
	this.init()
	this.actions[name] = action
	if len(desc) > 0 {
		this.desc["action_"+name] = desc[0]
	}
}

func (this *command) RegisterCmd(name string, runCmd func(), desc ...string) {
	this.init()
	this.cmds[name] = runCmd
	if len(desc) > 0 {
		this.desc["cmd_"+name] = desc[0]
	}
}

func (this *command) Run(app iApp) {

	this.init()

	for _, arg := range this.initArgs() {
		if input, err := json.Marshal(this.parseInput()); err != nil {

		} else if err := json.Unmarshal(input, &arg); err != nil {

		} else if err := arg.Prepare(app); err != nil {

		}
	}

}

func (this *command) initArgs() map[string]iArgs {
	if this.args == nil {
		this.args = map[string]iArgs{}
	}
	return this.args
}

func (this *command) parseInput() map[string]string {

	if this.input == nil {
		this.input = map[string]string{}
		count := len(os.Args)
		for k := 2; k < count; k++ {
			if strings.Contains(os.Args[k], "-") && strings.Contains(os.Args[k], "=") {
				arg := strings.SplitN(strings.Trim(os.Args[k], "-"), "=", 2)
				this.input[arg[0]] = arg[1]
			}
		}
	}
	return this.input
}

func (this *command) parseAction() (name string, action iAction) {

	if len(os.Args) > 1 {
		if false == strings.Contains(os.Args[1], "-") {
			name = os.Args[1]
		}
	}
	return
}
