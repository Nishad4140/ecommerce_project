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
	adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler) *ServerHTTP {

	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := engine.Group("/user")
	{
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)

		products := user.Group("/products")
		{
			// products.GET("listallproductItems", productHandler.DisaplyaAllProductItems)
			// products.GET("disaplyproductItem/:id", productHandler.DisaplyProductItem)

			products.GET("/listallproduct", productHandler.ListAllProduct)
			products.GET("/showproduct/:id", productHandler.ListProduct)

			products.GET("/listallcategories", productHandler.ListAllCategories)
			products.GET("/showcategories/:id", productHandler.ListCategory)
		}

		user.Use(middleware.UserAuth)
		{
			user.POST("/logout", userHandler.UserLogout)

			profile := user.Group("/profile")
			{
				profile.GET("/view", userHandler.ViewProfile)
				profile.PATCH("/edit", userHandler.EditProfile)
				profile.PATCH("/updatepassword", userHandler.UpdatePassword)
			}

			address := user.Group("/address")
			{
				address.POST("/add", userHandler.AddAddress)
				address.PATCH("/update/:addressId", userHandler.UpdateAddress)
			}
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
				adminUsers.GET("/list/:user_id", adminHandler.ShowUser)
				adminUsers.GET("/listall", adminHandler.ShowAllUsers)
			}

			// adminSellers := admin.Group("/seller")
			// {
			// 	adminSellers.POST("/create", adminHandler.CreateSeller)
			// }

			category := admin.Group("/category")
			{
				category.POST("/create", productHandler.CreateCategory)
				category.PATCH("/update/:id", productHandler.UpdatCategory)
				category.DELETE("/delete/:category_id", productHandler.DeleteCategory)
				category.GET("/listall", productHandler.ListAllCategories)
				category.GET("/list/:id", productHandler.ListCategory)
			}

			product := admin.Group("/product")
			{
				product.POST("/create", productHandler.AddProduct)
				product.PATCH("/update/:id", productHandler.UpdateProduct)
				product.DELETE("/delete/:id", productHandler.DeleteProduct)
				product.GET("/listall", productHandler.ListAllProduct)
				product.GET("/list/:id", productHandler.ListProduct)
			}

			model := admin.Group("/product-item")
			{
				model.POST("add", productHandler.AddModel)
				// model.PATCH("update/:id", productHandler.UpdateProductItem)
				// model.DELETE("delete/:id", productHandler.DeleteProductItem)
				// model.GET("listall", productHandler.DisaplyaAllProductItems)
				// model.GET("show/:id", productHandler.DisaplyProductItem)
				// model.POST("uploadimage/:id", productHandler.UploadImage)
			}
		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {

	sh.engine.Run(":3000")
}
