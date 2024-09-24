package main

import (
	"2task/internal/database"
	"2task/internal/handlers"
	"2task/internal/messagesService"
	"2task/internal/web/messages"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := messages.NewStrictHandler(handler, nil) // тут будет ошибка
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
