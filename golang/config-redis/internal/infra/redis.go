package infra

import (
	"config-redis/internal/config"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Redis() *redis.Client {
	env := config.SetEnv()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.Host, env.Port),
		Password: "",
		DB:       0,
	})

	return client
}
