package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"notify-server/internal/service/handle"
)

//var _ IRouter = (*ManageRouter)(nil)

// ProviderSetRouter 注入router
var ProviderSetRouter = wire.NewSet(wire.Struct(new(ManageRouter), "*"),
	wire.Struct(new(MonitorRouter), "*"))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	GetServerName() string
}

// ManageRouter 路由管理器
type ManageRouter struct {
	ManageApi *handle.UserService
} // end

// Register 注册路由
func (a *ManageRouter) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// GetServerName 获取路由绑定的server-name
func (a *ManageRouter) GetServerName() string {
	return "manage-server"
}

// RegisterAPI register api group router
func (a *ManageRouter) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	v1 := g.Group("/v1")
	pub := v1.Group("/user")
	pub.POST("", a.ManageApi.CreateUser)
	// v1 end
}

// MonitorRouter 路由管理器
type MonitorRouter struct {
	ManageApi *handle.UserService
} // end

// Register 注册路由
func (a *MonitorRouter) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// GetServerName 获取路由绑定的server-name
func (a *MonitorRouter) GetServerName() string {
	return "manage-server"
}

// RegisterAPI register api group router
func (a *MonitorRouter) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	v1 := g.Group("/v2")
	pub := v1.Group("/user")
	pub.POST("", a.ManageApi.CreateUser)
	// v1 end
}
