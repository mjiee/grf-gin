package api

import (
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

// @Summary "Register"
// @Description "用户注册"
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param name formData string true "用户名"
// @Param phone formData string true "手机号"
// @Param password formData string true "用户密码"
// @Success 200 {object} model.User '用户详情'
// @Failure 10001 {object} response.Response '错误信息'
// @Router /auth/register [post]
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

// @Summary "Login"
// @Description "用户登录"
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param phone formData string true "手机号"
// @Param password formData string true "用户密码"
// @Success 200 {object} lib.TokenOutput "登录成功"
// @Failure 10002 {object} response.Response "错误信息"
// @Router /auth/login [get]
func (h *AuthHandler) Login(c *gin.Context) {
	var form lib.Login
	if err := c.ShouldBindQuery(&form); err != nil {
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
			response.Success(c, tokenData)
		}
	}
}

// @Summary "RenewToken"
// @Description "更新Token"
// @Tags auth
// @Produce application/json
// @Success 200 {object} lib.TokenOutput "更新成功"
// @Failuer 0 {string} string 'token已更新'
// @Failure 10002 {object} response.Response "错误信息"
// @Router /auth/renewToken [get]
func (h *AuthHandler) RenewToken(c *gin.Context) {
	headerAuth := c.GetHeader("Authorization")
	claims, token, err := h.jwtSrv.RequestAuth(h.appName, headerAuth)

	if err != nil {
		response.Failure(c, apperr.TokenError, err.Error())
		c.Abort()
		return
	}

	// token renew
	if !h.jwtSrv.IsInBlackList(token.Raw) {
		user, err := h.userSrv.GetUserInfo(claims.ID)
		if err != nil {
			response.Failure(c, apperr.TokenError, err.Error())
			c.Abort()
			return
		} else {
			tokenData, _ := h.jwtSrv.GenToken(h.appName, user)
			_ = h.jwtSrv.JoinBlackList(token.Raw)
			response.Success(c, tokenData)
		}
	} else {
		response.Success(c, nil)
	}
}
