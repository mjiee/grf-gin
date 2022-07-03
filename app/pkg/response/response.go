package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 自定义响应体
type Response struct {
	Status  int    `json:"status"`  // 响应状态码
	Data    any    `json:"data"`    // 响应数据
	Message string `json:"message"` // 错误描述
}

// Success 成功的响应, 状态码为0
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{0, data, ""})
}

// Failure 失败的响应, 状态码非0
func Failure(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{code, nil, msg})
}
