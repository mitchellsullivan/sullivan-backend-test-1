package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	DB    *gorm.DB
	CACHE *redis.Client
	Ctx   = context.TODO()
)

func InitDb() *gorm.DB {
	dsn, connStrIsSet := os.LookupEnv("DB_CONN_STR")

	if !connStrIsSet {
		dsn = "host=localhost user=sullivan_backend_test password=yourStrongPassword123 port=43210 sslmode=disable"
	}

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("db failed to open: ", err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)

	DB = db
	return DB
}

func GetDb() *gorm.DB {
	return DB
}

func InitCache() *redis.Client {
	var redisAddr string

	redisAddr, ok := os.LookupEnv("REDIS_URL")

	if !ok {
		redisAddr = "localhost:6379"
	}

	var redisPass string

	redisPass, ok = os.LookupEnv("REDIS_PASSWORD")

	if !ok {
		redisPass = "yourStrongPassword123"
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		//DB:       0,
	})

	CACHE = cache
	return cache
}

func GetCache() *redis.Client {
	return CACHE
}

func GetCachedToken(token *jwt.Token) *redis.StringCmd {
	return CACHE.Get(token.Signature)
}

func SetBlacklistedToken(token *jwt.Token) *redis.StatusCmd {
	// set to expire in one hour, because JWT expires in one hour
	return CACHE.Set(token.Signature, 1, 1*time.Hour)
}
