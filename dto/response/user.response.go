package response

type UsersResponse struct {
	Id       int    `json:"id"`
	Username string `json:"name"`
	Email string `validate:"required,min=1,max=200" json:"email"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
