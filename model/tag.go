package model

import (
	"time"
)

type Tag struct {
	Id    int64   `json:"id"`
	Label string  `json:"label"`
	Posts []*Post `json:"posts,omitempty"`

	CreatedDate *time.Time `json:"created_date,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
}
