package domain

import (
	"github.com/google/uuid"
	"net/url"
	"time"
)

type Page struct {
	Id      uuid.UUID
	Name    string
	Url     url.URL
	Created time.Time
}

func CreatePage(name string, url url.URL) Page {
	return Page{
		Id:      uuid.New(),
		Name:    name,
		Url:     url,
		Created: time.Now(),
	}
}
