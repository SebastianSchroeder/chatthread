package main

import (
	"github.com/google/uuid"
	"net/url"
	"time"
)

var welcomePage = Page{
	Id:      uuid.New(),
	Name:    "welcome",
	Url:     url.URL{Scheme: "https", Host: "chatthread.net", Path: "welcome"},
	Created: time.Now(),
}

var helloPage = Page{
	Id:      uuid.New(),
	Name:    "hello",
	Url:     url.URL{Scheme: "https", Host: "chatthread.net", Path: "hello"},
	Created: time.Now(),
}

var pages = map[Page]*[]Post{
	welcomePage: {
		createPost(
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut ",
			welcomePage.Id,
		),
		createPost(
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt",
			welcomePage.Id,
		),
	},
	helloPage: {},
}

func retrievePageByName(pageName string, pages map[Page]*[]Post) (Page, *[]Post, bool) {
	for page, posts := range pages {
		if pageName == page.Name {
			return page, posts, true
		}
	}
	return Page{}, nil, false
}

func retrievePageById(pageId uuid.UUID, pages map[Page]*[]Post) (Page, *[]Post, bool) {
	for page, posts := range pages {
		if pageId == page.Id {
			return page, posts, true
		}
	}
	return Page{}, nil, false
}

func addPost(post Post, posts *[]Post) bool {
	*posts = append(*posts, post)
	return true
}

func addReply(postId uuid.UUID, reply Post, posts *[]Post) bool {
	for _, post := range *posts {
		if post.Id == postId {
			*post.Replies = append(*post.Replies, reply)
			return true
		} else if addReply(postId, reply, post.Replies) {
			return true
		}
	}
	return false
}
