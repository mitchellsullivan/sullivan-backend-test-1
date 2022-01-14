package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
	"github.com/mitchellsullivan/sullivan-backend-test-1/users"
	"gorm.io/gorm"
	"net/http"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func closeDb(db *gorm.DB) {
	sqlDb, err := db.DB()
	if err != nil {
		err := sqlDb.Close()

		if err != nil {
			return
		}
	}
}

func main() {
	common.InitCache()

	db := common.InitDb()
	Migrate(db)
	defer closeDb(db)

	r := gin.Default()

	apiV1Group := r.Group("api/v1")

	apiV1Group.HEAD("ping", Ping)
	apiV1Group.GET("ping", Ping)

	apiV1Group.POST("auth/register", users.AuthRegister)
	apiV1Group.POST("auth/login", users.AuthLogin)

	apiV1Group.Use(users.AuthMiddleware())

	apiV1Group.POST("auth/logout", users.AuthLogout)
	apiV1Group.GET("auth/refresh", users.AuthRefresh)

	apiV1Group.GET("user", users.UserRetrieveInfo)

	err := r.Run("0.0.0.0:8099")

	if err != nil {
		return
	}
}

func Ping(ctx *gin.Context) {
	status := http.StatusOK

	if ctx.Request.Method == http.MethodHead {
		ctx.AbortWithStatus(status)
	} else {
		ctx.JSON(status, gin.H{
			"message": "pong",
		})
	}
}
