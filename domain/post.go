package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id      uuid.UUID
	PageId  uuid.UUID
	Text    string
	Created time.Time
	Replies *[]Post
}

func CreatePost(text string, pageId uuid.UUID) Post {
	return Post{
		Id:      uuid.New(),
		PageId:  pageId,
		Text:    text,
		Created: time.Now(),
		Replies: &[]Post{}}
}
