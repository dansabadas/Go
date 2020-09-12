package main

import (
	"fmt"
	"reflect"
)

var (
	name                    string
	course                  string
	module                  float64
	name1, course1, module1 = "dan", "go", 17.2
)

func main() {
	fmt.Println("name is set to", name, "is of type", reflect.TypeOf(name))
	fmt.Println("module is set to", module, "is of type", reflect.TypeOf(module))
	fmt.Println("name1 is set to", name1, "is of type", reflect.TypeOf(name1))
	fmt.Println("module1 is set to", module1, "is of type", reflect.TypeOf(module1))

	a := 10.00000
	b := 3

	fmt.Println("a is set to", a, "is of type", reflect.TypeOf(a))
	fmt.Println("b is set to", b, "is of type", reflect.TypeOf(b))

	//c1 := a + b invalid mismatch from compiler
	//fmt.Println("c1 is set to", c1, "is of type", reflect.TypeOf(c1))

	c2 := int(a) + b
	fmt.Println("c2 is set to", c2, "is of type", reflect.TypeOf(c2))
}
