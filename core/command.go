package core

import (
	"github.com/rfyc/frame/command"
)

var Command = &command.Command{}

func init() {
	Command.Init()
	Command.RegisterArgs("config", Config)
}
