package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
)

type UserSerializer struct {
	c *gin.Context
}

// when returning a new token after login or refresh

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (s *UserSerializer) LoginResponse() LoginResponse {
	ctxUser := s.c.MustGet("ctx_user").(User)

	user := LoginResponse{
		Username: ctxUser.Username,
		Email:    ctxUser.Email,
		Token:    common.GenerateJwt(ctxUser.ID),
	}

	return user
}

// when simply retrieving info

type InfoResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (s *UserSerializer) InfoResponse() InfoResponse {
	ctxUser := s.c.MustGet("ctx_user").(User)

	user := InfoResponse{
		Username: ctxUser.Username,
		Email:    ctxUser.Email,
	}

	return user
}
