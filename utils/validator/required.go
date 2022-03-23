package validator

import (
	"fmt"
	"github.com/rfyc/frame/utils/structs"
	"strings"
)

type Required struct {
	Struct  interface{}
	Names   string
	IsEmpty bool
	Error   error
}

func (this *Required) Validate() (bool, error) {

	names := strings.Trim(this.Names, ",")
	refValue := structs.ValueOf(this.Struct)
	field_names := structs.Fields(this.Struct)
	for _, name := range strings.Split(strings.ToLower(names), ",") {
		if field_names[name] != "" {
			Field := refValue.FieldByName(field_names[name])
			if v, ok := Field.Interface().(string); ok && v == "" {
				return false, fmt.Errorf("%w: %s required", this.Error, name)
			}
			if this.IsEmpty && Field.IsZero() {
				return false, fmt.Errorf("%w: %s empty", this.Error, name)
			}
		}
	}
	return true, nil
}
