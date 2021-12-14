package core

import (
	"strings"
)

type Runner struct {
	Config string
	Logs   string
}

func (this *Runner) CommandRun(execApp iApp) {

}

func (this *Runner) parseArgs(args []string) map[string]string {

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
