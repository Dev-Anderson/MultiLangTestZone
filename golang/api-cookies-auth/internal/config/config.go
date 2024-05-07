package config

import "api-cookies/internal/database"

// import "github.com/redis/go-redis/v9"

type ApiConfig struct {
	DB *database.Queries
	//config redis
	// RedisClient *redis.Client
}
