package users

import (
	"go-go-go/posts"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserID uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Username  string `json:"username",gorm:"unique;not null"`
	Password  string `json:"-"`
	Token     string `json:"token",gorm:"-",sql:"-"`
	Posts     []posts.Post `json:"posts"`
	Email     string `json:"email",gorm:"unique;not null"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Set User's table name to be `profiles`
func (User) TableName() string {
	return "tbl_users"
}
