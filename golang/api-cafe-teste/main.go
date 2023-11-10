package main

import (
	"api-cafe/config"
	"api-cafe/db"
	"api-cafe/router"
	"api-cafe/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

// var port = os.Getenv("PORT")
var port = os.Getenv("api-coffe-port")

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	config.SetEnv(".env")
	fmt.Println(port)
	var c Config
	c.Port = port

	// dsn := os.Getenv("DSN")
	dsn := os.Getenv("api-coffe-dsn")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: c,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
