package command

import (
	"context"
	"fmt"
	"github.com/rfyc/frame/ext/validator"
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

type Command struct {
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

func (this *Command) Init() {
	this.input = &input{}
	this.action = (&action{}).init()
	this.cmd = (&cmd{}).init()
	this.args = (&args{}).init()
	this.stdio = &stdio{}
	this.done = make(chan bool, 1)
	this.signal = make(chan os.Signal, 1)
	this.ctx, this.cancel = context.WithCancel(context.Background())
}

func (this *Command) Run(servApp iApp) {

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
		this.run(servApp)
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

func (this *Command) run(servApp iApp) {

	//service cmds
	if servApp != nil {
		this.RegisterCmd("start", servApp.Start)
		this.RegisterCmd("stop", servApp.Stop)
		this.RegisterCmd("restart", servApp.Restart)
	}
	//init args
	for name, Args := range this.args.maps {
		//bind args
		if err := this.input.bind(Args); err != nil {
			this.stdio.format("error").format("args", name).format("bind", err.Error()).echo()
			return
		}
		//validate rules
		if ok, err := validator.Validate(Args.Rules()); !ok {
			this.stdio.format("error").format("args", name, "rules").format(err.Error()).echo()
			return
		}
		//args prepare
		if err := Args.Prepare(servApp); err != nil {
			this.stdio.format("error").format("args", name).format("prepare", err.Error()).echo()
			return
		}
		//args output
		this.stdio.format("args").format(name, Args.String()).echo()
	}
	//parse cmd
	if name, cmd := this.cmd.parse(); cmd != nil {
		this.stdio.format("cmd", name).format("running").echo()
		if err := cmd(); err != nil {
			this.stdio.format("cmd", name).format("run", "error", err.Error()).echo()
			return
		}
		return
	}
	//parse action
	if name, action := this.action.parse(); action != nil {
		//prepare
		this.stdio.format("action", name).format("parse").echo()
		if err := action.Prepare(); err != nil {
			this.stdio.format("action", name).format("prepare").format("error", err.Error()).echo()
			return
		}
		//run
		this.stdio.format("action", name).format("run...").echo()
		action.Run()
		action.End()
		return
	}
	//not found
	this.stdio.format("error").format("cmd not found").echo()
}

func (this *Command) RegisterArgs(name string, args iArgs, desc ...string) *Command {
	this.args.register(name, args, desc...)
	return this
}

func (this *Command) RegisterCmd(name string, runCmd func() error, desc ...string) *Command {
	this.cmd.register(name, runCmd, desc...)
	return this
}

func (this *Command) RegisterAction(name string, action iAction, desc ...string) *Command {
	this.action.register(name, action, desc...)
	return this
}

func (this *Command) catch(p interface{}) {
	buf := make([]byte, 64<<10)
	buf = buf[:runtime.Stack(buf, false)]
	fmt.Sprintf("%s -- cmd panic: %v  -- stack: %s\n", time.Now().Format("2006-01-02 15:04:05"), p, string(buf))
}
