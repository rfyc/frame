package core

import "github.com/rfyc/frame/ext/validator"

var Config = &config{}

type config struct {
	Config string
}

func (this *config) Rules() validator.IRules {

	return validator.IRules{
		&validator.File{
			Names:  "config",
			Struct: this,
		},
	}
}

func (this *config) Prepare(app interface{}) error {

	if this.Config != "" {

	}
	return nil
}

func (this *config) String() string {
	return this.Config
}

func (this *config) DefaultFile() string {

	return "/config/app.json"

}
