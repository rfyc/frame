package command

import (
	"errors"
	"fmt"
	"github.com/rfyc/frame/frame/func/object"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
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

	if err := runCmd(); !errors.Is(err, errNoCmd) {
		return err
	}

	if err := runApp(); err != nil {
		return err
	}

	return nil
}

func runCmd() error {
	//从commands中找cmd执行
	if cmd := commands[strings.ToLower(cmdname)]; cmd != nil {
		if execCmd, ok := cmd.(IRunCmd); ok {
			execCmd.Construct()
			if err := object.Set(execCmd, os.Args); err != nil {
				return fmt.Errorf("cmd set Args error: %w", err)
			}
			if err := execCmd.Init(); err != nil {
				return fmt.Errorf("cmd Init error: %w", err)
			}
			wait := make(chan bool, 1)
			go func() {
				execCmd.Run()
				execCmd.Destruct()
				wait <- true
			}()
			select {
			case <-ctx.Done():
			case <-wait:
			}
			return nil
		}
	}
	return errNoCmd
}

func runApp() error {

	//从app中找函数执行
	if method := object.FindMethod(execApp, cmdname); method != "" {
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
