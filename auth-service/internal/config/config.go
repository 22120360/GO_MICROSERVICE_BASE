package config

import (
	"auth-service/internal/database"
	"github.com/redis/go-redis/v9"
)

type ApiConfig struct {
	DB          *database.Queries
	RedisClient *redis.Client
}