package repository

import (
	"github.com/kim/go-crud/model"
)

type UserRepository interface {
	Save (Users model.Users)
	Update (Users model.Users)
	Delete (UserId int)
	FindById (UserId int) (Users model.Users, err error)
	FindAll() []model.Users
}