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
