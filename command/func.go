package command

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func echo(cmd string, action interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Sprintf(" %s -- %s -- %v", now, cmd, action)
}

func catch(p interface{}) {
	buf := make([]byte, 64<<10)
	buf = buf[:runtime.Stack(buf, false)]
	fmt.Sprintf("%s -- cmd panic: %v  -- stack: %s\n", time.Now().Format("2006-01-02 15:04:05"), p, string(buf))
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
