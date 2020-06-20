package presentation

import (
	"chatthread.net/app/main/domain"
	"chatthread.net/app/main/repository"
	"net/http"
	"regexp"
)

type PagePresentation struct {
	Page    domain.Page
	Posts   *[]domain.Post
	JsHash  string
	CssHash string
}

var pagesPath = regexp.MustCompile("^/pages/([a-zA-Z0-9]+)/?$")

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetPageRequest(w, r)
	default:
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetPageRequest(w http.ResponseWriter, r *http.Request) {
	m := pagesPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	pageName := m[1]
	page, posts, exists := repository.RetrievePageByName(pageName)
	if !exists {
		http.NotFound(w, r)
		return
	}
	pagePresentation := PagePresentation{page, posts, chatThreadJsHash, chatThreadCssHash}
	renderPage(w, "page.html", &pagePresentation)
}
