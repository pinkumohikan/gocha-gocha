package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		res := map[string]string{
			"message": "pong",
		}
		j, _ := json.Marshal(res)
		fmt.Fprintf(w, string(j))
	})

	type GreetInput struct {
		Name string `json:"name"`
	}

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		var i GreetInput
		ij, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}

		if err := json.Unmarshal(ij, &i); err != nil {
			log.Fatalln(err)
		}

		w.Header().Set("content-type", "application/json")
		j, _ := json.Marshal(map[string]string{
			"greeting": "Hello " + i.Name,
		})
		fmt.Fprintf(w, string(j))
	})

	log.Fatalln(http.ListenAndServe("0.0.0.0:8080", nil))
}
