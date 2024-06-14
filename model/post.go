package model

import (
	"time"
)

type Post struct {
	Id            int64      `json:"id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	Status        string     `json:"status,omitempty"`
	PublishedDate *time.Time `json:"published_date,omitempty"`
	Tags          []*Tag     `json:"tags,omitempty"`

	CreatedDate *time.Time `json:"created_date,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
}

type Param struct {
	Key   string
	Value string
}
