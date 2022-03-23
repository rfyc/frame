package structs

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

func TestFields(t *testing.T) {
	var (
		uv     = user{Name: "123"}
		ptr_uv = &uv
	)
	fmt.Println("user={}   :", Fields(uv))
	fmt.Println("*user={}  :", Fields(ptr_uv))
	fmt.Println("**user={} :", Fields(&ptr_uv))
	fmt.Println("*user=nil : panic not struct")
}
func TestPtrOf(t *testing.T) {
	var (
		uv     = user{Name: "123"}
		ptr_uv = &uv
		ptr_u  *user
	)
	fmt.Println("*user={}   :", PtrOf(&uv).MethodByName("Echo").Call([]reflect.Value{}))
	fmt.Println("*user set field :")
	PtrOf(&uv).Elem().FieldByName("Name").Set(reflect.ValueOf("666"))
	fmt.Println("*user={}   :", PtrOf(&uv).MethodByName("Echo").Call([]reflect.Value{}))
	fmt.Println("**user={}  :", PtrOf(&ptr_uv).MethodByName("Echo").Call([]reflect.Value{}))
	fmt.Println("*user=nil  :", PtrOf(ptr_u))
	fmt.Println("user={}    : error: argv not ptr")
	fmt.Println("argv=nil   : error: argv not ptr")
	//PtrOf(nil)
	//PtrOf(uv)
}

func TestValueOf(t *testing.T) {

	var (
		uv     = user{Name: "123"}
		ptr_uv = &uv
		ptr_u  *user
	)

	fmt.Println("123        :", ValueOf(123))
	fmt.Println("*user=nil  :", ValueOf(ptr_u))
	fmt.Println("user={}    :", ValueOf(uv).FieldByName("Name").String())
	fmt.Println("*user={}   :", ValueOf(&uv).FieldByName("Name").String())
	fmt.Println("**user={}  :", ValueOf(ptr_uv).FieldByName("Name").String())
	fmt.Println("**user={}  :", ValueOf(&ptr_uv).FieldByName("Name").String())
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
