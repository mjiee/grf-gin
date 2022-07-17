package main

import (
	"github.com/mjiee/grf-gin/cmd"
)

// @title grf-gin
// @version 0.0.1
// @description "grf-gin是一个基于go语言gin框架的web案例，专注于前后端分离的业务场景。\n1. 登陆后的请求必须添加Authorization. \n2. 响应数据都包装在response.Response中, status非0时表示出现自定义错误, 错误消息在message中。status为0表示正常响应, 数据在data中。"

// @contact.name mjiee
// @contact.url https://github.com/mjiee
// @contact.email mminjjie@gmail.com

// @tag.name grf-gin
// @tag.description "web项目api"
// @tag.docs.url https://github.com/mjiee/grf-gin
// @tag.docs.description "项目仓库地址"

// @license.name MIT
// @license.url http://www.mit.org

// @host http://127.0.0.1:8080
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
