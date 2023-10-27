package handler

import (
	"net/http"
	"strconv"

	handlerutil "github.com/Nishad4140/ecommerce_project/pkg/api/handlerUtil"
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

//-------------------------- Login --------------------------//

func (cr *AdminHandler) AdminLogin(c *gin.Context) {
	var admin helper.LoginReq
	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	token, err := cr.adminUseCase.AdminLogin(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("adminToken", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "login succesfully",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Log-Out --------------------------//

func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("adminToken", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "admin logouted",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Report-User --------------------------//

func (cr *AdminHandler) ReportUser(c *gin.Context) {
	var reason helper.ReportData
	err := c.BindJSON(&reason)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminID, err := handlerutil.GetAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find the admin id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.ReportUser(reason, adminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't report the user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user reported",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Show-Single-User --------------------------//

func (cr *AdminHandler) ShowUser(c *gin.Context) {
	paramID := c.Param("user_id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	user, err := cr.adminUseCase.ShowUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user details",
		Data:       user,
		Errors:     nil,
	})

}

//-------------------------- Show-All-Users --------------------------//

func (cr *AdminHandler) ShowAllUsers(c *gin.Context) {
	users, err := cr.adminUseCase.ShowAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "users are",
		Data:       users,
		Errors:     nil,
	})

}

// //-------------------------- Create-Seller --------------------------//

// func (cr *AdminHandler) CreateSeller(c *gin.Context) {
// 	var sellerData helper.CreateSeller
// 	err := c.Bind(&sellerData)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "bind faild",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// 	fmt.Println(sellerData)
// 	createrId, err := handlerutil.GetAdminIdFromContext(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "Can't find AdminId",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}

// 	sellerDetails, err := cr.adminUseCase.CreateSeller(sellerData, createrId)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "Can't Create Seller",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response.Response{
// 		StatusCode: 201,
// 		Message:    "Seller created",
// 		Data:       sellerDetails,
// 		Errors:     nil,
// 	})
// }
