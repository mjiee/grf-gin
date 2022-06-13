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
github.com/google/wire/cmd/wire  # 依赖注入, 需要安装


```

其他服务:

```bash
# swagger文档生成
swag init --parseDependency --output ./bin/docs

# 依赖代码文档生成
wire
```