package validator

import (
	"github.com/rfyc/frame/utils/file"
	"github.com/rfyc/frame/utils/structs"
	"reflect"
	"strings"
)

type File struct {
	Names    string
	Struct   interface{}
	Required bool
	Error    error
}

func (this *File) Validate() (bool, error) {

	names := strings.Trim(this.Names, ",")
	refValue := structs.ValueOf(this.Struct)
	field_names := structs.Fields(this.Struct)
	for _, name := range strings.Split(strings.ToLower(names), ",") {
		if field_names[name] != "" {
			Field := refValue.FieldByName(field_names[name])
			if Field.Kind() != reflect.String {
				return false, errorf(this.Error, "%s not string", name)
			}
			if this.Required && Field.String() == "" {
				return false, errorf(this.Error, "%s empty", name)
			}
			if false == file.IsFile(Field.String()) {
				return false, errorf(this.Error, "%s=%s not file", name, Field.String())
			}
		}
	}
	return true, nil
}
