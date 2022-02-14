package command

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type IAction interface {
	Construct()
	Init() error
	Run() error
	Stop() error
	Destruct()
}

type IArgs interface {
	Construct()
	Init() error
	Run() error
	Destruct()
}

type Command struct {
	done     chan bool
	signal   chan os.Signal
	ctx      context.Context
	cancel   context.CancelFunc
	args     []IArgs
	actions  map[string]map[string]IAction
	commands map[string]func()
}

func (this *Command) RegisterArgs(runArgs IArgs) *Command {

	return this
}

func (this *Command) RegisterCmd(cmd string, runCmd func()) *Command {

	if this.commands == nil {
		this.commands = make(map[string]func())
	}

	return this
}

func (this *Command) RegisterAction(cmd, action string, runAction IAction) *Command {

	if this.actions == nil {
		this.actions = make(map[string]map[string]IAction)
	}

	return this
}

func (this *Command) Run() {

	this.done = make(chan bool, 1)
	this.signal = make(chan os.Signal, 1)
	this.ctx, this.cancel = context.WithCancel(context.Background())

	//信号量绑定
	signal.Notify(this.signal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	//捕获异常
	defer func() {
		if p := recover(); p != nil {
			echo("main", "recover")
			catch(p)
			echo("catch", p)
			this.cancel()
			echo("main", "wait")
		}
	}()

	//执行cmd
	go func() {
		defer func() {
			this.done <- true
		}()
		this.run()
	}()

	//信号捕获
	for {
		select {
		case <-this.signal:
			echo("main", "stop")
			this.cancel()
			echo("main", "wait")
		case <-this.done:
			echo("main", "done")
			return
		}
	}
}

func (this *Command) run() {

}
