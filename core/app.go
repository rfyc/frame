package core

import "fmt"

type App struct {
}

func (this *App) Prepare() error {
	fmt.Println("app prepare")
	return nil
}
func (this *App) Start() error {
	fmt.Println("app start")
	return nil
}

func (this *App) Stop() error {
	return nil
}
