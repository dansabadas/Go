package main

import (
	"fmt"
	"os"
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

	name2 := "danson"
	fmt.Println("name2 is set to", name2, "is of type", reflect.TypeOf(name2))

	//name3 := "danson" will fail by compiler if we don't use it!

	module := 3

	fmt.Println("memory address of module variable is", &module)
	ptr := &module
	fmt.Println("memory address of module variable is", ptr, "and the value stored is", *ptr)

	course := "plural go"

	fmt.Println("we are currently watching", course)
	changeCourse(course)
	fmt.Println("we are currently watching", course)

	name := "dan"

	fmt.Println("we are ", name)
	changeName(&name)
	fmt.Println("we are ", name)

	const speedOfLightKm = 300000
	fmt.Println("speed of light = ", speedOfLightKm)

	//speedOfLightKm = 34 we cannot re-assign to a constant

	fmt.Println(os.Environ())

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	currentName := os.Getenv("USERNAME")

	fmt.Println("currentName is set to", currentName, "is of type", reflect.TypeOf(currentName))

	fmt.Println("address of currentName=", &currentName)

	namePtr := &currentName
	fmt.Println("namePtr is set to", namePtr, "is of type", reflect.TypeOf(namePtr))
	fmt.Println("&namePtr is set to", &namePtr, "is of type", reflect.TypeOf(&namePtr))
	fmt.Println("&namePtr=", &namePtr)
	testNameAddress(namePtr)
}

func changeCourse(course string) string {
	course = course + "2" // it does not work := assignment outside main!because is not declaration
	fmt.Println("trying to change course to", course)
	return course
}

func changeName(name *string) string {
	*name = "ioan"
	fmt.Println("we are trying to be ", *name)
	return *name
}

func testNameAddress(name *string) {
	fmt.Println("address of currentName=", name)
	fmt.Println("&namePtr=", &name)
}

// C:\Dan\Code\Personal\Go> go run c:\Dan\Code\Personal\Go\src\Fundamentals\variables-and-constants.go
