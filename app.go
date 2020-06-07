package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Running chatTHREAD app...")
	http.HandleFunc("/page/", pagePresentationHandler)
	http.HandleFunc("/api/page/", pageApiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
