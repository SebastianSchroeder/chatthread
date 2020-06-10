package main

import (
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var templates = template.Must(template.ParseFiles("templates/page.html"))

type PagePresentation struct {
	Page  Page
	Posts *[]Post
}

type Page struct {
	PageId  uuid.UUID
	Name    string
	Url     url.URL
	Created time.Time
}

type Post struct {
	PostId  uuid.UUID
	Text    string
	Created time.Time
	Replies *[]Post
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
	pageName := m[1]
	page, posts, exists := retrievePageByName(pageName, pages)
	if !exists {
		http.NotFound(w, r)
		return
	}
	pagePresentation := PagePresentation{page, posts}
	renderPage(w, &pagePresentation)
}

var pageApiPath = regexp.MustCompile("^/api/page/([a-zA-Z0-9\\-]+)/?$")
var postApiPath = regexp.MustCompile("^/api/page/([a-zA-Z0-9\\-]+)/posts/?$")
var replyApiPath = regexp.MustCompile("^/api/page/([a-zA-Z0-9\\-]+)/posts/([a-zA-Z0-9\\-]+)/replies/?$")

func pageApiHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case pageApiPath.MatchString(r.URL.Path):
		log.Print("matching page handler")
		handlePageRequest(w, r)
	case postApiPath.MatchString(r.URL.Path):
		log.Print("matching post handler")
		handlePostRequest(w, r)
	case replyApiPath.MatchString(r.URL.Path):
		log.Print("matching reply handler")
		handleReplyRequest(w, r)
	default:
		log.Print("path ", r.URL.Path, " does not match")
		http.NotFound(w, r)
	}
}

func handlePageRequest(w http.ResponseWriter, r *http.Request) {

}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	m := postApiPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	pageName := m[1]
	page, _, exists := retrievePageByName(pageName, pages)
	if !exists {
		http.NotFound(w, r)
		return
	}
	post := r.FormValue("post")
	addPost(pageName, Post{PostId: uuid.New(), Text: post, Created: time.Now()}, pages)
	http.Redirect(w, r, "/page/"+page.Name, http.StatusFound)
}

func handleReplyRequest(w http.ResponseWriter, r *http.Request) {

}
