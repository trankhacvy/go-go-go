package posts

import(
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Content string
	UserID uint
}

func (Post) TableName() string {
	return "tbl_posts"
}