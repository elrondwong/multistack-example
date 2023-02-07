# 项目介绍

这是一个使用pulumi创建腾讯云存储桶的示例项目，使用gin框架提供rest api服务

# 服务架构

- gin提供http服务
- 使用 pulumi automation api 操作 pulumi
- 使用 pulumi 腾讯云插件 操作腾讯云资源

# 目录结构

```bash
.
|-- Pulumi.yaml -- pulumi项目初始化配置
|-- README.md
|-- cmd
|   `-- multistack
|       `-- main.go -- 项目入口文件
|-- go.mod
|-- go.sum
`-- pkg
    |-- apiserver
    |   |-- controller -- controller
    |   |   `-- qclouds3.go
    |   |-- domain -- pulumi program
    |   |   `-- qs3
    |   |       `-- service.go
    |   `-- route -- 路由
    |       |-- init.go
    |       |-- routes.go
    |       `-- tencents3
    |           `-- qcloud.go
    |-- consts
    |   `-- consts.go
    `-- infra
        `-- log
```

- 服务地址

`cmd/multistack/main.go:12`

- 腾讯云凭证

`pkg/apiserver/domain/qs3/service.go:28-30`

- pulumi日志

`pkg/apiserver/domain/qs3/service.go:46`

# 使用

> 安装所有pulumi相关插件，不走代理非常慢

- 安装pulumi

```bash
curl -fsSL https://get.pulumi.com | sh
```

- 安装pulumi tencentcloud plugin

```bash
pulumi plugin install resource tencentcloud v0.1.2
```

- 登陆pulumi

```bash
pulumi login
```

- 下载项目依赖包

```bash
go mod tidy
```

- 运行项目

```bash
go run cmd/multistack/main.go
# command-line-arguments
ld: warning: -no_pie is deprecated when targeting new OS versions
ld: warning: non-standard -pagezero_size is deprecated when targeting macOS 13.0 or later
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /multistack/apis/v1/qs3bucket --> github.com/elrondwong/multistack-example/pkg/apiserver/controller.(*Qclouds3Controller).CreateBucket-fm (5 handlers)
2023/02/07 15:40:20 Run server at 0.0.0.0:8800
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8800
```

- 测试

```bash
curl --location \
  --request POST '127.0.0.1:8800/multistack/apis/v1/qs3bucket' \
  --header 'Content-Type: application/json' \
  --data-raw '{"name": "test007"}'
```

这样就在腾讯云上创建了一个 test007-appid的存储桶，在pulumi上创建一个project和stack
