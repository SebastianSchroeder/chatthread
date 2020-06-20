package api

import (
	"chatthread.net/app/main/domain"
	"chatthread.net/app/main/repository"
	"github.com/google/uuid"
	"log"
	"net/http"
	"regexp"
)

var pagesPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/?$")
var postsPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/posts/?$")
var repliesPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/posts/([a-zA-Z0-9\\-]+)/replies/?$")

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case pagesPath.MatchString(r.URL.Path):
		log.Print("matching page handler")
		handlePagesRequest(w, r)
	case postsPath.MatchString(r.URL.Path):
		log.Print("matching post handler")
		handlePostsRequest(w, r)
	case repliesPath.MatchString(r.URL.Path):
		log.Print("matching reply handler")
		handleRepliesRequest(w, r)
	default:
		log.Print("path ", r.URL.Path, " does not match")
		http.NotFound(w, r)
	}
}

func handlePostsRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	m := postsPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	pageId, failure := uuid.Parse(m[1])
	if failure != nil {
		http.NotFound(w, r)
		return
	}
	page, posts, exists := repository.RetrievePageById(pageId)
	if !exists {
		http.NotFound(w, r)
		return
	}
	text := r.FormValue("text")
	repository.AddPost(domain.CreatePost(text, page.Id), posts)
	http.Redirect(w, r, "/pages/"+page.Name, http.StatusFound)
}

func handleRepliesRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	m := repliesPath.FindStringSubmatch(r.URL.Path)
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
	page, posts, exists := repository.RetrievePageById(pageId)
	if !exists {
		http.NotFound(w, r)
		return
	}
	text := r.FormValue("text")
	repository.AddReply(postId, domain.CreatePost(text, page.Id), posts)
	http.Redirect(w, r, "/pages/"+page.Name, http.StatusFound)
}

func handlePagesRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	m := pagesPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	pageId, failure := uuid.Parse(m[1])
	if failure != nil {
		http.NotFound(w, r)
		return
	}
	deleted := repository.DeletePageById(pageId)
	if !deleted {
		http.NotFound(w, r)
		return
	}
}
