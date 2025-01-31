package exceptions

import (
	"net/http"

	"zen-test/app/helpers"
	"zen-test/app/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")

		statusCode := http.StatusBadRequest
		statusText := "Bad Request"

		if request.URL.Path == "/login" && request.Method == "POST" {
			statusCode = http.StatusUnauthorized
			statusText = "Unauthorized"
		}

		writer.WriteHeader(statusCode)

		webResponse := web.WebResponse{
			Code:   statusCode,
			Status: statusText,
			Data:   exception.Error(),
		}

		helpers.WriteResponseBody(writer, webResponse)
		return true
	}
	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helpers.WriteResponseBody(writer, webResponse)
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helpers.WriteResponseBody(writer, webResponse)
}
