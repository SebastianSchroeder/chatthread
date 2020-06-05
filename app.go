package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/chatthread/", chatThreadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
