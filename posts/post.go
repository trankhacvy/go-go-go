package posts

import (
	"time"
)

type Post struct {
	ID      uint `json:"id"`
	Content string `json:"content"`
	UserID  uint `json:"userId"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func (Post) TableName() string {
	return "tbl_posts"
}
