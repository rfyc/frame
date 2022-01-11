package command

import (
	"os"
	"strings"
)

var (
	execApp  IRunApp
	commands = map[string]*runCmd{}
	actions  = map[string]map[string]*runAction{}
)

type IRunApp interface {
	Construct()
	Init() error
	Start() error
	Stop() error
	Destruct()
}

type IRunCmd interface {
	Construct()
	Init() error
	Run(IRunAction) error
	Stop() error
	Destruct()
}

type IRunAction interface {
	Construct()
	Init() error
	Run() error
	Stop() error
	Destruct()
}

type runAction struct {
	execAction IRunAction
	desc       string
}

type runCmd struct {
	execCmd IRunCmd
	desc    string
}

func findCmd() IRunCmd {
	if len(os.Args) > 1 {
		nameCmd = strings.ToLower(os.Args[1])
		if runCmd := commands[nameCmd]; runCmd != nil && runCmd.execCmd != nil {
			return runCmd.execCmd
		}
	}
	return nil
}

func findAction() IRunAction {
	if len(os.Args) > 2 {
		nameCmd = strings.ToLower(os.Args[1])
		nameAction = strings.ToLower(os.Args[2])
		if runActions := actions[nameCmd]; len(runActions) > 0 {
			if runAction := actions[nameCmd][nameAction]; runAction != nil && runAction.execAction != nil {
				return runAction.execAction
			}
		}
	}
	return nil
}
func RegisterAction(cmdName, actionName string, execAction IRunAction, desc ...string) {

	cmdName = strings.ToLower(cmdName)
	actionName = strings.ToLower(actionName)

	if len(actions[cmdName]) == 0 {
		actions[cmdName] = map[string]*runAction{}
	}

	actions[cmdName][actionName] = &runAction{
		execAction: execAction,
		desc:       strings.Join(desc, ","),
	}
}

func RegisterCmd(cmdName string, execCmd IRunCmd, desc ...string) {

	cmdName = strings.ToLower(cmdName)
	commands[cmdName] = &runCmd{
		execCmd: execCmd,
		desc:    strings.Join(desc, ","),
	}
}
