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

// @Summary "SignUp"
// @Description "用户注册"
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param name formData string true "用户名"
// @Param phone formData string true "手机号"
// @Param password formData string true "用户密码"
// @Param category formData bool false "人员类别: manager: true | user: false"
// @response default {object} response.Response "响应包装"
// @Success 200 {object} model.User '用户详情'
// @Router /auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var form lib.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Failure(c, apperr.ValidateErr, request.GetErrorMsg(form, err))
		return
	}

	if user, err := h.userSrv.SignUp(form); err != nil {
		response.Failure(c, apperr.BusinessErr, err.Error())
		return
	} else {
		if tokenData, err := h.jwtSrv.GenToken(h.appName, form.IsAdmin, user); err != nil {
			response.Failure(c, apperr.BusinessErr, err.Error())
			return
		} else {
			response.Success(c, tokenData)
		}
	}
}

// loginResponse 登陆成功响应数据
type loginResponse struct {
	user  any
	token *lib.TokenOutput
}

// @Summary "SignIn"
// @Description "用户登录"
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param phone query string true "手机号"
// @Param password query string true "用户密码"
// @Param category query bool false "人员类别: manager: true | user: false"
// @response default {object} response.Response "响应包装"
// @Success 200 {object} loginResponse "登录成功后response中的data数据"
// @Router /auth/signin [get]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var form lib.Login
	if err := c.ShouldBindQuery(&form); err != nil {
		response.Failure(c, apperr.ValidateErr, request.GetErrorMsg(form, err))
		return
	}

	if user, err := h.userSrv.SignIn(form); err != nil {
		response.Failure(c, apperr.BusinessErr, err.Error())
		return
	} else {
		if tokenData, err := h.jwtSrv.GenToken(h.appName, form.IsAdmin, user); err != nil {
			response.Failure(c, apperr.BusinessErr, err.Error())
			return
		} else {
			response.Success(c, loginResponse{user, tokenData})
		}
	}
}

// @Summary "RenewToken"
// @Description "更新Token"
// @Tags auth
// @Produce application/json
// @Param category query bool false "人员类别: manager: true | user: false"
// @response default {object} response.Response "响应包装"
// @Success 200 {object} lib.TokenOutput "更新成功, 如果data为空, 表示token已更新过"
// @Router /auth/renewToken [get]
func (h *AuthHandler) RenewToken(c *gin.Context) {
	headerAuth := c.GetHeader("Authorization")
	claims, token, err := h.jwtSrv.RequestAuth(h.appName, headerAuth)

	if err != nil {
		response.Failure(c, apperr.TokenErr, err.Error())
		c.Abort()
		return
	}

	// token renew
	if !h.jwtSrv.IsInBlackList(token.Raw) {
		user, err := h.userSrv.GetUserInfo(claims.ID, claims.IsAdmin)
		if err != nil {
			response.Failure(c, apperr.TokenErr, err.Error())
			c.Abort()
			return
		} else {
			tokenData, _ := h.jwtSrv.GenToken(h.appName, claims.IsAdmin, user)
			_ = h.jwtSrv.JoinBlackList(token.Raw)
			response.Success(c, tokenData)
		}
	} else {
		response.Success(c, nil)
	}
}
