package common

import (
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"os"
	"time"
)

func KeyFunc(jwtToken *jwt.Token) (interface{}, error) {
	secretKey, ok := os.LookupEnv("JWT_SECRET_KEY")

	if !ok {
		secretKey = "Default Secret Key 123456789"
	}

	b := []byte(secretKey)
	return b, nil
}

func GetJwtAccessTokenExtractor() *request.MultiExtractor {
	filter := func(token string) (string, error) {
		if len(token) > 6 && token[0:7] == "Bearer " {
			return token[7:], nil
		}
		return token, nil
	}

	extractor := &request.MultiExtractor{
		&request.PostExtractionFilter{
			Extractor: request.HeaderExtractor{
				"Authorization",
			},
			Filter: filter,
		},
		request.ArgumentExtractor{
			"access_token",
		},
	}
	return extractor
}

func GenerateJwt(ulid string) string {
	newJwt := jwt.New(jwt.GetSigningMethod("HS256"))

	newJwt.Claims = jwt.MapClaims{
		"sub": ulid,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	key, _ := KeyFunc(nil)

	token, _ := newJwt.SignedString(key)
	return token
}
