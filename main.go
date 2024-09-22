package main

import (
	"fajar7xx/pzn-golang-restful-api/controller"
	"fajar7xx/pzn-golang-restful-api/database"
	"fajar7xx/pzn-golang-restful-api/exception"
	"fajar7xx/pzn-golang-restful-api/helper"
	"fajar7xx/pzn-golang-restful-api/repository"
	"fajar7xx/pzn-golang-restful-api/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.NewDB()
	validate := validator.New()

	postRepository := repository.NewPostRepository()
	postService := service.NewPostService(postRepository, db, validate)
	postController := controller.NewPostController(postService)

	router := httprouter.New()

	router.GET("/api/v1/posts", postController.FindAll)
	router.GET("/api/v1/posts/:postId", postController.FindById)
	router.POST("/api/v1/posts", postController.Create)
	router.PUT("/api/v1/posts/:postId", postController.Update)
	router.DELETE("/api/v1/posts/:postId", postController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	log.Println("Server running on: ", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
