package command

import "strings"

var (
	execApp  IRunApp
	commands = map[string]cmdOrAction{}
)

type cmdOrAction interface{}

type action struct {
	runAction IRunAction
	desc      string
}

type cmd struct {
	runCmd  IRunCmd
	actions map[string]*action
	desc    string
}

func (this *cmd) findAction(name string) IRunAction {
	if act := this.actions[strings.ToLower(name)]; act != nil {
		return act.runAction
	}
	return nil
}

func RegisterAction(cmdName string, runAction IRunAction, desc ...string) {

	act := &action{runAction: runAction}
	if len(desc) > 0 {
		act.desc = desc[0]
	}
	commands[cmdName] = act
}

func RegisterCmd(cmdName string, runCmd IRunCmd, desc ...string) {

	c := &cmd{runCmd: runCmd}
	if len(desc) > 0 {
		c.desc = desc[0]
	}
	if cc, ok := commands[cmdName].(*cmd); ok {
		c.actions = cc.actions
	}
	commands[cmdName] = c
}

func RegisterCmdAction(cmdName string, actionName string, runAction IRunAction, desc ...string) {

	act := &action{runAction: runAction}
	if len(desc) > 0 {
		act.desc = desc[0]
	}
	c := &cmd{}
	if cc, ok := commands[cmdName].(*cmd); ok {
		c = cc
	}
	c.actions[actionName] = act
}
