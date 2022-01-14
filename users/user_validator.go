package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
)

const DefaultUserPassword = "Default User Password 123456789"

type UserValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`

	UserModel User `json:"-"`
}

func (v *UserValidator) Bind(c *gin.Context) error {
	err := common.CustomShouldBind(c, v)

	if err != nil {
		return err
	}

	v.UserModel.Username = v.User.Username
	v.UserModel.Email = v.User.Email

	if v.User.Password != DefaultUserPassword {
		err := v.UserModel.setPassword(v.User.Password)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewUserValidator() UserValidator {
	userModelValidator := UserValidator{}
	return userModelValidator
}
