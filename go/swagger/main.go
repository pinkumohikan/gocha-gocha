package main

import (
	"encoding/json"
	"fmt"
	"github.com/pinkumohikan/gocha-gocha/go/swagger/api"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("POST /greet: called")

		ij, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}

		var i api.GreetRequest
		if err := json.Unmarshal(ij, &i); err != nil {
			log.Fatalln(err)
		}

		w.Header().Set("content-type", "application/json")
		j, _ := json.Marshal(api.GreetResponse{
			Greeting: "Hello " + i.Name,
		})
		fmt.Fprintf(w, string(j))
	})

	log.Fatalln(http.ListenAndServe("0.0.0.0:8080", nil))
}
