package main

import (
	"chatthread.net/app/main/api"
	"chatthread.net/app/main/presentation"
	"chatthread.net/app/main/presentation/admin"
	"chatthread.net/app/main/util/compression"
	"log"
	"net/http"
)

func main() {
	log.Print("Running chatTHREAD app...")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", compression.GzipCompression(fs)))
	http.HandleFunc("/pages/", presentation.PagesHandler)
	http.HandleFunc("/api/pages/", api.PagesHandler)
	http.HandleFunc("/admin/pages/", admin.PagesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
