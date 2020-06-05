package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Running chatTHREAD app...")
	http.HandleFunc("/chatthread/", chatThreadPresentationHandler)
	http.HandleFunc("/api/chatthread/", chatThreadApiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
