package usecase

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/repository/mockRepo"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      helper.UserReq
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(helper.UserReq)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg helper.UserReq, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestUserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := mockRepo.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo)
	testData := []struct {
		name           string
		input          helper.UserReq
		buildStub      func(userRepo mockRepo.MockUserRepository)
		expectedOutput response.UserData
		expectedError  error
	}{
		{
			name: "new user",
			input: helper.UserReq{
				Name:     "nishad",
				Email:    "nishadshanid40@gmail.com",
				Mobile:   "8848994140",
				Password: "1234abcd",
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(
					EqCreateUserParams(helper.UserReq{
						Name:     "nishad",
						Email:    "nishadshanid40@gmail.com",
						Mobile:   "8848994140",
						Password: "1234abcd",
					},
						"1234abcd")).
					Times(1).
					Return(response.UserData{
						Id:     1,
						Name:   "nishad",
						Email:  "nishadshanid40@gmail.com",
						Mobile: "8848994140",
					}, nil)
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "nishad",
				Email:  "nishadshanid40@gmail.com",
				Mobile: "8848994140",
			},
			expectedError: nil,
		},
		{
			name: "alredy exits",
			input: helper.UserReq{
				Name:     "nishad",
				Email:    "nishadshanid40@gmail.com",
				Mobile:   "8848994140",
				Password: "1234abcd",
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(
					EqCreateUserParams(helper.UserReq{
						Name:     "nishad",
						Email:    "nishadshanid40@gmail.com",
						Mobile:   "8848994140",
						Password: "1234abcd",
					},
						"1234abcd")).
					Times(1).
					Return(response.UserData{},
						errors.New("user alredy exits"))
			},
			expectedOutput: response.UserData{},
			expectedError:  errors.New("user alredy exits"),
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userRepo)
			actualUser, err := userUseCase.UserSignUp(tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, actualUser, tt.expectedOutput)
		})
	}
}
