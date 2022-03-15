package command

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type iApp interface {
	Prepare() error
	Start() error
	Stop() error
	Restart() error
}

type command struct {
	input  *input
	action *action
	cmd    *cmd
	args   *args
	stdio  *stdio
	done   chan bool
	signal chan os.Signal
	ctx    context.Context
	cancel context.CancelFunc
}

func (this *command) Init() {
	this.input = &input{}
	this.action = (&action{}).init()
	this.cmd = (&cmd{}).init()
	this.args = (&args{}).init()
	this.done = make(chan bool, 1)
	this.signal = make(chan os.Signal, 1)
	this.ctx, this.cancel = context.WithCancel(context.Background())
}

func (this *command) Run(app iApp) {

	//信号量绑定
	signal.Notify(this.signal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	//捕获异常
	defer func() {
		if p := recover(); p != nil {
			this.stdio.format("main", "recover").echo()
			this.catch(p)
			this.stdio.format("catch").format(fmt.Sprintf("%v", p)).echo()
			this.cancel()
			this.stdio.format("main", "wait").echo()
		}
	}()

	//执行cmd
	go func() {
		defer func() {
			this.done <- true
		}()
		this.run(app)
	}()

	//信号捕获
	for {
		select {
		case <-this.signal:
			this.stdio.format("main", "stop").echo()
			this.cancel()
			this.stdio.format("main", "wait").echo()
		case <-this.done:
			this.stdio.format("main", "done").echo()
			return
		}
	}

}

func (this *command) run(app iApp) {

	//init input
	if err := this.input.json(); err != nil {
		this.stdio.format("input", "json").format("error", err.Error()).echo()
		return
	}
	//service cmds
	if app != nil {
		this.RegisterCmd("start", app.Start)
		this.RegisterCmd("stop", app.Stop)
		this.RegisterCmd("restart", app.Restart)
	}
	//init args
	for name, Args := range this.args.maps {
		//bind args
		if err := json.Unmarshal(this.input.bytes, &Args); err != nil {
			this.stdio.format("cmd", "args").format("name", name).format("bind").format("error", err.Error()).echo()
			return
		}
		//args prepare
		if err := Args.Prepare(app); err != nil {
			this.stdio.format("cmd", "args").format("name", name).format("prepare").format("error", err.Error()).echo()
			return
		}
	}
	//parse cmd
	if name, cmd := this.cmd.parse(); cmd != nil {
		this.stdio.format("cmd", name).format("running").echo()
		if err := cmd(); err != nil {
			this.stdio.format("cmd", name).format("run", "error", err.Error()).echo()
			return
		}
		this.stdio.format("cmd", name).format("over").echo()
		return
	}
	//parse action
	if name, action := this.action.parse(); action != nil {
		//prepare
		this.stdio.format("action", name).format("parse", "ok").echo()
		if err := action.Prepare(); err != nil {
			this.stdio.format("action", name).format("prepare").format("error", err.Error()).echo()
			return
		}
		//run
		this.stdio.format("action", name).format("running").echo()
		action.Run()
		this.stdio.format("action", name).format("runned").echo()
		action.End()
		this.stdio.format("action", name).format("end").echo()
		return
	}
	//not found
	this.stdio.format("cmd").format("error", "not found").echo()
}

func (this *command) RegisterArgs(name string, args iArgs, desc ...string) *command {
	this.args.register(name, args, desc...)
	return this
}

func (this *command) RegisterCmd(name string, runCmd func() error, desc ...string) *command {
	this.cmd.register(name, runCmd, desc...)
	return this
}

func (this *command) RegisterAction(name string, action iAction, desc ...string) *command {
	this.action.register(name, action, desc...)
	return this
}

func (this *command) catch(p interface{}) {
	buf := make([]byte, 64<<10)
	buf = buf[:runtime.Stack(buf, false)]
	fmt.Sprintf("%s -- cmd panic: %v  -- stack: %s\n", time.Now().Format("2006-01-02 15:04:05"), p, string(buf))
}
