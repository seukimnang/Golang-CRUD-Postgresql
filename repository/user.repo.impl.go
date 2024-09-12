package repository

import (
	"errors"
	"github.com/kim/go-crud/dto/request"
	"github.com/kim/go-crud/helper"
	"github.com/kim/go-crud/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl (Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (t *UserRepositoryImpl) Save(Users model.Users) {
	result := t.Db.Create(&Users)
	helper.ErrorPanic(result.Error)
}

func (t *UserRepositoryImpl) Update(Users model.Users){
	var updateUser = request.UpdateUsersRequest{
		Id: Users.Id,
		Username: Users.Username,
		Email: Users.Email,
	}
	result := t.Db.Model(&Users).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func (t *UserRepositoryImpl) Delete(UserId int){
	var users model.Users
	result := t.Db.Where("id = ?", UserId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

func (t *UserRepositoryImpl) FindById(userId int) (model.Users, error) {
	var user model.Users
	result := t.Db.Find(&user, userId)
	if result != nil {
		return user, nil 
	} else {
		return user, errors.New("User not found")
	}
}

func (t *UserRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	results := t.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}