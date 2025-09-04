# ClarkGo - 基于Hertz的高性能Go Web框架

ClarkGo是一个仿Goravel风格，基于Cloudwego Hertz框架开发的现代化Go Web框架，旨在提供简单、高效且功能丰富的开发体验。

## 特性

- 🚀 基于Hertz的高性能HTTP服务器
- 🛠️ 模块化设计，易于扩展
- 📦 内置常用组件（数据库、缓存、日志等）
- 🔌 中间件支持
- 🗺️ 灵活的路由管理
- ⚙️ 统一的配置管理
- 📝 详细的请求上下文封装

## 快速开始

### 安装

1. 确保已安装Go (1.18+)
2. 克隆项目：
   ```bash
   git clone https://github.com/clarkgo/clarkgo.git
   cd clarkgo
   ```
3. 初始化依赖：
   ```bash
   go mod tidy
   ```

### 创建新项目

```bash
go run cmd/clarkgo/main.go new project-name
```

### 运行项目

```bash
go run main.go
```

## 项目结构

```
clarkgo/
├── cmd/               # 命令行工具
├── config/            # 配置文件
├── internal/          # 内部应用代码
│   ├── app/          # 应用核心
│   ├── controllers/  # 控制器
│   ├── models/       # 数据模型
│   └── services/     # 业务服务
├── pkg/              # 框架核心
│   ├── cache/        # 缓存系统
│   ├── config/       # 配置管理
│   ├── database/     # 数据库连接
│   ├── framework/     # 框架核心
│   ├── http/         # HTTP客户端
│   └── log/          # 日志系统
└── public/           # 静态文件
```

## 基本用法

### 创建应用实例

```go
app := framework.NewApplication()
```

### 配置管理

```go
app.SetConfigPath("config") // 设置配置目录
app.SetEnv("development")  // 设置环境
app.SetDebug(true)         // 设置调试模式
```

### 路由定义

```go
app.RegisterRoutes(func(router *framework.Router) {
    router.GET("/", func(ctx context.Context, c *framework.RequestContext) {
        c.String(200, "Hello ClarkGo!")
    })
    
    api := router.Group("/api")
    {
        api.GET("/users", UserController.Index)
        api.POST("/users", UserController.Store)
    }
})
```

### 中间件

```go
app.RegisterMiddleware(
    framework.Cors(),
    framework.Logger(),
    framework.Recovery(),
)
```

### 数据库操作

```go
// 获取数据库连接
db := database.GetDB()

// 查询示例
var user User
db.First(&user, 1)
```

### 日志记录

```go
log.Info("This is an info message")
log.Error("This is an error message")
```

## 配置示例

项目使用`.env`文件管理环境变量，请复制`.env.example`并重命名为`.env`，然后修改相应配置：

```bash
cp .env.example .env
```

`.env`文件示例：

```ini
# 应用配置
APP_ENV=development
APP_DEBUG=true

# 服务器配置
SERVER_HOST=0.0.0.0
SERVER_PORT=8888

# 数据库配置
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=clarkgo
DB_USERNAME=root
DB_PASSWORD=secret
```

框架会自动加载`.env`文件中的配置，并通过`config.GetEnv()`方法获取：

```go
host := config.GetEnv("SERVER_HOST", "0.0.0.0")
port := config.GetEnvInt("SERVER_PORT", 8888)
```

## 高级功能

### 自定义中间件

```go
func AuthMiddleware() framework.HandlerFunc {
    return func(ctx context.Context, c *framework.RequestContext) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatus(401)
            return
        }
        c.Next(ctx)
    }
}
```

### 数据库迁移

```bash
go run cmd/clarkgo/main.go migrate
```

### 任务队列

```go
queue.Dispatch(NewEmailJob(user))
```

## 贡献指南

欢迎提交Pull Request或Issue。在提交代码前请确保：

1. 代码通过所有测试
2. 遵循项目代码风格
3. 更新相关文档

## 许可证

ClarkGo采用MIT许可证开源。

## 联系方式

如有任何问题，请联系：clark@example.com