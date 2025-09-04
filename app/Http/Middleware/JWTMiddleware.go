package middleware

import (
	"github.com/clarkgo/clarkgo/pkg/framework"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(app *framework.Application) app.HandlerFunc {
	return func(ctx *app.RequestContext) {
		// JWT验证逻辑
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, map[string]string{"error": "Authorization header missing"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(app.Config.GetString("app.key")), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(401, map[string]string{"error": "Invalid token"})
			return
		}

		ctx.Next()
	}
}
