package command

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"

	"github.com/rfyc/frame/utils/object"
)

func Run(app ...IRunApp) {

	//信号量绑定
	signal.Notify(stopSig, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	//捕获异常
	defer func() {
		if p := recover(); p != nil {
			echo("main", "recover")
			catch(p)
			echo("catch", p)
			cancel()
			echo("main", "wait")
		}
	}()

	//执行cmd
	go func() {
		registerApp(app...)
		run()
	}()

	//信号捕获
	for {
		select {
		case <-stopSig:
			echo("main", "stop")
			cancel()
			echo("main", "wait")
		case <-done:
			echo("main", "done")
			return

		}
	}
}

func run() error {

	if execApp != nil {
		execApp.Construct()
		defer execApp.Destruct()
		if err := object.Set(execApp, os.Args); err != nil {
			return err
		}
		if err := execApp.Init(); err != nil {
			return err
		}
	}

	if regCmd := commands[strings.ToLower(nameCmd)]; regCmd != nil {
		if execAction, ok := regCmd.(*action); ok {
			return runAction(execAction.runAction)
		}
		if execCmd, ok := regCmd.(*cmd); ok {
			if execAction := execCmd.findAction(nameAction); execAction != nil {
				return runAction(execAction)
			}
		}
	}

	return runApp()
}

func runAction(execAction IRunAction) error {

	execAction.Construct()
	if err := object.Set(execAction, os.Args); err != nil {
		return fmt.Errorf("cmd set Args error: %w", err)
	}
	if err := execAction.Init(); err != nil {
		return fmt.Errorf("cmd Init error: %w", err)
	}
	wait := make(chan bool, 1)
	go func() {
		execAction.Run()
		execAction.Destruct()
		wait <- true
	}()
	select {
	case <-ctx.Done():
	case <-wait:
	}
	return nil
}

func runApp() error {

	//从app中找函数执行
	if method := object.FindMethod(execApp, nameCmd); method != "" {
		wait := make(chan bool, 1)
		go func() {
			run := reflect.ValueOf(execApp).MethodByName(method)
			run.Call([]reflect.Value{})
			wait <- true
		}()
		select {
		case <-ctx.Done():
		case <-wait:
		}
		return nil
	}
	return errNoCmd
}
