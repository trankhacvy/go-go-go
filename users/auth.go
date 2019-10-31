package users

import (
	"encoding/json"
	"fmt"
	"go-go-go/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type AuthenticationController struct {
	services   *UserServices
	validators *validator.Validate
	router     *mux.Router
}

func NewAuthenticationController(r *mux.Router, s *UserServices, v *validator.Validate) *AuthenticationController {
	return &AuthenticationController{
		services:   s,
		validators: v,
		router:     r,
	}
}

func (controller *AuthenticationController) Register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		registerParams := RegisterValidator{}
		
		if err := json.NewDecoder(r.Body).Decode(&registerParams); err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}

		fmt.Println("Register handler", registerParams)

		if err := controller.validators.Struct(registerParams); err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}

		user := User{
			Username: registerParams.Username,
			Password: registerParams.Password,
			Email: registerParams.Email,
			Firstname: registerParams.Firstname,
			Lastname: registerParams.Lastname,
		}

		_, err := controller.services.Register(&user)
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "register successfully!")
		resp["user"] = user
		utils.Respond(w, resp)

	})
}

func (controller *AuthenticationController) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginParams := LoginValidator{}

		if err := json.NewDecoder(r.Body).Decode(&loginParams); err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}

		if err := controller.validators.Struct(loginParams); err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, err.Error()))
			return
		}

		fmt.Println("Login handler", loginParams)

		user := &User{
			Username: loginParams.Username,
			Password: loginParams.Password,
		}

		user, err := controller.services.Login(user.Username, user.Password)
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusUnauthorized, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "login successfully!")
		resp["user"] = user
		utils.Respond(w, resp)

	})
}

func (controller *AuthenticationController) FetchAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := controller.services.FetchAll()
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "Get book info successfully!")
		resp["users"] = users
		utils.Respond(w, resp)

	})
}

func (controller *AuthenticationController) Profile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user")

		user, err := controller.services.Profile(userID.(uint))
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "success")
		resp["user"] = user
		utils.Respond(w, resp)

	})
}

func (controller *AuthenticationController) UserProfile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := strconv.ParseUint(vars["id"], 10, 32)

		user, err := controller.services.Profile(uint(userID))
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		resp := utils.Message(http.StatusOK, "success")
		resp["user"] = user
		utils.Respond(w, resp)

	})
}

func (controller *AuthenticationController) MakeUserHandler() {
	controller.router.Handle("/api/user", controller.Profile()).Methods("GET")
	controller.router.Handle("/api/user/{id}", controller.UserProfile()).Methods("GET")
	controller.router.Handle("/api/user/register", controller.Register()).Methods("POST")
	controller.router.Handle("/api/user/login", controller.Login()).Methods("POST")
}
