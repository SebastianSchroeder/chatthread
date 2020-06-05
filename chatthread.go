package main

import (
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("templates/chatThread.html"))

type ChatThread struct {
	Name  string
	Posts []Post
}

type Post struct {
	Text string
}

func renderPosts(w http.ResponseWriter, chatThread *ChatThread) {
	err := templates.ExecuteTemplate(w, "chatThread.html", chatThread)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var chatThreads = map[string]ChatThread{
	"foo": {
		Name: "foo",
		Posts: []Post{
			{"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed"},
			{"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed"},
		},
	},
}

var presentationPath = regexp.MustCompile("^/chatthread/([a-zA-Z0-9]+)$")

func chatThreadPresentationHandler(w http.ResponseWriter, r *http.Request) {
	m := presentationPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	chatThread, exists := chatThreads[m[1]]
	if !exists {
		http.NotFound(w, r)
		return
	}
	renderPosts(w, &chatThread)
}

var apiPath = regexp.MustCompile("^/api/chatthread/([a-zA-Z0-9]+)$")

func chatThreadApiHandler(w http.ResponseWriter, r *http.Request) {
	m := apiPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	chatThread, exists := chatThreads[m[1]]
	if !exists {
		http.NotFound(w, r)
		return
	}
	post := r.FormValue("post")
	chatThread.Posts = append(chatThread.Posts, Post{post})
	chatThreads[chatThread.Name] = chatThread
	http.Redirect(w, r, "/chatthread/"+chatThread.Name, http.StatusFound)
}
