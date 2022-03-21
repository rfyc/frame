package core

import (
	"github.com/rfyc/frame/utils/object"
)

var Config = &config{}

type config struct {
	Config string
}

func (this *config) Rules() {

	validator.
}

func (this *config) Prepare(app interface{}) error {
	return object.LoadFile(app, this.Config)
}

func (this *config) String() string {
	return this.Config
}

func (this *config) DefaultFile() string {

	return "/config/app.json"

}
