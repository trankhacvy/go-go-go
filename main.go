package main

import (
	"fmt"
	"log"

	database "github.com/Levi-ackerman/go-go-go/db"

	"net/http"
	"os"

	"github.com/Levi-ackerman/go-go-go/files"
	"github.com/Levi-ackerman/go-go-go/middlewares"
	"github.com/Levi-ackerman/go-go-go/posts"
	"github.com/Levi-ackerman/go-go-go/users"

	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// func init() {
// 	e := godotenv.Load()
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	database.Init()

	router := mux.NewRouter()
	router.Use(middlewares.JwtAuthentication)

	translator := en.New()
	uit := ut.New(translator, translator)
	trans, found := uit.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	validators := validator.New()
	if err := en_translations.RegisterDefaultTranslations(validators, trans); err != nil {
		log.Fatal(err)
	}

	_ = validators.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	userRepo := users.NewUserRepository(database.GetDB())
	userServices := users.NewUserServices(userRepo)
	users.NewAuthenticationController(router, userServices, validators, trans).MakeUserHandler()

	files.NewFileController(router).MakeUserHandler()

	// posts
	postRepo := posts.NewMySqlPostRepository(database.GetDB())
	postService := posts.NewPostService(postRepo)
	posts.NewPostController(router, postService).MakePostHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
