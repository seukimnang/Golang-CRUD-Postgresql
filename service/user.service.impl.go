package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/kim/go-crud/dto/request"
	"github.com/kim/go-crud/dto/response"
	"github.com/kim/go-crud/helper"
	"github.com/kim/go-crud/model"
	"github.com/kim/go-crud/repository"
)

type UsersServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(UserRepository repository.UserRepository, validate *validator.Validate) UsersServices {
	return &UsersServiceImpl{
		UserRepository: UserRepository,
		Validate:       validate,
	}
}

func (t *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := t.Validate.Struct(users)
	helper.ErrorPanic(err)
	userModel := model.Users{
		Username: users.Username,
		Password: users.Password,
		Email:    users.Email,
	}
	t.UserRepository.Save(userModel)
}

func (t UsersServiceImpl) Update(user request.UpdateUsersRequest) {
	userData, err := t.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Username = user.Username
	userData.Email = user.Email
	t.UserRepository.Update(userData)
}

func (t UsersServiceImpl) Delete(tagId int) {
	t.UserRepository.Delete(tagId)
}

func (t UsersServiceImpl) FindById(userId int) response.UsersResponse {
	userData, err := t.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	tagResponse := response.UsersResponse{
		Id:       userData.Id,
		Username: userData.Username,
		Email:    userData.Email,
	}
	return tagResponse
}

func (t UsersServiceImpl) FindAll() []response.UsersResponse {
	result := t.UserRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			Id:       value.Id,
			Username: value.Username,
			Email:    value.Email,
		}
		users = append(users, user)
	}
	return users
}
