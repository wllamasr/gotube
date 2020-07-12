package errors

import "net/http"

type RestError struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Error   string      `json:"error"`
}

func BadRequestError(message interface{}) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func InternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func UnauthorizedError() *RestError {
	return &RestError{
		Message: "Unauthorized",
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}
