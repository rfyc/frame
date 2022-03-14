package core

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var Command = &command{}

type iApp interface {
	Prepare() error
	Start() error
	Stop() error
}
type iCmd interface {
	Prepare() error
	Run()
	End()
}
type iArgs interface {
	String() string
	Prepare(app iApp) error
}

func init() {
	Command.init()
	Command.RegisterArgs("config", Config)
}

type stdio struct {
	out []interface{}
}

func (this *stdio) format(args ...string) *stdio {
	this.out = append(this.out, "["+strings.Join(args, ":")+"]")
	return this
}

func (this *stdio) echo() {
	out := []interface{}{"[" + time.Now().Format("2006:01:02:15:04:05.000") + "]"}
	fmt.Println(append(out, this.out...)...)
	this.out = []interface{}{}
}

type command struct {
	actions map[string]iCmd
	args    map[string]iArgs
	cmds    map[string]func() error
	desc    map[string]string
	input   []byte
	stdio   stdio
}

func (this *command) init() {

	this.actions = make(map[string]iCmd)
	this.args = make(map[string]iArgs)
	this.cmds = make(map[string]func() error)
	this.desc = make(map[string]string)
}

func (this *command) Run(app iApp) {

	//init error
	var err error
	//init input
	if this.input, err = json.Marshal(this.parseInput()); err != nil {
		this.stdio.format("input", "json").format("error", err.Error()).echo()
		return
	}
	//init args
	for name, Args := range this.initArgs() {
		//bind args
		if err = json.Unmarshal(this.input, &Args); err != nil {
			this.stdio.format("cmd", "args").format("name", name).format("bind").format("error", err.Error()).echo()
			return
		}
		//args prepare
		if err = Args.Prepare(app); err != nil {
			this.stdio.format("cmd", "args").format("name", name).format("prepare").format("error", err.Error()).echo()
			return
		}
		this.stdio.format("cmd", "args ").format(name, Args.String()).echo()
	}
	//run cmd
	this.RegisterCmd("start", app.Start)
	this.RegisterCmd("stop", app.Stop)

	if name, cmd, action := this.parseAction(); cmd != nil {
		this.stdio.format("cmd", name).format("running").echo()
		cmd()
		this.stdio.format("cmd", name).format("over").echo()
	} else if action != nil {
		this.stdio.format("action", name).format("prepare").echo()
		if err := action.Prepare(); err != nil {
			this.stdio.format("action", name).format("prepare").format("error", "not found").echo()
			return
		}
		this.stdio.format("action", name).format("running").echo()
		action.Run()
		this.stdio.format("action", name).format("end").echo()
		action.End()
		this.stdio.format("action", name).format("over").echo()

	} else {
		this.stdio.format("cmd", name).format("error", "not found").echo()
	}
}

func (this *command) RegisterArgs(name string, args iArgs, desc ...string) {

	this.args[name] = args
	if len(desc) > 0 {
		this.desc["args_"+name] = desc[0]
	}
}

func (this *command) RegisterAction(name string, action iCmd, desc ...string) {

	this.actions[name] = action
	if len(desc) > 0 {
		this.desc["action_"+name] = desc[0]
	}
}

func (this *command) RegisterCmd(name string, runCmd func() error, desc ...string) {

	this.cmds[name] = runCmd
	if len(desc) > 0 {
		this.desc["cmd_"+name] = desc[0]
	}
}

func (this *command) echo(args ...interface{}) {
}

func (this *command) initArgs() map[string]iArgs {
	return this.args
}

func (this *command) parseInput() map[string]string {

	this.stdio.format("cmd", "input")
	input := map[string]string{}
	count := len(os.Args)
	for k := 2; k < count; k++ {
		if strings.Contains(os.Args[k], "-") && strings.Contains(os.Args[k], "=") {
			args := strings.SplitN(strings.Trim(os.Args[k], "-"), "=", 2)
			input[args[0]] = args[1]
			this.stdio.format("--" + args[0] + "=" + args[1])
		}
	}
	this.stdio.echo()
	return input
}

func (this *command) parseAction() (name string, cmd func() error, action iCmd) {

	if len(os.Args) > 1 {
		name = os.Args[1]
		if cmd = this.cmds[name]; cmd != nil {
			return
		}
		if action = this.actions[name]; action != nil {
			return
		}
		return
	}
	return
}
