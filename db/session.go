package db

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var Store *session.Store

func InitSession() {
	storage := redis.New()
	Store = session.New(session.Config{
		Expiration: time.Hour * 24 * 15,
		Storage:    storage,
	})
}
