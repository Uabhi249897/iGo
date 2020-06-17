package gocodes

import (
	"fmt"
	"reflect"
)

var (
// name   string //initialized with empty string
// course string
// module float64 //initialized with zero value

// name, course, module = "Abhishek", "Golang", 4.2
)

func Typeofvar() {
	// fmt.Println("Name is set to", name, "and is type of", reflect.TypeOf(name))
	// fmt.Println("Course is set to", course, "and is type of", reflect.TypeOf(course))
	// fmt.Println("Module is set to", module, "and is type of", reflect.TypeOf(module))

	a := 10.00
	b := 3

	fmt.Println("\nA is type of ", reflect.TypeOf(a), "and B is type of ", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("\nC is type of ", reflect.TypeOf(c))
}
