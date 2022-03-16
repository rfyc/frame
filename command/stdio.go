package command

import (
	"fmt"
	"strings"
	"time"
)

type stdio struct {
	out []interface{}
}

func (this *stdio) format(args ...interface{}) *stdio {

	if len(args) == 0 {
		return this
	}
	var data interface{}
	fmtStr := "["
	for _, arg := range args {
		fmtStr += fmt.Sprintf("%v:", arg)
	}
	data = strings.Trim(fmtStr, ":") + "]"
	this.out = append(this.out, data)
	return this
}

func (this *stdio) echo() {
	out := []interface{}{"[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"}
	fmt.Println(append(out, this.out...)...)
	this.out = []interface{}{}
}
