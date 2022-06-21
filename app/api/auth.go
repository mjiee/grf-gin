package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mjiee/scaffold-gin/app/lib"
	"github.com/mjiee/scaffold-gin/app/pkg/apperr"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"github.com/mjiee/scaffold-gin/app/pkg/request"
	"github.com/mjiee/scaffold-gin/app/pkg/response"
)

// AuthHander 用户认证处理器
type AuthHandler struct {
	appName string
	jwtSrv  *lib.JwtService
	userSrv *lib.UserService
}

// NewAuthHandler 创建新的AuthHandler
func NewAuthHandler(cfg *conf.Config, jwtSrv *lib.JwtService, userSrv *lib.UserService) *AuthHandler {
	return &AuthHandler{cfg.App.Name, jwtSrv, userSrv}
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
		return
	} else {
		response.Success(c, user)
	}
}

// Login
func (h *AuthHandler) Login(c *gin.Context) {
	var form lib.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Failure(c, apperr.ValidateErr, request.GetErrorMsg(form, err))
		return
	}

	if user, err := h.userSrv.Login(form); err != nil {
		response.Failure(c, apperr.BusinessErr, err.Error())
		return
	} else {
		if tokenData, err := h.jwtSrv.GenToken(h.appName, user); err != nil {
			response.Failure(c, apperr.BusinessErr, err.Error())
			return
		} else {
			c.Header("New-Token", tokenData.AccessToken)
			c.Header("New-Expires-In", strconv.Itoa(tokenData.ExpiresAt))
			response.Success(c, "login success")
		}
	}
}
