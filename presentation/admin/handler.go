package admin

import (
	"chatthread.net/app/main/domain"
	"chatthread.net/app/main/repository"
	"net/http"
	"regexp"
)

type AdminPagesPresentation struct {
	Pages   *[]domain.Page
	JsHash  string
	CssHash string
}

var adminPagesPresentationPath = regexp.MustCompile("^/admin/pages/?$")

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetPagesRequest(w, r)
	default:
		http.Error(w, "415 method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetPagesRequest(w http.ResponseWriter, r *http.Request) {
	if !adminPagesPresentationPath.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}
	pages := repository.ListPages()
	adminPagesPresentation := AdminPagesPresentation{&pages, ChatThreadJsHash, ChatThreadCssHash}
	RenderPage(w, "pages.html", &adminPagesPresentation)
}
