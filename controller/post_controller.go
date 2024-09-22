package controller

import (
	"fajar7xx/pzn-golang-restful-api/helper"
	"fajar7xx/pzn-golang-restful-api/model/web"
	"fajar7xx/pzn-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PostController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type PostControllerImpl struct {
	PostService service.PostService
}

// kita buat interface dulu
// kemudian kita dapat membuat sebuah function yang dpat mengexpose
// struct postControllerImpl => ini mirip polimorpysm
// dimana kembaliannya PostControllerl interface 
func NewPostController(postService service.PostService) PostController {
	return &PostControllerImpl{
		PostService: postService,
	}
}

func (c *PostControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// baca request json
	// decoder := json.NewDecoder(r.Body)
	// postCreateRequest := web.PostCreateRequest{}
	// err := decoder.Decode(&postCreateRequest)
	// helper.PanicIfError(err)

	// refactoring dari atas
	postCreateRequest := web.PostCreateRequest{}
	helper.ReadFromRequestBody(r, &postCreateRequest)

	postResponse := c.PostService.Create(r.Context(), postCreateRequest)
	webResponse := web.WebResponse{
		Code:    201,
		Status:  "OK",
		Message: "Post has been successfuley created",
		Data:    postResponse,
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)

	// refactoring dari atas
	helper.WriteToResponseBody(w, webResponse)
}

func (c *PostControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// baca request json
	// decoder := json.NewDecoder(r.Body)
	// postUpdateRequest := web.PostUpdateRequest{}
	// err := decoder.Decode(&postUpdateRequest)
	// helper.PanicIfError(err)

	postUpdateRequest := web.PostUpdateRequest{}
	helper.ReadFromRequestBody(r, &postUpdateRequest)

	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)
	postUpdateRequest.Id = id

	postResponse := c.PostService.Update(r.Context(), postUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Post has been successfully updated",
		Data:    postResponse,
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)

	helper.WriteToResponseBody(w, webResponse)
}

func (c *PostControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	c.PostService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:    204,
		Status:  "OK",
		Message: "Post has been successfully deleted",
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(w, webResponse)
}

func (c *PostControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	postResponse := c.PostService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Get post by id",
		Data:    postResponse,
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(w, webResponse)
}

func (c *PostControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postResponses := c.PostService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Get all posts",
		Data:    postResponses,
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err := encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(w, webResponse)
}
