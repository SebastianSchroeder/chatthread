package api

import (
	"chatthread.net/app/main/domain"
	"chatthread.net/app/main/repository"
	"github.com/google/uuid"
	"net/http"
	"regexp"
)

var repliesPath = regexp.MustCompile("^/api/pages/([a-zA-Z0-9\\-]+)/posts/([a-zA-Z0-9\\-]+)/replies/?$")

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
