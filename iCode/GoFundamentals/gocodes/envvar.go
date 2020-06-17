package gocodes

import (
	"fmt"
	"os"
)

func Envvar() {
	// fmt.Println(os.Environ())

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	name := os.Getenv("USERNAME")
	fmt.Println(name)
}
