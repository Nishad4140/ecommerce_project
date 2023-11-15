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
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	supadminHandler *handler.SupAdminHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler) *ServerHTTP {

	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := engine.Group("/user")
	{
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)
		user.POST("/forgotpass", userHandler.ForgotPassword)

		products := user.Group("/products")
		{
			products.GET("/listallmodels", productHandler.ListAllModel)
			products.GET("/listmodel/:id", productHandler.ListModel)

			products.GET("/listallbrands", productHandler.ListAllProduct)
			products.GET("/listbrand/:id", productHandler.ListProduct)

			products.GET("/listallcategories", productHandler.ListAllCategories)
			products.GET("/listcategory/:id", productHandler.ListCategory)
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

			wallet := user.Group("/wallet")
			{
				wallet.PATCH("/verify", userHandler.VerifyWallet)
			}

			cart := user.Group("/cart")
			{
				cart.POST("/add/:model_id", cartHandler.AddToCart)
				cart.PATCH("/remove/:model_id", cartHandler.RemoveFromCart)
				cart.GET("/list", cartHandler.ListCart)
			}

			order := user.Group("/order")
			{
				order.POST("/orderall/:payment_id", orderHandler.OrderAll)
				order.PATCH("/cancel/:orderId", orderHandler.UserCancelOrder)
				order.GET("/view/:orderId", orderHandler.ListOrder)
				order.GET("/listall", orderHandler.ListAllOrders)
				order.PATCH("/return/:orderId", orderHandler.ReturnOrder)
			}

			//Payment
			user.GET("/order/razorpay/:orderId", paymentHandler.CreateRazorpayPayment)
			user.GET("/payment-handler", paymentHandler.PaymentSuccess)

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

			category := admin.Group("/category")
			{
				category.POST("/create", productHandler.CreateCategory)
				category.PATCH("/update/:id", productHandler.UpdatCategory)
				category.DELETE("/delete/:category_id", productHandler.DeleteCategory)
				category.GET("/listall", productHandler.ListAllCategories)
				category.GET("/list/:id", productHandler.ListCategory)
			}

			brand := admin.Group("/brand")
			{
				brand.POST("/create", productHandler.AddProduct)
				brand.PATCH("/update/:id", productHandler.UpdateProduct)
				brand.DELETE("/delete/:id", productHandler.DeleteProduct)
				brand.GET("/listall", productHandler.ListAllProduct)
				brand.GET("/list/:id", productHandler.ListProduct)
			}

			model := admin.Group("/model")
			{
				model.POST("/add", productHandler.AddModel)
				model.PATCH("/update/:id", productHandler.UpdateModel)
				model.DELETE("/delete/:id", productHandler.DeleteModel)
				model.GET("/listall", productHandler.ListAllModel)
				model.GET("/list/:id", productHandler.ListModel)
				model.POST("/uploadimage/:id", productHandler.UploadImage)
			}

			dashboard := admin.Group("/dashboard")
			{
				dashboard.GET("/get", adminHandler.AdminDashBoard)
			}

			order := admin.Group("/order")
			{
				order.PATCH("/update", orderHandler.UpdateOrder)
			}

			//Sales report
			sales := admin.Group("/sales")
			{
				sales.GET("/get", adminHandler.ViewSalesReport)
				sales.GET("/download", adminHandler.DownloadSalesReport)
			}

		}

	}

	supadmin := engine.Group("/supadmin")
	{
		supadmin.POST("/login", supadminHandler.SupAdminLogin)

		supadmin.Use(middleware.SupAdminAuth)
		{
			supadmin.POST("/logout", supadminHandler.SupAdminLogout)

			supAdminUsers := supadmin.Group("/user")
			{
				supAdminUsers.PATCH("/block", supadminHandler.BlockUser)
				supAdminUsers.PATCH("/unblock/:user_id", supadminHandler.UnblockUser)
			}

		}
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {

	sh.engine.LoadHTMLGlob("../../template/*.html")

	sh.engine.Run(":3000")
}
