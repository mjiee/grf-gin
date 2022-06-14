# Scaffold-gin

![GitHub](https://img.shields.io/github/license/mjiee/scaffold-gin)

Scaffold-gin是一个基于go语言gin框架的web案例，专注于前后端分离的业务场景。目的是为初学者提供一个清晰的web项目主线逻辑，对基础功能实现封装。

## Deployment

快速部署项目

```bash
# 获取项目
git clone https://github.com/mjiee/scaffold-gin.git

# 编译项目
go build

# 修改配置文件
vi ./conf/default.yaml

# 初始化项目
./scaffold-gin check -c ./conf/default.yaml
./scaffold-gin init -c ./conf/default.yaml

# 运行项目
./scaffold-gin run -c ./conf/default.yaml
```

使用docker一键部署项目

```bash
docker-compose -f ./docker-compose.yaml build
```

## Documentation

项目目录结构：

```bash

```

项目依赖:

```bash
github.com/spf13/cobra  # 命令行
github.com/spf13/viper  # 配置管理
go.uber.org/zap  # 日志服务
github.com/natefinch/lumberjack  # 日志切割服务
github.com/go-playground/validator/v10  # 数据校验服务
gorm.io/gorm  # 数据库orm
gorm.io/driver/mysql  # mysql驱动
gorm.io/gorm/logger  # 数据库日志
github.com/go-redis/redis/v8  # redis客户端
github.com/gin-gonic/gin  # web服务框架
github.com/golang-jwt/jwt/v4  # jwt认证服务
github.com/swaggo/files  # swagger内置文件
github.com/swaggo/gin-swagger  # swagger gin中间件
github.com/swaggo/swag/cmd/swag  # 生成api文档, 需要安装
github.com/google/wire/cmd/wire  # 依赖注入, 需要安装
```

其他服务:

```bash
# swagger文档生成
swag init --parseDependency --output ./app/docs
go build -tags "dev"
# api访问地址: /api/v1/swagger/index.html

# 依赖代码生成
wire
```
