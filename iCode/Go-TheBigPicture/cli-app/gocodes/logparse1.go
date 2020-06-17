package gocodes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Logparse1() {
	f, err := os.Open("error.log")
	if err != nil {
		log.Fatal(err)
	}

	//after defer calls exit, its going to goahead and close those at the end (Resource Clean up)
	defer f.Close()

	//input string : Network Conn, String
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, "HTTP") {
			fmt.Println(s)
		}
		if strings.Contains(s, "http") {
			fmt.Println(s)
		}
	}
}
