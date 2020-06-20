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

var templates = template.Must(template.ParseFiles(
	"templates/page.html",
	"templates/admin/pages-admin.html",
	"templates/admin/page-admin.html",
))
var chatthreadJsHash = computeFileHash("static/js/chatthread.js")
var chatThreadAdminJsHash = computeFileHash("static/js/chatthread-admin.js")
var chatthreadCssHash = computeFileHash("static/css/chatthread.css")
var chatThreadAdminCssHash = computeFileHash("static/css/chatthread-admin.css")

type PagePresentation struct {
	Page    Page
	Posts   *[]Post
	JsHash  string
	CssHash string
}

type Page struct {
	Id      uuid.UUID
	Name    string
	Url     url.URL
	Created time.Time
}

type Post struct {
	Id      uuid.UUID
	PageId  uuid.UUID
	Text    string
	Created time.Time
	Replies *[]Post
}

func createPost(text string, pageId uuid.UUID) Post {
	return Post{
		Id:      uuid.New(),
		PageId:  pageId,
		Text:    text,
		Created: time.Now(),
		Replies: &[]Post{}}
}

func renderPage(w http.ResponseWriter, template string, data interface{}) {
	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var pagesPresentationPath = regexp.MustCompile("^/pages/([a-zA-Z0-9]+)/?$")

func pagesPresentationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	m := pagesPresentationPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	pageName := m[1]
	page, posts, exists := retrievePageByName(pageName, pagesToPosts)
	if !exists {
		http.NotFound(w, r)
		return
	}
	pagePresentation := PagePresentation{page, posts, chatthreadJsHash, chatthreadCssHash}
	renderPage(w, "page.html", &pagePresentation)
}

var pagesApiPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/?$")
var postsApiPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/posts/?$")
var repliesApiPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/posts/([a-zA-Z0-9\\-]+)/replies/?$")

func pagesApiHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case pagesApiPath.MatchString(r.URL.Path):
		log.Print("matching page handler")
		handlePagesRequest(w, r)
	case postsApiPath.MatchString(r.URL.Path):
		log.Print("matching post handler")
		handlePostsRequest(w, r)
	case repliesApiPath.MatchString(r.URL.Path):
		log.Print("matching reply handler")
		handleRepliesRequest(w, r)
	default:
		log.Print("path ", r.URL.Path, " does not match")
		http.NotFound(w, r)
	}
}

func handlePagesRequest(w http.ResponseWriter, r *http.Request) {

}

func handlePostsRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	m := postsApiPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	pageId, failure := uuid.Parse(m[1])
	if failure != nil {
		http.NotFound(w, r)
		return
	}
	page, posts, exists := retrievePageById(pageId, pagesToPosts)
	if !exists {
		http.NotFound(w, r)
		return
	}
	text := r.FormValue("text")
	addPost(createPost(text, page.Id), posts)
	http.Redirect(w, r, "/pages/"+page.Name, http.StatusFound)
}

func handleRepliesRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	m := repliesApiPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	pageId, failure := uuid.Parse(m[1])
	if failure != nil {
		http.NotFound(w, r)
		return
	}
	postId, failure := uuid.Parse(m[2])
	if failure != nil {
		http.NotFound(w, r)
		return
	}
	page, posts, exists := retrievePageById(pageId, pagesToPosts)
	if !exists {
		http.NotFound(w, r)
		return
	}
	text := r.FormValue("text")
	addReply(postId, createPost(text, page.Id), posts)
	http.Redirect(w, r, "/pages/"+page.Name, http.StatusFound)
}

type AdminPagesPresentation struct {
	Pages   *[]Page
	JsHash  string
	CssHash string
}

var adminPagesPresentationPath = regexp.MustCompile("^/admin/pages/?$")

func adminPagesPresentationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !adminPagesPresentationPath.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}
	i := 0
	pages := make([]Page, len(pagesToPosts))
	for k := range pagesToPosts {
		pages[i] = k
		i++
	}
	adminPagesPresentation := AdminPagesPresentation{&pages, chatThreadAdminJsHash, chatThreadAdminCssHash}
	renderPage(w, "pages-admin.html", &adminPagesPresentation)
}
