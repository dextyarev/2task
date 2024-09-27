package main

import (
	"2task/internal/api"
	"2task/internal/database"
	"2task/internal/handlers"
	"2task/internal/repositories"
	"2task/internal/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&services.Message{}, &services.User{})

	if err != nil {
		log.Fatal("[DB] Error AutoMigrate")
	}

	userRepo := repositories.NewUserRepository(database.DB)
	messageRepo := repositories.NewMessageRepository(database.DB)

	messageService := services.NewMessageService(messageRepo)
	userService := services.NewUserService(userRepo)

	handler := handlers.NewHandler(messageService, userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := api.NewStrictHandler(handler, nil) // тут будет ошибка
	api.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
