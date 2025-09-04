package controllers

import (
	"github.com/clarkgo/clarkgo/pkg/framework"
	"github.com/cloudwego/hertz/pkg/app"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	App *framework.Application
}

func NewUserController(app *framework.Application) *UserController {
	return &UserController{App: app}
}

func (c *UserController) Register(ctx *app.RequestContext) {
	// 用户注册逻辑
}

func (c *UserController) Login(ctx *app.RequestContext) {
	// 用户登录逻辑
}

func (c *UserController) Profile(ctx *app.RequestContext) {
	// 获取用户信息
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
