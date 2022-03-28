package core

import (
	"os"
	"testing"
)

func TestCmd(t *testing.T) {

	os.Args = []string{"frame", "start", "--config=app.json"}
	Command.Run(&App{})
}
