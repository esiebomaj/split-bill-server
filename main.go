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
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	router.Get("/health/", appBackend.HealthHandler)

	router.Get("/groups/", appBackend.FetchGroupsHandler)
	router.Post("/groups/", appBackend.CreateGroupHandler)

	router.Get("/groups/{groupID}/items/", appBackend.AuthMiddleWare(appBackend.FetchItemsByGroupHandler))
	router.Post("/groups/{groupID}/items/", appBackend.AuthMiddleWare(appBackend.CreateItemHandler))
	router.Patch("/groups/{groupID}/items/{itemID}/", appBackend.AuthMiddleWare(appBackend.UpdateItemHandler))

	router.Get("/groups/{groupID}/", appBackend.AuthMiddleWare(appBackend.GetGroupDetailsHandler))
	router.Post("/groups/{groupID}/members/", appBackend.AuthMiddleWare(appBackend.AddUserToGroup))
	router.Get("/groups/{groupID}/allusers/", appBackend.AuthMiddleWare(appBackend.GetAllUsers))

	port := appBackend.Envs.PORT
	fmt.Println("Listening on port ", port)
	server := http.Server{Handler: router, Addr: fmt.Sprintf(":%v", port)}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
