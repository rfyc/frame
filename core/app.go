package core

import "fmt"

type App struct {
	HTTP string
}

func (this *App) Prepare() error {
	fmt.Println("app prepare")
	return nil
}

func (this *App) Start() error {
	fmt.Println("app start")
	return nil
}

func (this *App) Restart() error {
	fmt.Println("app restart")
	return nil
}

func (this *App) Stop() error {
	return nil
}
