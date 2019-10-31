package main

// import (
// 	"fmt"
// 	"go-go-go/controllers"
// 	database "go-go-go/db"

// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// 	"go-go-go/middlewares"
// 	"go-go-go/posts"
// 	"go-go-go/users"
// 	"net/http"
// 	"os"

// 	"gopkg.in/go-playground/validator.v9"
// )



// func main() {
// 	database.Init()

// 	router := mux.NewRouter()
// 	router.Use(middlewares.JwtAuthentication)

// 	validators := validator.New()

// 	userRepo := users.NewUserRepository(database.GetDB())
// 	userServices := users.NewUserServices(userRepo)
// 	users.NewAuthenticationController(router, userServices, validators).MakeUserHandler()

// 	controllers.NewFileController(router).MakeUserHandler()

// 	// posts
// 	postRepo := posts.NewMySqlPostRepository(database.GetDB())
// 	postService := posts.NewPostService(postRepo)
// 	posts.NewPostController(router, postService).MakePostHandler()

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "9999" //localhost
// 	}

// 	err := http.ListenAndServe(":"+port, router)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// }

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// func init() {
// 	e := godotenv.Load()
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", TestHandler)

	port := "9999"

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Category: 123")
}
