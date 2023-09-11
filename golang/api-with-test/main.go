package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type applicaton struct {
	Router *gin.Engine
}

func newApp() *applicaton {
	router := gin.Default()
	router.GET("/user/:id", getUser())
	return &applicaton{Router: router}
}

func (a *applicaton) Start() {
	log.Fatal(http.ListenAndServe("localhost:8888", a.Router))
}

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("TomsFavouriteDring", "Beer", 0, "/", "here.com", false, false)

		id := c.Param("id")
		if id == "1234" {
			user := &User{ID: id, Name: "Anderson"}
			c.JSON(http.StatusOK, user)
			return
		}

		c.AbortWithStatus(http.StatusNotFound)
	}
}

func main() {
	newApp().Start()
}
