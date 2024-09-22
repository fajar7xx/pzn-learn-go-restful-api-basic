package web

type PostCreateRequest struct {
	Name string `validate:"required,max=255" json:"name"`
	Post string `validate:"required,max=255" json:"post"`
}

type PostUpdateRequest struct {
	Id   int    `validate:"required"` //karena tidak masuk dalam payload makanya tidak dibuat jsonnya
	Name string `validate:"required,max=255" json:"name"`
	Post string `validate:"required,max=255" json:"post"`
}
