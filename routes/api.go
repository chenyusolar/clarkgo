package routes

import (
	controllers "github.com/clarkgo/clarkgo/app/Http/Controllers"
	middleware "github.com/clarkgo/clarkgo/app/Http/Middleware"
	"github.com/clarkgo/clarkgo/pkg/framework"
)

func APIRoutes(app *framework.Application) {
	userController := controllers.NewUserController(app)

	app.RegisterRoutes(func(r *framework.Router) {
		// 公开路由
		r.POST("/register", userController.Register)
		r.POST("/login", userController.Login)

		// 需要认证的路由
		authGroup := r.Group("/user", middleware.JWTMiddleware(app))
		{
			authGroup.GET("/profile", userController.Profile)
		}
	})
}
