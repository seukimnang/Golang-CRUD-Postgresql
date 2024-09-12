package request

type CreateUsersRequest struct {
	Username string `validate:"required,min=1,max=200" json:"name"`
	Password string `validate:"required,min=1,max=200" json:"password"`
	Email  string `validate:"required,min=1,max=200" json:"email"`
}

type UpdateUsersRequest struct {
	Id       int    `validate:"required"`
	Username string `validate:"required,max=200,min=1" json:"name"`
	Email string `validate:"required,min=1,max=200" json:"email"`
}
