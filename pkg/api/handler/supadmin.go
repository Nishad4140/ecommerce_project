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

type SupAdminHandler struct {
	supadminUseCase services.SupAdminUseCase
}

func NewSupAdminHandler(usecase services.SupAdminUseCase) *SupAdminHandler {
	return &SupAdminHandler{
		supadminUseCase: usecase,
	}
}

//-------------------------- Login --------------------------//

func (cr *SupAdminHandler) SupAdminLogin(c *gin.Context) {
	var supadmin helper.LoginReq
	err := c.BindJSON(&supadmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	token, err := cr.supadminUseCase.SupAdminLogin(supadmin)
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
	c.SetCookie("supadminToken", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "login succesfully",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Logout --------------------------//

func (cr *SupAdminHandler) SupAdminLogout(c *gin.Context) {
	c.SetCookie("supadminToken", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "supadmin logouted",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Logout --------------------------//

func (cr *SupAdminHandler) CreateAdmin(c *gin.Context) {
	var adminData helper.AdminData
	err := c.Bind(&adminData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	admin, err := cr.supadminUseCase.CreateAdmin(adminData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable create admin",
			Data:       response.UserData{},
			Errors:     err.Error(),
		})
		return
	}

	var details = struct {
		Name   string
		Email  string
		Mobile string
	}{
		admin.Name,
		admin.Email,
		admin.Mobile,
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "admin created Successfully",
		Data:       details,
		Errors:     nil,
	})

}

//-------------------------- Block-User --------------------------//

func (cr *SupAdminHandler) BlockUser(c *gin.Context) {
	var body helper.BlockData
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminID, err := handlerutil.GetSupAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find SupAdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.supadminUseCase.BlockUser(body, adminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Block",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "User Blocked",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- UnBlock-User --------------------------//

func (cr *SupAdminHandler) UnblockUser(c *gin.Context) {
	paramsId := c.Param("user_id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.supadminUseCase.UnblockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant unblock user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user unblocked",
		Data:       nil,
		Errors:     nil,
	})
}
