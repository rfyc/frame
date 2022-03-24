package validator

import (
	"github.com/rfyc/frame/utils/structs"
	"strings"
)

type Required struct {
	Struct interface{}
	Names  string
	Error  error
}

func (this *Required) Validate() (bool, error) {

	names := strings.Trim(this.Names, ",")
	refValue := structs.ValueOf(this.Struct)
	field_names := structs.Fields(this.Struct)
	for _, name := range strings.Split(strings.ToLower(names), ",") {
		if field_names[name] != "" {
			Field := refValue.FieldByName(field_names[name])
			if v, ok := Field.Interface().(string); ok && v == "" {
				return false, errorf(this.Error, "%s not required", name)
			}
			if Field.IsZero() {
				return false, errorf(this.Error, "%s empty", name)
			}
		}
	}
	return true, nil
}
