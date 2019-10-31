package users

import (
	"errors"
	"os"
	// "fmt"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	repo UserRepository
}

func NewUserServices(repo UserRepository) *UserServices {
	return &UserServices{repo: repo}
}

func (s *UserServices) Register(user *User) (bool, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err := s.repo.Create(user)
	if err != nil {
		return false, err
	}
	tk := Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""

	return true, nil
}

func (s *UserServices) Login(username, password string) (*User, error) {

	user := s.repo.FindByUsername(username)
	if user == nil || user.ID == 0 {
		return nil, errors.New("Invalid Username")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid Password")
	}

	user.Password = ""

	//Create JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	return user, nil
}

func (s *UserServices) Profile(id uint) (*User, error) {
	return s.repo.FindById(id)
}

func (s *UserServices) FetchAll() ([]*User, error) {
	return s.repo.GetAll()
}
