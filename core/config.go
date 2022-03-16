package core

var Config = &config{}

type config struct {
	Config string
}

func (this *config) Prepare(app interface{}) error {
	return nil
}

func (this *config) String() string {
	return this.Config
}
