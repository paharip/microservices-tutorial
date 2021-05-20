package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/paharip/microservices-tutorial/data"
)

func main() {

	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	http.HandleFunc("/helloworld1", helloWorldHandler1)

	cathandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	reposne := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(reposne)
	if err != nil {
		panic("Ooops")
	}
	fmt.Fprint(w, string(data))
}

func helloWorldHandler1(w http.ResponseWriter, r *http.Request) {

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
