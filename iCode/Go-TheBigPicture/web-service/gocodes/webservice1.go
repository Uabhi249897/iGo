package gocodes

import (
	"log"
	"net/http"
)

func Webservice1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

//go build -o webservice1 main.go
//chmod +x main
//./webservice1
