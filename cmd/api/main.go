package main

import (
	"log"
	"net/http"
	"os"

	"github.com/L200160149/be-sewa-alat-berat/app"
	"github.com/L200160149/be-sewa-alat-berat/config"
	"github.com/L200160149/be-sewa-alat-berat/controller"
	"github.com/L200160149/be-sewa-alat-berat/repository"
	"github.com/L200160149/be-sewa-alat-berat/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.InitEnv()

    db := app.NewDB()
    usersRepository := repository.NewUsersRepository()
    validate := validator.New()
    usersService := service.NewUsersService(usersRepository, db, validate)
    usersController := controller.NewUsersController(usersService)
    router := app.NewRouter(usersController)

    appPort := os.Getenv("APP_PORT")
    if appPort == "" {
        appPort = "8080"
    }

    server := http.Server{
        Addr:    ":" + appPort,
        Handler: router,
    }

    log.Printf("Server running on port %s", appPort)
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Server failed: %v", err)
    }
}