package main

import (
	"context"
	"encoding/json"
	"fmt"
	"set-env-redis/schemas"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx = context.Background()

func Redis() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
	}
	return client
}

func configMysqlBasicInfo() (schemas.ConfigLPR, error) {
	val, err := Redis().HGetAll(ctx, "mysql-basic-info").Result()
	if err != nil {
		fmt.Println("Falha ao carregar o mysql-basic-info", err)
	}
	jsonStr, err := json.Marshal(val)
	if err != nil {
		fmt.Println("Falh ao Marshal JSON", err)
	}

	config := schemas.ConfigLPR{}
	json.Unmarshal([]byte(jsonStr), &config)

	return config, nil
}

func main() {
	configMysqlBasicInfo()
	fmt.Println("Teste de configuracao dentro do Redis")
}
