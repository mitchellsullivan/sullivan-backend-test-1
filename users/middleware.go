package users

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		SetContextUser(c, "")

		token, err := request.ParseFromRequest(c.Request, common.GetJwtAccessTokenExtractor(), common.KeyFunc)

		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// is token revoked (logged out)
		existing := common.GetCachedToken(token)

		if len(existing.Val()) > 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["sub"].(string)
			SetContextUser(c, userId)
		}
	}
}
