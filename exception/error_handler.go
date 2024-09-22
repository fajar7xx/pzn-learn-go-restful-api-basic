package exception

import (
	"fajar7xx/pzn-golang-restful-api/helper"
	"fajar7xx/pzn-golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(w, request, err) {
		return
	}

	if validationErrors(w, request, err) {
		return
	}

	internalServerError(w, request, err)
}

func internalServerError(w http.ResponseWriter, request *http.Request, err interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	WebResponse := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Status:  "Internal server error",
		Message: "Internal Server error",
		Data:    err,
	}

	helper.WriteToResponseBody(w, WebResponse)
}

func notFoundError(w http.ResponseWriter, request *http.Request, err interface{}) bool {
	// kita konversi
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		WebResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Not Found",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(w, WebResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(w http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		WebResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    exception.Error(),
		}

		helper.WriteToResponseBody(w, WebResponse)
		return true
	} else {
		return false
	}
}
