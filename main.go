package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello!")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", HelloServer)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
