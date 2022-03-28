package structs

import (
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strings"
)

//传入非struct参数 会panic
//return map[ToLower(field)]field
func Fields(argv interface{}) map[string]string {
	refValue := ValueOf(argv)
	if refValue.Kind() != reflect.Struct {
		panic("structs::Fields argv not struct")
	}
	names := map[string]string{}
	refType := refValue.Type()
	for k := refType.NumField() - 1; k >= 0; k-- {
		if false == refType.Field(k).Anonymous {
			names[strings.ToLower(refType.Field(k).Name)] = refType.Field(k).Name
		}
	}
	return names
}

func IsPtr(argv interface{}) bool {
	if argv == nil {
		return false
	}
	return reflect.TypeOf(argv).Kind() == reflect.Ptr
}

func IsStruct(argv interface{}) bool {
	if argv == nil {
		return false
	}
	refType := reflect.TypeOf(argv)
	fmt.Println(refType.Kind())
	if refType.Kind() == reflect.Ptr {
		for {
			if refType.Kind() != reflect.Ptr {
				break
			}
			refType = refType.Elem()
			fmt.Println(refType.Kind())
		}
	}
	return refType.Kind() == reflect.Struct
}

/**
 * argv 传入struct 会得到一个struct 可以获取struct的field 不能调用指针的函数
 * argv 如果直接传 nil 会得到一个invalid的 Value 只能使用 Value.Kind()  Value.IsValid() 等函数
 */
func ValueOf(argv interface{}) reflect.Value {

	refValue := reflect.ValueOf(argv)
	for {
		if refValue.Kind() != reflect.Ptr {
			break
		}
		refValue = refValue.Elem()
	}
	return refValue
}

/**
 * argv 直接传nil 会panic
 * struct 调用call method 必须用 PtrOf().Method().Call()
 * struct 调用field会panic  PtrOf().Field()
 * struct set field 需要用 PtrOf().Elem().Field().Set()
 */
func PtrOf(argv interface{}) reflect.Value {

	refValue := reflect.ValueOf(argv)
	if refValue.Kind() != reflect.Ptr {
		panic("structs::PtrOf argv not ptr")
	}
	for {
		if refValue.Elem().Kind() != reflect.Ptr {
			break
		}
		refValue = refValue.Elem()
	}
	return refValue
}

func Set(dst interface{}, src interface{}) error {

	if !IsPtr(dst) {
		return errors.New("dst error: not ptr")
	}
	value := ValueOf(src)
	if !value.IsValid() || !value.CanInterface() {
		return errors.New("src error: invalid")
	}
	// []byte string
	switch v := value.Interface().(type) {
	case []byte:
		if len(v) == 0 {
			return errors.New("src error: empty")
		}
		if err := jsoniter.Unmarshal(v, dst); err != nil {
			return fmt.Errorf("src error: json parse: %w", err)
		}
	case string:
		if len(v) == 0 {
			return errors.New("src error: empty")
		}
		if err := jsoniter.Unmarshal([]byte(v), dst); err != nil {
			return fmt.Errorf("src error: json parse: %w", err)
		}
	default: //map struct
		switch value.Kind() {
		case reflect.Map, reflect.Struct:
			if jbytes, err := jsoniter.Marshal(value.Interface()); err != nil {
				return fmt.Errorf("src error: json fail: %w", err)
			} else if err := jsoniter.Unmarshal(jbytes, dst); err != nil {
				return fmt.Errorf("src error: json parse: %w", err)
			}
		default:
			return errors.New("src error: type invalid")
		}
	}

	return nil
}

func Values(obj interface{}) map[string]interface{} {

	return values(ValueOf(obj))
}

func values(valueOf reflect.Value) map[string]interface{} {

	maps := make(map[string]interface{})

	switch valueOf.Kind() {

	case reflect.Map:
		if valueOf.CanInterface() {
			if jsonData, err := json.Marshal(valueOf.Interface()); err == nil {
				fmt.Println("json:", json.Unmarshal(jsonData, &maps))
			}
		}

	case reflect.Struct:
		for k := 0; k < valueOf.Type().NumField(); k++ {
			if valueOf.Type().Field(k).Anonymous {
				vals := values(valueOf.Field(k))
				for key, val := range vals {
					maps[key] = val
				}
			} else {
				if valueOf.Type().Field(k).IsExported() && valueOf.Field(k).CanInterface() {
					maps[strings.ToLower(valueOf.Type().Field(k).Name)] = valueOf.Field(k).Interface()
				}
			}
		}
	}
	return maps
}
