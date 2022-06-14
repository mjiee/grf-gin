package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 自定义响应体
type Response struct {
	ErrorCode int    `json:"error_code"`
	Data      any    `json:"data"`
	Message   string `json:"message"`
}

// Success 成功的响应, 错误码为0
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{0, data, "ok"})
}

// Failure 失败的响应, 错误码非0
func Failure(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{code, nil, msg})
}
