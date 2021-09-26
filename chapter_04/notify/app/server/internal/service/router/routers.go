package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"notify-server/internal/service/handle"
)

var _ IRouter = (*Router)(nil)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
}

// Router 路由管理器
type Router struct {
	ManageApi *handle.UserService
} // end

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	v1 := g.Group("/v1")
	pub := v1.Group("/user")
	pub.POST("", a.ManageApi.CreateUser)
	// v1 end
}
