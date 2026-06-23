package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/errorcode"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationData struct {
	Items    interface{} `json:"items"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type ListData struct {
	Items interface{} `json:"items"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    errorcode.Success,
		Message: errorcode.Message(errorcode.Success),
		Data:    data,
	})
}

func Error(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: errorcode.Message(code),
		Data:    struct{}{},
	})
}

func ErrorWithMessage(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    struct{}{},
	})
}
