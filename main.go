package main

import (
	"github.com/mjiee/scaffold-gin/cmd"
)

// @title scaffold-gin
// @version 0.0.1
// @description "这是一个web项目。请求成功的消息都包装在: '{"err_code":"0","data":"xx","message":""}', 下文中的成功返回都只写了data数据类型. 请求失败的错误码为err_code值"

// @contact.name mjiee
// @contact.url https://mjiee.top
// @contact.email mminjjie@gmail.com

// @tag.name Scaffold-gin
// @tag.description "web项目api"
// @tag.docs.url https://github.com/mjiee/scaffold-gin
// @tag.docs.description "项目仓库地址"

// @license.name MIT
// @license.url http://www.mit.org

// @host http://127.0.0.1:8080
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
