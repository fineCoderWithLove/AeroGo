package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
	Age  int
	No   int
}
type Student struct {
	SNo int
	People
}

func main() {
	i := 333
	f(i)

	s := Student{1, People{"kika", 33, 1}}
	f(s)

	p := People{"lisi", 22, 123}
	f(p)

	f(&p)
}

func f(i interface{}) {
	t := reflect.TypeOf(i) // 真实类型，包括自定义类型
	v := reflect.ValueOf(i)
	k := v.Kind() // golang 内置类别，例如 map，struct，string，bool，int 等
	if k == reflect.Ptr {
		e := v.Elem() // 如果传入的是指针，通过 Elem() 拿到指向的值
		fmt.Println("Elem is:", e)
	}
	fmt.Printf("type is: %s\nvalue is: %v\nkind is: %s\n\n", t, v, k)
	fmt.Println()
}
