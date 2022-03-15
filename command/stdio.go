package command

import (
	"fmt"
	"time"
)

type stdio struct {
	out []interface{}
}

func (this *stdio) format(args ...interface{}) *stdio {
	this.out = append(this.out, "")
	return this
}

func (this *stdio) echo() {
	out := []interface{}{"[" + time.Now().Format("2006:01:02:15:04:05.000") + "]"}
	fmt.Println(append(out, this.out...)...)
	this.out = []interface{}{}
}
