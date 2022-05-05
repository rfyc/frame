package config

import (
	"bytes"
	"errors"
	"github.com/rfyc/frame/ext/validator"
	"github.com/rfyc/frame/utils/conv"
	"github.com/rfyc/frame/utils/structs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Config  string
	content []byte
}

func (this *Config) Rules() validator.IRules {

	return validator.IRules{
		&validator.File{
			Names:  "config",
			Struct: this,
		},
		&validator.Method{
			Func: this.Load,
		},
	}
}

func (this *Config) Load() (err error) {

	if this.Config != "" {
		if this.content, err = LoadFile(this.Config); err != nil {
			return err
		}
	}
	return nil
}

func (this *Config) Prepare(app interface{}) error {

	if len(this.content) > 0 {
		return structs.Set(app, this.content)
	}
	return nil
}

func (this *Config) String() string {
	return this.Config
}

func LoadFile(config string) ([]byte, error) {

	var (
		err      error
		dirpath  string
		scontent = []byte{}
		smaps    = make(map[string][]byte)
		bcontent = []byte{}
		bmaps    = make(map[string]interface{})
	)
	if strings.Index(config, ".json") <= 0 {
		return bcontent, errors.New("file not json")
	}
	if bcontent, err = ioutil.ReadFile(config); err != nil {
		return bcontent, errors.New("config read error: " + err.Error())
	}
	if err = structs.Set(&bmaps, bcontent); err != nil {
		return bcontent, errors.New("config set error: " + err.Error())
	}
	if dirpath, err = filepath.Abs(filepath.Dir(config)); err != nil {
		return bcontent, errors.New("filepath error: " + err.Error())
	}
	for key := range bmaps {
		var file = dirpath + "/" + conv.String(bmaps[key])
		if strings.Index(file, ".json") <= 0 {
			continue
		}
		if finfo, err := os.Stat(file); err != nil || finfo.IsDir() {
			return bcontent, err
		}
		if scontent, err = ioutil.ReadFile(file); err != nil {
			return bcontent, err
		}
		rkey := `"` + conv.String(bmaps[key]) + `"`
		smaps[rkey] = scontent
	}
	for key, value := range smaps {
		bcontent = bytes.Replace(bcontent, []byte(key), value, -1)
	}
	return bcontent, nil
}
