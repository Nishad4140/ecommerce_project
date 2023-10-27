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

//-------------------------- View-Profile --------------------------//

func (cr *UserHandler) ViewProfile(c *gin.Context) {
	userID, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	UserData, err := cr.userUseCase.ViewProfile(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find userprofile",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Profile",
		Data:       UserData,
		Errors:     nil,
	})
}

//-------------------------- Edit-Profile --------------------------//

func (cr *UserHandler) EditProfile(c *gin.Context) {
	userID, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find UserId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var updatingDetails helper.UserReq
	err = c.Bind(&updatingDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind details",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	updatedProfile, err := cr.userUseCase.EditProfile(userID, updatingDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find userprofile",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Profile updated",
		Data:       updatedProfile,
		Errors:     nil,
	})
}

//-------------------------- Update-Password --------------------------//

func (cr *UserHandler) UpdatePassword(c *gin.Context) {
	userID, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find UserId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var Passwords helper.UpdatePassword
	err = c.Bind(&Passwords)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.UpdatePassword(userID, Passwords)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update password",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Password updated",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Add-Address --------------------------//

func (cr *UserHandler) AddAddress(c *gin.Context) {
	userID, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find UserId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var address helper.Address
	err = c.Bind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.AddAddress(userID, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "address added",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Update-Address --------------------------//

func (cr *UserHandler) UpdateAddress(c *gin.Context) {
	paramsId := c.Param("addressId")
	addressID, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AddressId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find UserId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var address helper.Address
	err = c.Bind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.UpdateAddress(userID, addressID, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "address updated",
		Data:       nil,
		Errors:     nil,
	})
}
