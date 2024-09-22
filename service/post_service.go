package service

import (
	"context"
	"database/sql"
	"fajar7xx/pzn-golang-restful-api/model/web"
	"fajar7xx/pzn-golang-restful-api/repository"
)

// service function mengikuti api yang di buat
// contractnya
type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse
	Delete(ctx context.Context, postId int)
	FindById(ctx context.Context, postId int) web.PostResponse
	FindAll(ctx context.Context) []web.PostResponse
}

type PostServiceImpl struct{
	PostRepository repository.PostRepository //karena interface jadi gak perlu pointer*
	DB *sql.DB //karena struct maka tambahkan pointer*
}