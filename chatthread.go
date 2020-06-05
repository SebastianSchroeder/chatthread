package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/chatThread.html"))

func renderPosts(w http.ResponseWriter, posts *[]Post) {
	err := templates.ExecuteTemplate(w, "chatThread.html", posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Post struct {
	Text string
}

//var chatThreads = make(map[string][]Post)
var chatThreads = map[string][]Post{
	"/chatthread/foo": {
		Post{"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed"},
		Post{"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed"},
	},
}

func chatThreadHandler(w http.ResponseWriter, r *http.Request) {
	posts, exists := chatThreads[r.URL.Path]
	if !exists {
		http.NotFound(w, r)
		return
	}
	renderPosts(w, &posts)
}
