package service

import (
	"context"
	"database/sql"
	"fajar7xx/pzn-golang-restful-api/exception"
	"fajar7xx/pzn-golang-restful-api/helper"
	"fajar7xx/pzn-golang-restful-api/model/domain"
	"fajar7xx/pzn-golang-restful-api/model/web"
	"fajar7xx/pzn-golang-restful-api/repository"

	"github.com/go-playground/validator/v10"
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

type PostServiceImpl struct {
	PostRepository repository.PostRepository //karena interface jadi gak perlu pointer*
	DB             *sql.DB                   //karena struct maka tambahkan pointer*
	Validate       *validator.Validate
}

func NewPostService(postRespository repository.PostRepository, DB *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRespository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse {
	// validation
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	// untuk rollback
	// defer func(){
	// 	err := recover()
	// 	if err != nil{
	// 		tx.Rollback()
	// 		panic(err)
	// 	}else{
	// 		tx.Commit()
	// 	}
	// }()
	// refactoring ke
	defer helper.CommitOrRollback(tx)

	post := domain.Post{
		Name: request.Name,
		Post: request.Post,
	}

	post = service.PostRepository.Create(ctx, tx, post)
	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	// validation
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindyById(ctx, tx, request.Id)
	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	post.Name = request.Name
	post.Post = request.Post

	post = service.PostRepository.Update(ctx, tx, post)
	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindyById(ctx, tx, postId)
	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PostRepository.Delete(ctx, tx, post)
}

func (service *PostServiceImpl) FindById(ctx context.Context, postId int) web.PostResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindyById(ctx, tx, postId)
	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) FindAll(ctx context.Context) []web.PostResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	posts := service.PostRepository.FindAll(ctx, tx)

	// var postResponses []web.PostResponse
	// for _, post := range posts{
	// 	postResponses = append(postResponses, helper.ToPostResponse(post))
	// }

	return helper.ToPostResponses(posts)
}
