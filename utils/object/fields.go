package object

import (
	"reflect"
	"strings"
)

var Fields = &fields{}

type fields struct{}

func (this *fields) Names(obj interface{}) map[string]string {

	names := map[string]string{}
	if obj == nil {
		return names
	}
	refType := reflect.TypeOf(obj)
	if refType.Kind() == reflect.Ptr {
		for {
			refType = refType.Elem()
			if refType.Kind() != reflect.Ptr {
				break
			}
		}
	}
	if refType.Kind() == reflect.Struct {
		for k := refType.NumField() - 1; k >= 0; k-- {
			if false == refType.Field(k).Anonymous {
				names[strings.ToLower(refType.Field(k).Name)] = refType.Field(k).Name
			}
		}
	}
	return names
}

func (this *fields) Find(obj interface{}, name string) (bool, interface{}) {

	if !IsStruct(obj) {
		return false, nil
	}
	refValue := reflect.ValueOf(obj)
	if refValue.Kind() == reflect.Ptr {
		for {
			refValue = refValue.Elem()
			if refValue.Kind() == reflect.Struct {
				break
			}
		}
	}
	name = this.Names(obj)[strings.ToLower(name)]
	if refField := refValue.FieldByName(name); refField.IsValid() {
		return true, refField.Interface()
	}
	return false, nil
}
