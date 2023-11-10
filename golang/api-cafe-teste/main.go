package main

import (
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

var port = os.Getenv("PORT")

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	var config Config
	config.Port = port

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: config,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
