package core

type iApp interface {
	Start()
	Restart()
	Stop()
}

type App struct {
}
