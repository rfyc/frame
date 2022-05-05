package core

import (
	"github.com/rfyc/frame/command"
	"github.com/rfyc/frame/config"
	"github.com/rfyc/frame/route"
)

var Conf = &config.Config{}

var Command = &command.Command{}

var Router = &route.Router{
	Regsiter: &route.DefaultRegister{},
	HTTP:     &route.DefaultHTTP{},
}

func init() {
	Command.Init()
	Command.RegisterArgs("config", Conf)
	Router.Regsiter.Controller(&Controller{})
}
