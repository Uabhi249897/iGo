package gocodes

import (
	"log"
	"net/http"
)

func Querywebservice() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}

		w.Write([]byte("Hello " + name))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

//go build -o webservice1 main.go
//chmod +x main
//./Querywebservice
//http://localhost:3000/?name=Mike
