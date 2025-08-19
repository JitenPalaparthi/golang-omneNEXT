package models_test

import (
	"testing"
	"time"
	"users-service/models"

	"github.com/stretchr/testify/assert"
)

func TestUserValidateSuccess(t *testing.T) {
	user := models.User{Name: "Jiten", Email: "JItenP@Outlook.com", Mobile: "9618558500", CommonModel: models.CommonModel{ID: 101, Status: "active", LastModified: time.Now().Unix()}}
	err := (&user).Validate()
	if err != nil {
		t.Fatalf("The test failes as it should not return any error")
	}
}

func TestUserValidateFailure(t *testing.T) {
	user := models.User{Email: "JitenP@Outlook.com", Mobile: "9618558500", CommonModel: models.CommonModel{ID: 101, Status: "active", LastModified: time.Now().Unix()}}
	err := (&user).Validate()
	expected := models.ErrInvalidName
	if err != expected {
		t.Fatalf("The test failed as it should return error %v", expected)
	}
}

func TestUserValidateAll(t *testing.T) {
	type UserTestCase struct {
		Name     string
		User     *models.User
		Expected error
	}

	usertestCases := []UserTestCase{
		UserTestCase{"test nil error", &models.User{Name: "Jiten", Email: "JItenP@Outlook.com", Mobile: "9618558500", CommonModel: models.CommonModel{ID: 101, Status: "active", LastModified: time.Now().Unix()}}, nil},
		UserTestCase{"test user name not found error", &models.User{Email: "JItenP@Outlook.com", Mobile: "9618558500", CommonModel: models.CommonModel{ID: 101, Status: "active", LastModified: time.Now().Unix()}}, models.ErrInvalidName},
		UserTestCase{"test user email not found error", &models.User{Name: "Jiten", Mobile: "9618558500", CommonModel: models.CommonModel{ID: 101, Status: "active", LastModified: time.Now().Unix()}}, models.ErrInvalidEmail},
		UserTestCase{"test user mobile not found error", &models.User{Name: "Jiten", Email: "JitenP@Outlook.com", CommonModel: models.CommonModel{ID: 101, Status: "active", LastModified: time.Now().Unix()}}, models.ErrInvalidMobile},
		//UserTest{2, "test user not found error", models.User{Email: "JItenP@Outlook.com", Mobile: "9618558500", CommonModel: models.CommonModel{ID: 101, Status: "active", LastModified: time.Now().Unix()}}, models.ErrInvalidName},
	}
	t.Parallel()
	for _, userTest := range usertestCases {
		//userTest := userTest
		t.Run(userTest.Name, func(t *testing.T) {
			err := userTest.User.Validate()
			if err != userTest.Expected {
				t.Fatalf("The test failed--> %v", err)
			}
		})
	}

}

func TestUserToBytes(t *testing.T) {
	user := models.User{Name: "Jiten", Email: "Jitenp@Outlook.Com", Mobile: "9618558500"}
	expected := []byte{123, 34, 105, 100, 34, 58, 48, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 34, 44, 34, 108, 97, 115, 116, 95, 109, 111, 100, 105, 102, 105, 101, 100, 34, 58, 48, 44, 34, 110, 97, 109, 101, 34, 58, 34, 74, 105, 116, 101, 110, 34, 44, 34, 101, 109, 97, 105, 108, 34, 58, 34, 74, 105, 116, 101, 110, 112, 64, 79, 117, 116, 108, 111, 111, 107, 46, 67, 111, 109, 34, 44, 34, 109, 111, 98, 105, 108, 101, 34, 58, 34, 57, 54, 49, 56, 53, 53, 56, 53, 48, 48, 34, 125}
	actual := user.ToBytes()
	//fmt.Println(bytes)
	assert.Equal(t, expected, actual)
	//expected:= []byte{}
}
