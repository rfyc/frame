package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/rfyc/frame/utils/conv"
	"github.com/rfyc/frame/utils/object"
)

var (
	done        = make(chan bool, 1)
	stopSig     = make(chan os.Signal, 1)
	ctx, cancel = context.WithCancel(context.Background())
)

func RunApp(app IRunApp) {
	execApp = app
	Run()
}

func Run() {

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
		defer func() {
			done <- true
		}()
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

func initApp() error {

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
	return nil
}

func run() {

	if err := initApp(); err != nil {
		return
	}

	execCmd := findCmd()
	if execCmd == nil {
		cmdHelper()
		return
	}
	execAction := findAction()
	if execAction == nil {
		actionHelper()
		return
	}

	return
}

func cmdHelper() {

}

func actionHelper() {

}

func catch(p interface{}) {
	buf := make([]byte, 64<<10)
	buf = buf[:runtime.Stack(buf, false)]
	log := time.Now().Format("2006-01-02 15:04:05.000") + " -- " + "cmd panic:" + conv.String(p)
	log += " -- stack:" + string(buf)
	echo("catch", log)
}

func echo(cmd string, action interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now, fmt.Sprintf(" %-10s", cmd), action)
}

func parseArgs(args []string) map[string]string {

	maps := make(map[string]string)
	for _, value := range args {
		if strings.HasPrefix(value, "--") {
			params := strings.Split(value[2:], "=")
			count := len(params)
			if count > 1 {
				maps[params[0]] = params[1]
			} else if count == 1 {
				maps[params[0]] = "true"
			}
		}
	}
	return maps
}
