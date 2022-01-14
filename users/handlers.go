package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/request"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
	"net/http"
)

func AuthRegister(c *gin.Context) {
	validator := NewUserValidator()

	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.BindingValidationError(err))
		return
	}

	validator.UserModel.ID = GenerateULID()

	if err := Save(&validator.UserModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("db", err))
		return
	}

	c.Set("ctx_user", validator.UserModel)
	serializer := UserSerializer{c}

	c.JSON(http.StatusCreated, gin.H{
		"user": serializer.LoginResponse(),
	})
}

func AuthLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()

	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.BindingValidationError(err))
		return
	}

	userModel, err := FindOne(&User{
		Email: loginValidator.UserModel.Email,
	})

	if err != nil {
		c.JSON(
			http.StatusForbidden,
			common.NewError("login", errors.New("not registered")),
		)
		return
	}

	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(
			http.StatusForbidden,
			common.NewError("login", errors.New("invalid password")),
		)
		return
	}

	SetContextUser(c, userModel.ID)
	serializer := UserSerializer{c}

	c.JSON(http.StatusOK, gin.H{
		"user": serializer.LoginResponse(),
	})
}

func AuthRefresh(c *gin.Context) {
	// should not have error because protected by middleware which already checks jwt
	token, _ := request.ParseFromRequest(c.Request, common.GetJwtAccessTokenExtractor(), common.KeyFunc)

	// logout current jwt
	result := common.SetBlacklistedToken(token)
	_ = result

	serializer := UserSerializer{c}

	c.JSON(http.StatusOK, gin.H{
		"user": serializer.LoginResponse(),
	})
}

func AuthLogout(c *gin.Context) {
	// should not have error because protected by middleware which already checks jwt
	token, _ := request.ParseFromRequest(c.Request, common.GetJwtAccessTokenExtractor(), common.KeyFunc)

	result := common.SetBlacklistedToken(token)
	_ = result
	
	c.AbortWithStatus(http.StatusNoContent)
}

func UserRetrieveInfo(c *gin.Context) {
	serializer := UserSerializer{c}

	c.JSON(http.StatusOK, gin.H{
		"user": serializer.InfoResponse(),
	})
}
