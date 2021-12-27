package web

type IApp interface {
	Prepare() error
	Start() error
	Stop() error
}

type App struct {
}

func (this *App) Start() error {
	return nil
}

func (this *App) Stop() error {
	return nil
}
