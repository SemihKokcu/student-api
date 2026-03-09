package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorType string

const (
	ErrTypeValidation ErrorType = "ValidationError"
	ErrTypeNotFound   ErrorType = "NotFound"
	ErrTypeBadRequest ErrorType = "BadRequest"
	ErrTypeInternal   ErrorType = "InternalError"
)

type ErrorResponse struct {
	Type    ErrorType   `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{Data: data})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, SuccessResponse{Data: data})
}

func ValidationError(c *gin.Context, fields interface{}) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Type:    ErrTypeValidation,
		Message: "Validasyon hatası oluştu.",
		Data:    fields,
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Type:    ErrTypeNotFound,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Type:    ErrTypeBadRequest,
		Message: message,
	})
}

func InternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Type:    ErrTypeInternal,
		Message: message,
	})
}
