package main

import (
	"fmt"
	"reflect"
)

// 使用reflect.MakeFunc方法Mock任意一个函数

func main() {
	originFn := func(a, b int) (int, int) {
		return a + b, a+b
	}

	mockedFunc := MockFunc(originFn).(func(int, int) (int, int))
	result1, result2 := mockedFunc(1, 2)
	fmt.Println(result1)
	fmt.Println(result2)

	// originFunc赋值给newFunc，再改变newFunc的value，无需显式进行interface{}到具体函数类型的具体转换
	newFunc := originFn
	reflect.ValueOf(&newFunc).Elem().Set(reflect.ValueOf(MockFunc(originFn)))
	result1, result2 = newFunc(1, 2)
	fmt.Println(result1)
	fmt.Println(result2)

	// 改变newFunc的值，不会改变originFn的值
	result1, result2 = originFn(1, 2)
	fmt.Println(result1)
	fmt.Println(result2)
}

func MockFunc(originFunc interface{}) interface{} {

	numOut := reflect.TypeOf(originFunc).NumOut()
	mockedFuncValue := func(args []reflect.Value) (results []reflect.Value) {
		for i := 0; i < numOut; i++ {
			outParamType := reflect.TypeOf(originFunc).Out(i)
			results = append(results, reflect.Zero(outParamType))
		}
		return
	}
	mockedFunc := reflect.MakeFunc(reflect.TypeOf(originFunc), mockedFuncValue)
	return mockedFunc.Interface()
	return nil
}