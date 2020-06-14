package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Running chatTHREAD app...")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", gzipCompression(fs)))
	http.HandleFunc("/pages/", pagesPresentationHandler)
	http.HandleFunc("/api/pages/", pagesApiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
