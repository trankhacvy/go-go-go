package posts

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"go-go-go/utils"
)

type PostController struct {
	router *mux.Router
	service *PostService
}

func NewPostController(r *mux.Router, s *PostService) *PostController {
	return &PostController{router: r, service: s}
}

func (controller *PostController) AllPosts() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		posts, err := controller.service.AllPosts()
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "")
		resp["posts"] = posts
		utils.Respond(res, resp)
	})
}

func (controller *PostController) PostById() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		postID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}
		post, err := controller.service.repo.FindById(uint(postID))
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}
		resp := utils.Message(http.StatusOK, "success")
		resp["post"] = post
		utils.Respond(res, resp)
	})
}

func (controller *PostController) NewPost() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		post := Post{}
		if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
			utils.Respond(res, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}
		userID := req.Context().Value("user")
		post.UserID = userID.(uint)

		_, err := controller.service.Create(&post)
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "")
		resp["post"] = post
		utils.Respond(res, resp)
	})
}

func (controller *PostController) UpdatePost() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		postID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}

		post := Post{}
		if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
			utils.Respond(res, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}
		post.ID = uint(postID)

		_, err = controller.service.Update(&post)
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "")
		resp["post"] = post
		utils.Respond(res, resp)
	})
}

func (controller *PostController) DeletePost() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		postID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}
		
		_, err = controller.service.Delete(uint(postID))
		if err != nil {
			utils.Respond(res, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "success")
		utils.Respond(res, resp)
	})
}

func (controller *PostController) MakePostHandler() {
	controller.router.Handle("/api/post", controller.AllPosts()).Methods("GET")
	controller.router.Handle("/api/post/{id}", controller.PostById()).Methods("GET")
	controller.router.Handle("/api/post", controller.NewPost()).Methods("POST")
	controller.router.Handle("/api/post/{id}", controller.UpdatePost()).Methods("PUT")
	controller.router.Handle("/api/post/{id}", controller.DeletePost()).Methods("DELETE")
}