package validator

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	Name *string
	Age  uint8
}

func (this *user) Echo() {
	fmt.Println(*this.Name)
	fmt.Println(2222)
}

func TestRule(t *testing.T) {

	var name = "hong"
	var a = &user{&name, 22}

	return
	ref := reflect.Indirect(reflect.ValueOf(&a))
	ref.FieldByName("Age").SetUint(123)

	ref = reflect.Indirect(reflect.ValueOf(a))
	//fmt.Println(ref.FieldByName("Name").Elem().String())
	//fmt.Println(ref.FieldByName("Name").Kind().String())
	fmt.Println(ref.FieldByName("Age").Uint())
	//fmt.Println(ref.FieldByName("Age").Kind().String())
	//
	ref = reflect.ValueOf(&a)
	ref.MethodByName("Echo").Call([]reflect.Value{})
}
