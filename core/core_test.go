package core

import (
	"os"
	"testing"
)

func TestCmd(t *testing.T) {

	os.Args = []string{"frame", "started", "--config=app.conf"}
	Command.Run(&App{})
}
