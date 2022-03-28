package core

import (
	"github.com/rfyc/frame/command"
	"github.com/rfyc/frame/config"
)

var Conf = &config.Config{
	LoadFile: config.LoadFile,
}

var Command = &command.Command{}

func init() {
	Command.Init()
	Command.RegisterArgs("config", Conf)
}
