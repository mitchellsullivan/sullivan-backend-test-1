package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func SetContextUser(ctx *gin.Context, ctxUserId string) {
	var user User

	if ctxUserId != "" {
		db := common.GetDb()
		db.First(&user, "id = ?", ctxUserId)
	}

	ctx.Set("ctx_user", user)
}

func GenerateULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
