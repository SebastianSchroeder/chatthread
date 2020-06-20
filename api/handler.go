package api

import (
	"log"
	"net/http"
	"strings"
)

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.Contains(r.URL.Path, "/replies"):
		log.Print("matching reply handler")
		handleRepliesRequest(w, r)
	case strings.Contains(r.URL.Path, "/posts"):
		log.Print("matching post handler")
		handlePostsRequest(w, r)
	case strings.Contains(r.URL.Path, "/pages"):
		log.Print("matching page handler")
		handlePagesRequest(w, r)
	default:
		log.Print("path ", r.URL.Path, " does not match")
		http.NotFound(w, r)
	}
}
