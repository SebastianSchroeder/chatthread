package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Running chatTHREAD app...")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/page/", pagePresentationHandler)
	http.HandleFunc("/api/page/", pageApiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
