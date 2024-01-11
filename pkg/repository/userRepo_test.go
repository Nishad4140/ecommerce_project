package repository

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserSignUp(t *testing.T) {
	tests := []struct {
		name           string
		input          helper.UserReq
		expectedOutput response.UserData
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{
			name: "successful creations",
			input: helper.UserReq{
				Name:     "nishad",
				Email:    "nishadshanid0@gmail.com",
				Mobile:   "8848994999",
				Password: "1234abcd",
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "nishad",
				Email:  "nishadshanid0@gmail.com",
				Mobile: "8848994999",
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"}).
					AddRow(1, "nishad", "nishadshanid0@gmail.com", "8848994999")

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("nishad", "nishadshanid0@gmail.com", "8848994999", "1234abcd").
					WillReturnRows(rows)
			},
			expectedErr: nil,
		},
		{
			name: "duplicate user",
			input: helper.UserReq{
				Name:     "nishad",
				Email:    "nishadshanid0@gmail.com",
				Mobile:   "8848994999",
				Password: "1234abcd",
			},
			expectedOutput: response.UserData{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("nishad", "nishadshanid0@gmail.com", "8848994999", "1234abcd").
					WillReturnError(errors.New("email or phone number alredy used"))
			},
			expectedErr: errors.New("email or phone number alredy used"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			userRepository := NewUserRepository(gormDB)

			tt.buildStub(mock)

			actualOutput, actualErr := userRepository.UserSignUp(tt.input)

			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else {
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}

		})
	}
}
