package routes

import (
	"github.com/clarkgo/clarkgo/pkg/framework"
)

// TestRoutes 测试路由
func TestRoutes(app *framework.Application) {
	app.RegisterRoutes(func(r *framework.Router) {
		r.GET("/test", func(ctx framework.Context) {
			ctx.JSON(200, map[string]interface{}{
				"message": "测试成功",
			})
		})
	})
}
