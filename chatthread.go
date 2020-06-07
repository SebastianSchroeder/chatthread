package main

import (
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("templates/page.html"))

type PagePresentation struct {
	Name  string
	Posts *[]Post
}

type Page struct {
	Name string
}

type Post struct {
	Text  string
	Posts []Post
}

var pages = map[Page][]Post{
	Page{"foo"}: {
		{Text: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut "},
		{Text: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt"},
	},
	Page{"bar"}: {},
}

func renderPage(w http.ResponseWriter, page *PagePresentation) {
	err := templates.ExecuteTemplate(w, "page.html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var presentationPath = regexp.MustCompile("^/page/([a-zA-Z0-9]+)/?$")

func pagePresentationHandler(w http.ResponseWriter, r *http.Request) {
	m := presentationPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	page := Page{m[1]}
	posts, exists := pages[page]
	if !exists {
		http.NotFound(w, r)
		return
	}
	pagePresentation := PagePresentation{page.Name, &posts}
	renderPage(w, &pagePresentation)
}

var apiPath = regexp.MustCompile("^/api/page/([a-zA-Z0-9]+)/?$")

func pageApiHandler(w http.ResponseWriter, r *http.Request) {
	m := apiPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	page := Page{m[1]}
	posts, exists := pages[page]
	if !exists {
		http.NotFound(w, r)
		return
	}
	post := r.FormValue("post")
	posts = append(posts, Post{Text: post})
	pages[page] = posts
	http.Redirect(w, r, "/page/"+page.Name, http.StatusFound)
}
