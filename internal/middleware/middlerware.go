package middleware

import (
	"github.com/gin-gonic/gin"
	"newframework/pkg/errors"
	"newframework/pkg/errors/ecode"
	"newframework/pkg/response"
)

// middleware 实现Router接口
// 便于服务启动时加载, middleware本质跟handler无区别
type middleware struct {
}

func NewMiddleware() *middleware {
	return &middleware{}
}

// Load 注册中间件和公共路由
func (m *middleware) Load(g *gin.Engine) {
	// 注册中间件
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		response.JSON(c, errors.WithCode(ecode.NotFoundErr, "404 not found!"), nil)
	})
}
