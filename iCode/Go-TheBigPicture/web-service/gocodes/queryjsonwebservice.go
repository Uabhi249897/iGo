package gocodes

import (
	"encoding/json"
	"log"
	"net/http"
)

func Queryjsonwebservice() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}

		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)
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
