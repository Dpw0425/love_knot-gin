package response

import (
	"github.com/gin-gonic/gin"
	myErr "love_knot/pkg/error"
)

type NormalResponse struct {
	Status  int         `json:"code"` // 响应状态码
	Data    interface{} `json:"data"` // 返回内容
	Message string      `json:"msg"`  // 返回消息
}

type ErrorResponse struct {
	ID      string `json:"error_type"` // 错误码类型
	Status  int    `json:"code"`       // 响应状态码
	Message string `json:"message"`    // 返回消息
}

func normal(c *gin.Context, status int, data interface{}, message string) {
	c.Abort()
	c.JSON(status, &NormalResponse{
		Status:  status,
		Data:    data,
		Message: message,
	})
}

func NorResponse(c *gin.Context, status int, data interface{}, message string) {
	normal(c, status, data, message)
}

func err(c *gin.Context, id string, status int, message string) {
	c.Abort()
	c.JSON(status, ErrorResponse{
		ID:      id,
		Status:  status,
		Message: message,
	})
}

func ErrResponse(c *gin.Context, e error) {
	newErr := myErr.Parse(e.Error())

	err(c, newErr.ID, newErr.Code, newErr.Detail)
}
