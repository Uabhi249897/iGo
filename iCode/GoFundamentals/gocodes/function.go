package gocodes

import (
	"fmt"
	"strings"
)

func Function() {

	module := "Golang"
	author := "Abhishek"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
