package service

import (
	"github.com/kim/go-crud/dto/request"
	"github.com/kim/go-crud/dto/response"
)

type UsersServices interface {
	Create(users request.CreateUsersRequest)
	Update(users request.UpdateUsersRequest)
	Delete(usersId int)
	FindById(usersId int) response.UsersResponse
	FindAll() []response.UsersResponse
}
