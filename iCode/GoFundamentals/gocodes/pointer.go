package gocodes

import (
	"fmt"
)

func Pointer() {

	// < -- Pointer -- >

	// module := 4.5

	// fmt.Println("Memory address of *module variable is ", &module)

	// ptr := &module

	// fmt.Println("Memory address of *module variable is ", ptr, "and the value is ", *ptr)

	// < -- Pass by Value -- >

	// name := "Abhishek"
	// course := "Golang"

	// fmt.Println("\n Hi", name, "you're currently watching", course)

	// changeCoursebyval(course)

	// fmt.Println("\n Hi", name, "Unchanged course value is ", course)

	// < -- Pass by Reference -- >

	name := "Abhishek"
	course := "Golang"

	fmt.Println("\n Hi", name, "you're currently watching", course)

	changeCoursebyref(&course)

	fmt.Println("\n Hi", name, "Changed course reference value is ", course)

}

func changeCoursebyval(course string) string {
	//return value could be name(like string) or unnamed
	course = "Pointer in Golang"

	fmt.Println("\n Changed the course is ", course)

	return course
}

func changeCoursebyref(course *string) string {
	//return value could be name(like string) or unnamed
	*course = "Pointer in Golang"

	fmt.Println("\n Changed the course is ", *course)

	return *course
}
