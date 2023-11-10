package router

import (
	"api-cafe/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("api/v1/coffes", controllers.GetAllCoffes)
	router.Get("api/v1/coffes/coffe/{id}", controllers.GetCoffeeById)
	router.Post("api/v1/coffes/coffe", controllers.CreateCoffe)
	router.Put("api/v1/coffes/coffe/{id}", controllers.UpdateCoffee)
	router.Delete("api/v1/coffes/coffe/{id}", controllers.DeleteCoffe)

	return router
}
