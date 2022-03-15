package command

import (
	"encoding/json"
	"os"
	"strings"
)

type input struct {
	bytes []byte
	maps  map[string]string
}

func (this *input) parse() map[string]string {

	this.maps = make(map[string]string)
	count := len(os.Args)
	for k := 2; k < count; k++ {
		if strings.Contains(os.Args[k], "-") && strings.Contains(os.Args[k], "=") {
			args := strings.SplitN(strings.Trim(os.Args[k], "-"), "=", 2)
			this.maps[args[0]] = args[1]
		}
	}

	return this.maps
}
func (this *input) json() (err error) {
	this.bytes, err = json.Marshal(this.parse())
	return
}
