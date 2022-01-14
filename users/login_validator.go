package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
)

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`

	UserModel User `json:"-"`
}

func (v *LoginValidator) Bind(c *gin.Context) error {
	err := common.CustomShouldBind(c, v)

	if err != nil {
		return err
	}

	v.UserModel.Email = v.User.Email
	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
