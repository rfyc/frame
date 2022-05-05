package core

import (
	"fmt"
	"github.com/rfyc/frame/connect"
)

type App struct {
	HTTP connect.HTTPServer
}

func (this *App) Init() {
	this.HTTP.Init()
}

func (this *App) Prepare() error {
	return nil
}
func (this *App) Start() error {

	fmt.Printf("http: %+v\n", this.HTTP)
	this.HTTP.Handle("/", Router)
	fmt.Println(this.HTTP.Start())
	return nil
}

func (this *App) Restart() error {
	fmt.Println("app restart")
	return nil
}

func (this *App) Stop() error {
	return nil
}
