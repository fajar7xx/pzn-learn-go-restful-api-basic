package web

type PostCreateRequest struct {
	Name string
	Post string
}

type PostUpdateRequest struct {
	Id   int
	Name string
	Post string
}
