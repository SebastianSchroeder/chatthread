package api

import (
	"chatthread.net/app/main/domain"
	"chatthread.net/app/main/repository"
	"github.com/google/uuid"
	"net/http"
	"regexp"
)

var postsPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/posts/?$")

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
