package handler

import (
	"errors"
	"net/http"
	"strconv"

	handlerutil "github.com/Nishad4140/ecommerce_project/pkg/api/handlerUtil"
	"github.com/Nishad4140/ecommerce_project/pkg/api/middleware"
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	cartUseCase services.CartUsecase
}

func NewUserHandler(usecase services.UserUseCase, cartUseCase services.CartUsecase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		cartUseCase: cartUseCase,
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

	if user.OTP == "" {
		err = middleware.SendOTP(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "Error sending otp",
				Data:       nil,
				Errors:     err,
			})
			return
		}
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "OTP send successfully, Please enter the otp",
			Data:       nil,
			Errors:     nil,
		})
		return
	} else {
		res := middleware.VerifyOTP(user.Email, user.OTP)
		if !res {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "OTP is incorrect",
				Data:       nil,
				Errors:     errors.New("OTP is incorrect"),
			})
			return
		}
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

	err = cr.cartUseCase.CreateCart(userData.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable create cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.userUseCase.CreateWallet(userData.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to create wallet",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var details = struct {
		Name   string
		Email  string
		Mobile string
	}{
		userData.Name,
		userData.Email,
		userData.Mobile,
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       details,
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
	c.SetCookie("userToken", token, 3600*24*30, "/", "localhost", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "login succesfully",
		Data:       nil,
		Errors:     nil,
	})

}

//-------------------------- Forgot-Password --------------------------//

func (cr *UserHandler) ForgotPassword(c *gin.Context) {
	var newPass helper.ForgotPassword
	err := c.BindJSON(&newPass)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "Unable Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	if newPass.OTP == "" {
		if err := middleware.SendOTP(newPass.Email); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "error in sending the otp",
				Data:       nil,
				Errors:     err,
			})
			return
		}
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "otp send successfully",
			Data:       nil,
			Errors:     nil,
		})
		return
	} else {
		res := middleware.VerifyOTP(newPass.Email, newPass.OTP)
		if !res {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "OTP is incorrect",
				Data:       nil,
				Errors:     errors.New("OTP is incorrect"),
			})
			return
		}
	}
	err = cr.userUseCase.ForgotPassword(newPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to update the password",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "password updated successfully",
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
	var updatingDetails helper.UpdateProfile
	err = c.Bind(&updatingDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
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

//-------------------------- User-Wallet --------------------------//

func (cr *UserHandler) VerifyWallet(c *gin.Context) {
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find UserId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var wallet helper.VerifyWallet
	err = c.BindJSON(&wallet)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "Unable Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	if wallet.OTP == "" {
		if err := middleware.SendOTP(wallet.Email); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "error in sending the otp",
				Data:       nil,
				Errors:     err,
			})
			return
		}
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "otp send successfully",
			Data:       nil,
			Errors:     nil,
		})
		return
	} else {
		res := middleware.VerifyOTP(wallet.Email, wallet.OTP)
		if !res {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "OTP is incorrect",
				Data:       nil,
				Errors:     errors.New("OTP is incorrect"),
			})
			return
		}
	}
	err = cr.userUseCase.VerifyWallet(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to verify",
			Data:       nil,
			Errors:     err,
		})
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "wallet verified successfully",
		Data:       nil,
		Errors:     nil,
	})
}
