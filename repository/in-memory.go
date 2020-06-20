package repository

import (
	"chatthread.net/app/main/domain"
	"github.com/google/uuid"
	"net/url"
	"time"
)

var welcomePage = domain.Page{
	Id:      uuid.New(),
	Name:    "welcome",
	Url:     url.URL{Scheme: "https", Host: "chatthread.net", Path: "welcome"},
	Created: time.Now(),
}

var helloPage = domain.Page{
	Id:      uuid.New(),
	Name:    "hello",
	Url:     url.URL{Scheme: "https", Host: "chatthread.net", Path: "hello"},
	Created: time.Now(),
}

var pagesToPosts = map[domain.Page]*[]domain.Post{
	welcomePage: {
		domain.CreatePost(
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut ",
			welcomePage.Id,
		),
		domain.CreatePost(
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt",
			welcomePage.Id,
		),
	},
	helloPage: {},
}

func RetrievePageByName(pageName string) (domain.Page, *[]domain.Post, bool) {
	for page, posts := range pagesToPosts {
		if pageName == page.Name {
			return page, posts, true
		}
	}
	return domain.Page{}, nil, false
}

func RetrievePageById(pageId uuid.UUID) (domain.Page, *[]domain.Post, bool) {
	for page, posts := range pagesToPosts {
		if pageId == page.Id {
			return page, posts, true
		}
	}
	return domain.Page{}, nil, false
}

func DeletePageById(pageId uuid.UUID) bool {
	for page := range pagesToPosts {
		if pageId == page.Id {
			delete(pagesToPosts, page)
			return true
		}
	}
	return false
}

func AddPage(page domain.Page) {
	posts := make([]domain.Post, 5)
	pagesToPosts[page] = &posts
}

func ListPages() []domain.Page {
	i := 0
	pages := make([]domain.Page, len(pagesToPosts))
	for k := range pagesToPosts {
		pages[i] = k
		i++
	}
	return pages
}

func AddPost(post domain.Post, posts *[]domain.Post) bool {
	*posts = append(*posts, post)
	return true
}

func AddReply(postId uuid.UUID, reply domain.Post, posts *[]domain.Post) bool {
	for _, post := range *posts {
		if post.Id == postId {
			*post.Replies = append(*post.Replies, reply)
			return true
		} else if AddReply(postId, reply, post.Replies) {
			return true
		}
	}
	return false
}
