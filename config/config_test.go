package config

import (
	"fmt"
	"github.com/rfyc/frame/ext/validator"
	"testing"
)

type address struct {
	Addr string
}
type user struct {
	Name    string
	Address []address
}

func TestConfig(t *testing.T) {

	var Conf = &Config{
		Config:   "app.json",
		LoadFile: LoadFile,
	}
	var app = user{}
	fmt.Println(validator.Validate(Conf.Rules()))
	fmt.Println("content :", string(Conf.content))
	fmt.Println("prepare :", Conf.Prepare(&app))
	fmt.Println("app     :", app)
}
