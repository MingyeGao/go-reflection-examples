// 将struct中的每一个字段置为zero值

package main

import (
	"fmt"
	"reflect"
)

type A struct {
	One int
	Two string
	Three *string
	Four chan int
	Five func(int, string) map[string]map[string]bool
}

func main() {
	s := "3"
	a := &A{
		One: 1,
		Two: "2",
		Three: &s,
	}
	for i := 0; i < reflect.ValueOf(a).Elem().NumField(); i++ {
		field := reflect.ValueOf(a).Elem().Field(i)
		fieldType := reflect.TypeOf(field.Interface())
		field.Set(reflect.Zero(fieldType))
	}
	fmt.Println(a)
}
