package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/grf-gin/app/lib"
	"github.com/mjiee/grf-gin/app/pkg/apperr"
	"github.com/mjiee/grf-gin/app/pkg/response"
)

// OssHandler Oss文件处理
type OssHandler struct {
	ossSrv *lib.OssService
}

func NewOssHandler(ossSrv *lib.OssService) *OssHandler {
	return &OssHandler{ossSrv}
}

// @Summary 'GetStsToken'
// @description '获取Sts临时授权'
// @Tags sts
// @Security ApiKeyAuth
// @response default {object} response.Response "响应包装"
// @Success 200 {object} sts.Credentials "token信息"
// @Router /oss/getStsToken [get]
func (h *OssHandler) GetStsToken(c *gin.Context) {
	token, err := h.ossSrv.GenStsToken()
	if err != nil {
		response.Failure(c, apperr.BusinessErr, "获取sts授权失败")
		return
	}
	response.Success(c, token)
}
