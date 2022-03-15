package command

import "os"

type cmd struct {
	cmds map[string]func() error
	desc map[string]string
}

func (this *cmd) init() *cmd {
	this.cmds = make(map[string]func() error)
	this.desc = make(map[string]string)
	return this
}

func (this *cmd) register(name string, runCmd func() error, desc ...string) {

	this.cmds[name] = runCmd
	if len(desc) > 0 {
		this.desc["cmd_"+name] = desc[0]
	}
}

func (this *cmd) parse() (name string, runCmd func() error) {

	if len(os.Args) > 1 {
		name = os.Args[1]
		if runCmd = this.cmds[name]; runCmd != nil {
			return
		}
	}
	return
}
