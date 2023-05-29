package response

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors"`
	Status  int         `json:"status"`
}

type EmptyObj struct{}

func ConvertToBaseResponse(message string, status int, data interface{}) BaseResponse {
	return BaseResponse{
		Message: message,
		Data:    data,
		Status:  status,
		Errors:  nil,
	}
}

func ConvertErrorToBaseResponse(message string, status int, data interface{}, err string) BaseResponse {
	splittedError := strings.Split(err, "\n")
	return BaseResponse{
		Message: message,
		Data:    data,
		Status:  status,
		Errors:  splittedError,
	}
}

func Success(ctx *gin.Context, message string, status int, data interface{}) {
	response := ConvertToBaseResponse(message, status, data)
	ctx.JSON(status, response)
}

func Error(ctx *gin.Context, message string, status int, err error) {
	response := ConvertErrorToBaseResponse(message, status, EmptyObj{}, err.Error())
	ctx.JSON(status, response)
}
