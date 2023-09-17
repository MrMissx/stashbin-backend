package response

import (
	"net/http"

	"gorm.io/gorm"
)

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewError(msg string) *ErrorResponse {
	return &ErrorResponse{Ok: false, Message: msg}
}

func GormErrorToResponse(err error) (int, *ErrorResponse) {
	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound, ErrDocumentNotFound
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}

var (
	ErrInternalServerError = NewError("internal server error")
	ErrDocumentNotFound    = NewError("document not found")
	ErrContentEmpty        = NewError("content is empty")
	ErrNoKeyQuery          = NewError("no document key provided")
)
