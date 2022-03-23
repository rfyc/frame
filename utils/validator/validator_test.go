package validator

import (
	"errors"
	"fmt"
	"testing"
)

type user struct {
	Name string
	Age  int
}

func (this *user) Uniq() error {

	return errors.New("name not uniq")
}

func TestMethod(t *testing.T) {
	u := &user{
		Name: "hong",
		Age:  0,
	}
	var method Method = u.Uniq
	fmt.Println(method.Validate())
}

func TestRequired(t *testing.T) {
	u := &user{
		Name: "hong",
		Age:  0,
	}
	required := Required{
		Struct:  u,
		Names:   "name,age",
		IsEmpty: true,
		Error:   errors.New("params required"),
	}
	fmt.Println(required.Validate())
}
