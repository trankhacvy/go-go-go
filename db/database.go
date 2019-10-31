package db

import (
	"fmt"
	"os"

	"github.com/Levi-ackerman/go-go-go/posts"
	"github.com/Levi-ackerman/go-go-go/users"
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	// dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPass)
	fmt.Println(dbUri)

	db, err = gorm.Open(dbDriver, dbUri)
	db.LogMode(true)
	if err != nil {
		panic(err)
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
