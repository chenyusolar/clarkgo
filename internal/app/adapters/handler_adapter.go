package adapters

import (
	"context"

	"github.com/clarkgo/clarkgo/pkg/framework"
	"github.com/cloudwego/hertz/pkg/app"
)

// HertzToFramework 将Hertz的处理函数转换为framework的处理函数
func HertzToFramework(handler app.HandlerFunc) framework.HandlerFunc {
	return func(ctx interface{}) {
		// 这里假设framework.HandlerFunc接受一个interface{}参数
		// 实际实现可能需要根据framework包的定义进行调整
		handler(context.Background(), ctx.(*app.RequestContext))
	}
}

// ControllerToFramework 将控制器方法转换为framework的处理函数
func ControllerToFramework(handler func(ctx interface{})) framework.HandlerFunc {
	return func(ctx interface{}) {
		handler(ctx)
	}
}
