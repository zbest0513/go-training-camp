package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"notify-server/internal/service"
	api "notify/api/server/service"
)

//var _ IRouter = (*ManageRouter)(nil)

// ProviderSetRouter 注入router
//var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))
var ProviderSetRouter = wire.NewSet(wire.Struct(new(ManageRouter), "*"),
	wire.Struct(new(DocRouter), "*"))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	GetServerName() string
}

// ManageRouter 路由管理器
type ManageRouter struct {
	UserApi     api.UserApi
	TagApi      api.TagApi
	TemplateApi api.TemplateApi
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
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g := app.Group("/api")
	v1 := g.Group("/v1")
	{
		user := v1.Group("/user")
		{
			create := user.Group("/create")
			create.POST("", a.UserApi.CreateUser)
			all := user.Group("/all")
			all.GET("", a.UserApi.AllUsers)
			addTags := user.Group("/addTags")
			addTags.POST("", a.UserApi.AddTags)
			updateStatus := user.Group("/updateStatus")
			updateStatus.POST("", a.UserApi.UpdateUserStatus)
		}
	}
	{
		tag := v1.Group("/tag")
		{
			create := tag.Group("/create")
			create.POST("", a.TagApi.CreateTag)
			del := tag.Group("/delete")
			del.POST("", a.TagApi.DeleteTag)
		}
	}
	{
		template := v1.Group("/template")
		{
			create := template.Group("/create")
			create.POST("", a.TemplateApi.CreateTemplate)
		}
	}

	// v1 end
}

// DocRouter 路由管理器
type DocRouter struct {
	UserApi *service.UserService
} // end

// Register 注册路由
func (a *DocRouter) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// GetServerName 获取路由绑定的server-name
func (a *DocRouter) GetServerName() string {
	return "doc-server"
}

// RegisterAPI register api group router
func (a *DocRouter) RegisterAPI(app *gin.Engine) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
