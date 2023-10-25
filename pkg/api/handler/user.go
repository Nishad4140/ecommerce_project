package handler

import (
	"net/http"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

//-------------------------- Sign-Up --------------------------//

func (cr *UserHandler) UserSignUp(c *gin.Context) {
	var user helper.UserReq
	err := c.BindJSON(&user)
	// fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userData, err := cr.userUseCase.UserSignUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable signup ",
			Data:       response.UserData{},
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       userData,
		Errors:     nil,
	})
}

//-------------------------- Login --------------------------//

func (cr *UserHandler) UserLogin(c *gin.Context) {
	var user helper.LoginReq
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	token, err := cr.userUseCase.UserLogin(user)
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
	c.SetCookie("userToken", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "login succesfully",
		Data:       nil,
		Errors:     nil,
	})

}

//-------------------------- Log-Out --------------------------//

func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("userToken", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user logouted",
		Data:       nil,
		Errors:     nil,
	})
}
