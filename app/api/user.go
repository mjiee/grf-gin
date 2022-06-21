package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/scaffold-gin/app/lib"
	"github.com/mjiee/scaffold-gin/app/pkg/apperr"
	"github.com/mjiee/scaffold-gin/app/pkg/response"
)

// UserHandler 用户处理器
type UserHandler struct {
	userSrv *lib.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userSrv *lib.UserService) *UserHandler {
	return &UserHandler{userSrv}
}

// GetUserInfo
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	user, err := h.userSrv.GetUserInfo(c.MustGet("id").(string))
	if err != nil {
		response.Failure(c, apperr.BusinessErr, err.Error())
		return
	}
	response.Success(c, user)
}
