package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-go-go/users"
	"go-go-go/posts"
	"os"
)

var db *gorm.DB
var err error

func Init() {

	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbName := os.Getenv("db_name")
	dbDriver := os.Getenv("db_type")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(dbUri)

	db, err = gorm.Open(dbDriver, dbUri)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connect to db success")
	db.AutoMigrate(&users.User{}, &posts.Post{})

	//defer func() {
	//	err := db.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()
}

func GetDB() *gorm.DB {
	return db
}
