package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	ErrorCodeBadRequest ErrorCode = "BAD_REQUEST"
	ErrorCodeNotFound   ErrorCode = "NOT_FOUND"
	ErrorCodeConflict   ErrorCode = "CONFLICT"
	ErrorCodeInternal   ErrorCode = "INTERNAL_SERVER_ERROR"
)

type AppError struct {
	Message string
	Code    ErrorCode
	Err     error
}

// de ham error mac dinh trong interface
func (ae *AppError) Error() string {
	return ""
}

func NewError(message string, code ErrorCode) error {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func WrapError(err error, message string, code ErrorCode) error {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func ResponseError(ctx *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		status := httpStatusFromCode(appErr.Code)
		response := gin.H{
			"error": appErr.Message,
			"code":  appErr.Code,
		}

		if appErr.Err != nil {
			response["detail"] = appErr.Err.Error()
		}

		ctx.JSON(status, response)
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
		"code":  string(ErrorCodeInternal),
	})
}

func ResponseSuccess(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, gin.H{
		"status": status,
		"data": data,
	})
} 

func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ErrorCodeBadRequest:
		return http.StatusBadRequest
	case ErrorCodeNotFound:
		return http.StatusNotFound
	case ErrorCodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
