package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/rfyc/frame/utils/conv"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	nameCmd     = "start"
	nameAction  = ""
	args        = parseArgs(os.Args)
	done        = make(chan bool, 1)
	stopSig     = make(chan os.Signal, 1)
	ctx, cancel = context.WithCancel(context.Background())
	errNoCmd    = errors.New("no cmd")
)

func init() {
	if len(os.Args) > 1 {
		nameCmd = strings.ToLower(os.Args[1])
	}
	if len(os.Args) > 2 {
		nameAction = strings.ToLower(os.Args[2])
	}
}

type IRunApp interface {
	Construct()
	Init() error
	Start() error
	Stop() error
	Destruct()
}

type IRunCmd interface {
	Construct()
	Init() error
	Run() error
	Destruct()
}

type IRunAction interface {
	Construct()
	Init() error
	Run() error
	Destruct()
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
