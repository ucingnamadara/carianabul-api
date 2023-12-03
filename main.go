package main

import (
	"dana/anabul-rest-api/src/database"
	"dana/anabul-rest-api/src/entities"
	"dana/anabul-rest-api/src/handlers"
	"dana/anabul-rest-api/src/repositories"
	"dana/anabul-rest-api/src/services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().PathPrefix(("/api/v1")).Subrouter()
	db := database.NewPostrgresClient()

	db.DB.AutoMigrate(&entities.User{})

	repositories := repositories.NewRepository(db)

	userService := services.NewUserService(repositories.UserRepositoryImpl)
	authService := services.NewAuthService(repositories.UserRepositoryImpl)
	handlers := handlers.NewHttpHandler(userService, authService)

	handlers.Router(router)
	//config all cors
	r := AllowAll().Handler(router)

	fmt.Print("Server Start")
	http.ListenAndServe(":7777", r)
}
