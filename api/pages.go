package api

import (
	"chatthread.net/app/main/domain"
	"chatthread.net/app/main/repository"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

var pagesPath = regexp.MustCompile("^/api/pages/?$")
var pagePath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/?$")

func handlePagesRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "DELETE":
		handleDeletePageRequest(w, r)
	case "POST":
		handlePostPageRequest(w, r)
	default:
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
	}
}

type createPage struct {
	Name string
	Url  string
}

func handlePostPageRequest(w http.ResponseWriter, r *http.Request) {
	m := pagesPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pageToCreate := createPage{}
	err = json.Unmarshal(body, &pageToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	parsedUrl, err := url.Parse(pageToCreate.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page := domain.CreatePage(pageToCreate.Name, *parsedUrl)
	repository.AddPage(page)
}

func handleDeletePageRequest(w http.ResponseWriter, r *http.Request) {
	m := pagePath.FindStringSubmatch(r.URL.Path)
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
