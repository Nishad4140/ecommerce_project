package http

import (
	"github.com/Nishad4140/ecommerce_project/pkg/api/handler"
	"github.com/Nishad4140/ecommerce_project/pkg/api/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler) *ServerHTTP {

	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := engine.Group("/user")
	{
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)

		user.Use(middleware.UserAuth)
		{
			user.POST("/logout", userHandler.UserLogout)
		}

	}

	admin := engine.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)

		admin.Use(middleware.AdminAuth)
		{
			admin.POST("/logout", adminHandler.AdminLogout)

			adminUsers := admin.Group("/user")
			{
				adminUsers.PATCH("/report", adminHandler.ReportUser)
				adminUsers.GET("/list/:user_id",adminHandler.ShowUser)
			}
		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {

	sh.engine.Run(":3000")
}
