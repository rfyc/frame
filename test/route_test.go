package test

import (
	"fmt"
	"github.com/rfyc/frame/utils/structs"
	"testing"
)

type A struct {
	Name string
	Age  int
}

func TestSet(t *testing.T) {

	ua := &A{}
	ma := map[string]interface{}{
		"name": "hong",
		"age":  true,
	}
	structs.Set(ua, ma)
	fmt.Println(ua)
}
