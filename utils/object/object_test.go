package object

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	Name string
}

func (this *user) Set(name string) {
	this.Name = name
}

func (this *user) Echo() {
	fmt.Println("Name:", this.Name)
}

func TestMethod(t *testing.T) {
	var u = user{Name: "100"}
	reflect.ValueOf(&u).MethodByName("Echo").Call([]reflect.Value{})
	reflect.ValueOf(&u).MethodByName("Set").Call([]reflect.Value{reflect.ValueOf(123)})
	reflect.ValueOf(&u).MethodByName("Echo").Call([]reflect.Value{})

}

func TestFields(t *testing.T) {
	var u = user{Name: "100"}
	var ptr_u = &u
	var ptr_uu = &ptr_u
	var ptr_uuu = &ptr_uu
	fmt.Println(Fields.Find(ptr_uuu, "Name"))
	fmt.Println(Fields.Find(ptr_uuu, "name"))
}

func TestIsStruct(t *testing.T) {
	var u user
	var uv = user{}
	fmt.Println("user=nil   :", IsStruct(u))
	fmt.Println("user={}    :", IsStruct(uv))
	var ptr_u *user
	fmt.Println("*user=nil  :", IsStruct(ptr_u))
	fmt.Println("**user=nil :", IsStruct(&ptr_u))
	ptr_u = &user{}
	fmt.Println("*user={}   :", IsStruct(ptr_u))
	fmt.Println("**user={}  :", IsStruct(&ptr_u))
	ptr_uu := &ptr_u
	fmt.Println("***user={} :", IsStruct(&ptr_uu))
	ptr_uu = nil
	fmt.Println("***user=nil:", IsStruct(&ptr_uu))

	fmt.Println("nil        :", IsStruct(nil))
	fmt.Println("1          :", IsStruct(1))
	fmt.Println("ok         :", IsStruct("ok"))
	fmt.Println("true       :", IsStruct(true))
}

func TestIsPtr(t *testing.T) {
	var u user
	var uv = user{}
	fmt.Println("user=nil   :", IsPtr(u))
	fmt.Println("user={}    :", IsPtr(uv))
	var ptr_u *user
	fmt.Println("*user=nil  :", IsPtr(ptr_u))
	fmt.Println("**user=nil :", IsPtr(&ptr_u))
	ptr_u = &user{}
	fmt.Println("*user={}   :", IsPtr(ptr_u))
	fmt.Println("**user={}  :", IsPtr(&ptr_u))
	ptr_uu := &ptr_u
	fmt.Println("***user={} :", IsPtr(&ptr_uu))
	ptr_uu = nil
	fmt.Println("***user=nil:", IsPtr(&ptr_uu))

	age := 100
	fmt.Println("nil        :", IsPtr(nil))
	fmt.Println("int        :", IsPtr(age))
	fmt.Println("*int       :", IsPtr(&age))
	fmt.Println("ok         :", IsPtr("ok"))
	fmt.Println("true       :", IsPtr(true))
}
