package gocodes

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func Flaglogparse() {

	path := flag.String("path", "error.log", "The path to the log that should be analyzed")
	level := flag.String("level", "http", "Log level to be search for. Options are HTTP, http, WARNING, INFO")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}

//go run main.go -help
//go run main.go
//go run main.go -level HTTP
