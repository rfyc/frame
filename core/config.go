package core

var Config = &config{}

type config struct {
	Config string
}

func (this *config) Prepare(app iApp) error {
	return nil
}

func (this *config) String() string {
	return this.Config
}
