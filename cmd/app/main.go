package main

import (
	"2task/internal/database"
	"2task/internal/handlers"
	"2task/internal/messagesService"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/messages", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/messages", handler.PatchMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/messages", handler.DeleteMessageHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
