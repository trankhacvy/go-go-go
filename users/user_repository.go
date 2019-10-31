package users

import (
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(user *User) (bool, error)
	Update(user *User) (bool, error)
	// Delete(id uint) (bool, error)
	GetAll() ([]*User, error)
	FindById(id uint) (*User, error)
	FindByUsername(username string) *User
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (userRepo *UserRepositoryImpl) Create(user *User) (bool, error) {
	if err := userRepo.DB.Create(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryImpl) Update(user *User) (bool, error) {
	if err := userRepo.DB.Save(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepositoryImpl) GetAll() ([]*User, error) {
	users := make([]*User, 0)
	if err := userRepo.DB.Find(&users).Error; err != nil {
		return []*User{}, err
	}
	return users, nil
}

func (userRepo *UserRepositoryImpl) FindById(id uint) (*User, error) {
	user := &User{}
	if err := userRepo.DB.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserRepositoryImpl) FindByUsername(username string) *User {
	user := &User{}
	userRepo.DB.Where("username = ?", username).First(user)
	return user
}

