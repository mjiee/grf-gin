package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/scaffold-gin/app/lib"
	"github.com/mjiee/scaffold-gin/app/pkg/apperr"
	"github.com/mjiee/scaffold-gin/app/pkg/request"
	"github.com/mjiee/scaffold-gin/app/pkg/response"
)

// AuthHander 用户认证处理器
type AuthHandler struct {
	jwtSrv  *lib.JwtService
	userSrv *lib.UserService
}

// NewAuthHandler 创建新的AuthHandler
func NewAuthHandler(jwtSrv *lib.JwtService, userSrv *lib.UserService) *AuthHandler {
	return &AuthHandler{jwtSrv, userSrv}
}

// Register
func (h *AuthHandler) Register(c *gin.Context) {
	var form lib.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Failure(c, apperr.ValidateErr, request.GetErrorMsg(form, err))
		return
	}

	if user, err := h.userSrv.Register(form); err != nil {
		response.Failure(c, apperr.BusinessErr, err.Error())
	} else {
		response.Success(c, user)
	}
}
