package main

import (
	"api-dindin/config"
	"api-dindin/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("COnfig initilization erorr: %v", err)
		return
	}

	router.Initialize()
}
