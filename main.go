package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

type AppBackend struct {
	DB   *gorm.DB
	Envs *AppEnvs
}

func main() {
	Envs := ConfigureEnvs()
	db := SetUpDB(Envs)

	appBackend := AppBackend{
		DB:   db,
		Envs: &Envs,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	router.Get("/health/", appBackend.HealthHandler)
	router.Get("/groups/", appBackend.FetchGroupsHandler)
	router.Post("/groups/", appBackend.CreateGroupHandler)
	router.Get("/groups/{groupID}/items/", appBackend.FetchItemsByGroupHandler)
	router.Post("/groups/{groupID}/items/", appBackend.CreateItemHandler)

	port := appBackend.Envs.PORT
	fmt.Println("Listening on port ", port)
	server := http.Server{Handler: router, Addr: fmt.Sprintf(":%v", port)}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
