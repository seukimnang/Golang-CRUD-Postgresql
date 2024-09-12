package model

import (
	"time"
)

type Users struct {

	Id int `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(50)"`
	Created time.Time 
	Status bool 
}