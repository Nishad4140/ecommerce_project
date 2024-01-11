package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nishad4140/ecommerce_project/pkg/api/middleware"
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/usecase/mockUsecase"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)



func TestUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
	cartUseCase := mockUsecase.NewMockCartUsecase(ctrl)
	UserHandler := NewUserHandler(userUseCase, cartUseCase)

	testData := []struct {
		name             string
		loginData        helper.LoginReq
		buildStub        func(userUsecase mockUsecase.MockUserUseCase)
		expectedCode     int
		expectedResponse response.Response
		expectedError    error
	}{
		{
			name: "valid login",
			loginData: helper.LoginReq{
				Email:    "nishadshanid40@gmail.com",
				Password: "1234abcd",
			},
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().UserLogin(helper.LoginReq{
					Email:    "nishadshanid40@gmail.com",
					Password: "1234abcd",
				}).Times(1).
					Return("validToken", nil)
			},
			expectedCode: 200,
			expectedResponse: response.Response{
				StatusCode: 200,
				Message:    "login successfully",
				Data:       nil,
				Errors:     nil,
			},
			expectedError: nil,
		},
		{
			name: "invalid login",
			loginData: helper.LoginReq{
				Email:    "invalid@example.com",
				Password: "invalidPassword",
			},
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().UserLogin(helper.LoginReq{
					Email:    "invalid@example.com",
					Password: "invalidPassword",
				}).Times(1).
					Return("", errors.New("invalid credentials"))
			},
			expectedCode: 400,
			expectedResponse: response.Response{
				StatusCode: 400,
				Message:    "failed to login",
				Data:       nil,
				Errors:     "invalid credentials",
			},
			expectedError: errors.New("invalid credentials"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)
			engine := gin.Default()
			recorder := httptest.NewRecorder()
			engine.POST("/user/login", UserHandler.UserLogin)
			var body []byte
			body, err := json.Marshal(tt.loginData)
			assert.NoError(t, err)
			url := "/user/login"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			engine.ServeHTTP(recorder, req)   
			var actual response.Response
			err = json.Unmarshal(recorder.Body.Bytes(), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
		})
	}
}

func TestViewProfile(t *testing.T) {
	ctrl := gomock.NewController(t)

	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
	cartUseCase := mockUsecase.NewMockCartUsecase(ctrl)
	UserHandler := NewUserHandler(userUseCase, cartUseCase)

	testData := []struct {
		name             string
		userID           int64
		buildStub        func(userUsecase mockUsecase.MockUserUseCase)
		expectedCode     int
		expectedResponse response.Response
		expectedData     response.Userprofile
		expectedError    error
	}{
		{
			name:   "valid profile",
			userID: 1,
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().ViewProfile(gomock.Any()).Times(1).
					Return(response.Userprofile{
						Name:   "TestUser",
						Email:  "test@example.com",
						Mobile: "1234567890",
					}, nil)
			},
			expectedCode: 200,
			expectedResponse: response.Response{
				StatusCode: 200,
				Message:    "Profile",
				Data: response.Userprofile{
					Name:   "TestUser",
					Email:  "test@example.com",
					Mobile: "1234567890",
				},
				Errors: nil,
			},
			expectedData: response.Userprofile{
				Name:   "TestUser",
				Email:  "test@example.com",
				Mobile: "1234567890",
			},
			expectedError: nil,
		},
		{
			name:   "invalid profile",
			userID: 2,
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().ViewProfile(gomock.Any()).Times(1).
					Return(response.Userprofile{}, errors.New("user not found"))
			},
			expectedCode: 400,
			expectedResponse: response.Response{
				StatusCode: 400,
				Message:    "Can't find userprofile",
				Data:       nil,
				Errors:     "user not found",
			},
			expectedData:  response.Userprofile{},
			expectedError: errors.New("user not found"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)
			engine := gin.Default()
			recorder := httptest.NewRecorder()
			engine.GET("/user/profile/", middleware.TestUserAuth, UserHandler.ViewProfile)
			// url := fmt.Sprintf("/user/profile?user_id=%d", tt.userID)
			url := "/user/profile/"
			req := httptest.NewRequest(http.MethodGet, url, nil)
			engine.ServeHTTP(recorder, req)
			var actual response.Response
			err := json.Unmarshal(recorder.Body.Bytes(), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
		})
	}
}

// func TestUserSignUp(t *testing.T) {
// 	ctrl := gomock.NewController(t)

// 	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
// 	cartUseCase := mockUsecase.NewMockCartUsecase(ctrl)
// 	UserHandler := NewUserHandler(userUseCase, cartUseCase)

// 	testData := []struct {
// 		name             string
// 		userData         helper.UserReq
// 		buildStub        func(userUsecase mockUsecase.MockUserUseCase)
// 		expectedCode     int
// 		expectedResponse response.Response
// 		expectedData     response.UserData
// 		expectedError    error
// 	}{
// 		{
// 			name: "new user",
// 			userData: helper.UserReq{
// 				Name:     "nishad",
// 				Email:    "nishadshanid40@gmail.com",
// 				Mobile:   "8848994140",
// 				Password: "1234abcd",

// 			},
// 			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
// 				userUsecase.EXPECT().UserSignUp(helper.UserReq{
// 					Name:     "nishad",
// 					Email:    "nishadshanid40@gmail.com",
// 					Mobile:   "8848994140",
// 					Password: "1234abcd",
// 				}).Times(1).
// 					Return(response.UserData{
// 						Id:     1,
// 						Name:   "nishad",
// 						Email:  "nishadshanid40@gmail.com",
// 						Mobile: "8848994140",
// 					}, nil)
// 				cartUseCase.EXPECT().CreateCart(1).Times(1).Return(nil)
// 			},
// 			expectedCode: 201,
// 			expectedResponse: response.Response{
// 				StatusCode: 201,
// 				Message:    "user signed successfully",
// 				Data: response.UserData{
// 					Id:     1,
// 					Name:   "nishad",
// 					Email:  "nishadshanid40@gmail.com",
// 					Mobile: "8848994140",
// 				},
// 				Errors: nil,
// 			},
// 			expectedData: response.UserData{
// 				Id:     1,
// 				Name:   "nishad",
// 				Email:  "nishadshanid40@gmail.com",
// 				Mobile: "8848994140",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "duplicate user",
// 			userData: helper.UserReq{
// 				Name:     "nishad",
// 				Email:    "nishadshanid40@gmail.com",
// 				Mobile:   "8848994140",
// 				Password: "1234abcd",
// 			},
// 			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
// 				userUsecase.EXPECT().UserSignUp(helper.UserReq{
// 					Name:     "nishad",
// 					Email:    "nishadshanid40@gmail.com",
// 					Mobile:   "8848994140",
// 					Password: "1234abcd",
// 				}).Times(1).
// 					Return(
// 						response.UserData{},
// 						errors.New("user already exist"),
// 					)
// 			},
// 			expectedCode: 400,
// 			expectedResponse: response.Response{
// 				StatusCode: 400,
// 				Message:    "unable signup",
// 				Data:       response.UserData{},
// 				Errors:     "user already exits",
// 			},
// 			expectedData:  response.UserData{},
// 			expectedError: errors.New("user already exists"),
// 		},
// 	}

// 	for _, tt := range testData {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.buildStub(*userUseCase)
// 			engine := gin.Default()
// 			recorder := httptest.NewRecorder()
// 			engine.POST("/user/signup", UserHandler.UserSignUp)
// 			var body []byte
// 			body, err := json.Marshal(tt.userData)
// 			assert.NoError(t, err)
// 			url := "/user/signup"
// 			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
// 			engine.ServeHTTP(recorder, req)
// 			var actual response.Response
// 			err = json.Unmarshal(recorder.Body.Bytes(), &actual)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedCode, recorder.Code)
// 			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
// 		})
// 	}
// }

// func TestEditProfile(t *testing.T) {
// 	ctrl := gomock.NewController(t)

// 	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
// 	cartUseCase := mockUsecase.NewMockCartUsecase(ctrl)
// 	UserHandler := NewUserHandler(userUseCase, cartUseCase)

// 	testData := []struct {
// 		name             string
// 		userID           int64
// 		updateDetails    helper.UpdateProfile
// 		buildStub        func(userUsecase mockUsecase.MockUserUseCase)
// 		expectedCode     int
// 		expectedResponse response.Response
// 		expectedData     response.Userprofile
// 		expectedError    error
// 	}{
// 		{
// 			name:   "valid update",
// 			userID: 1,
// 			updateDetails: helper.UpdateProfile{
// 				Name:   "UpdatedUser",
// 				Mobile: "9876543210",
// 			},
// 			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
// 				userUsecase.EXPECT().EditProfile(gomock.Any(), gomock.Any()).Times(1).
// 					Do(func(userID int64, details helper.UpdateProfile) {
// 						assert.Equal(t, int64(1), userID)
// 						assert.Equal(t, "UpdatedUser", details.Name)
// 						assert.Equal(t, "9876543210", details.Mobile)
// 					}).
// 					Return(response.Userprofile{
// 						Name:   "UpdatedUser",
// 						Mobile: "9876543210",
// 					}, nil)
// 			},
// 			expectedCode: 200,
// 			expectedResponse: response.Response{
// 				StatusCode: 200,
// 				Message:    "Profile updated",
// 				Data: response.Userprofile{
// 					Name:   "UpdatedUser",
// 					Mobile: "9876543210",
// 				},
// 				Errors: nil,
// 			},
// 			expectedData: response.Userprofile{
// 				Name:   "UpdatedUser",
// 				Mobile: "9876543210",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name:   "invalid update",
// 			userID: 2,
// 			updateDetails: helper.UpdateProfile{
// 				Name:   "UpdatedUser",
// 				Mobile: "9876543210",
// 			},
// 			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
// 				userUsecase.EXPECT().EditProfile(gomock.Any(), gomock.Any()).Times(1).
// 					Return(response.Userprofile{}, errors.New("user not found"))
// 			},
// 			expectedCode: 400,
// 			expectedResponse: response.Response{
// 				StatusCode: 400,
// 				Message:    "Can't find userprofile",
// 				Data:       nil,
// 				Errors:     "user not found",
// 			},
// 			expectedData:  response.Userprofile{},
// 			expectedError: errors.New("user not found"),
// 		},
// 	}

// 	for _, tt := range testData {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.buildStub(*userUseCase)
// 			engine := gin.Default()
// 			recorder := httptest.NewRecorder()
// 			engine.PUT("/user/profile", middleware.TestUserAuth, UserHandler.EditProfile)
// 			var body []byte
// 			body, err := json.Marshal(tt.updateDetails)
// 			assert.NoError(t, err)
// 			// url := fmt.Sprintf("/user/profile?user_id=%d", tt.userID)
// 			url := "/user/profile"
// 			req := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
// 			engine.ServeHTTP(recorder, req)
// 			var actual response.Response
// 			err = json.Unmarshal(recorder.Body.Bytes(), &actual)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedCode, recorder.Code)
// 			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
// 			if tt.expectedCode == 200 {
// 				// Only check data and errors when the status code is 200
// 				assert.Equal(t, tt.expectedResponse.Data, actual.Data)
// 				assert.Equal(t, tt.expectedResponse.Errors, actual.Errors)
// 			}
// 		})
// 	}
// }
