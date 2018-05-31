package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/handle", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v\n", err)
			return
		}

		fmt.Printf("method = %s\n", r.Method)

		switch r.Method {
		case "GET":
			fmt.Printf("OK!\n")
		case "POST":
			m := make(map[string]string)

			for key, value := range r.Form {
				m[key] = strings.Join(value, "")
			}

			json, err := json.Marshal(m)
			fmt.Println(string(json), err)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.\n")
		}
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
