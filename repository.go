package main

import (
	"github.com/google/uuid"
	"net/url"
	"time"
)

var pages = map[Page]*[]Post{
	Page{
		PageId:  uuid.New(),
		Name:    "welcome",
		Url:     url.URL{Scheme: "https", Host: "chatthread.net", Path: "welcome"},
		Created: time.Now(),
	}: {
		{
			PostId:  uuid.New(),
			Text:    "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut ",
			Created: time.Now(),
		},
		{
			PostId:  uuid.New(),
			Text:    "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt",
			Created: time.Now(),
		},
	},
	Page{
		PageId:  uuid.New(),
		Name:    "hello",
		Url:     url.URL{Scheme: "https", Host: "chatthread.net", Path: "hello"},
		Created: time.Now(),
	}: {},
}

func retrievePageByName(name string, pages map[Page]*[]Post) (Page, *[]Post, bool) {
	for page, posts := range pages {
		if name == page.Name {
			return page, posts, true
		}
	}
	return Page{}, nil, false
}

func addPost(name string, post Post, pages map[Page]*[]Post) bool {
	_, posts, exits := retrievePageByName(name, pages)
	if !exits {
		return false
	}
	*posts = append(*posts, post)
	return true
}
