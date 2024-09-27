package handlers

import (
	"2task/internal/services"
)

type Handler struct {
	MessageService *services.MessageService
	UserService    *services.UserService
}

func NewHandler(messageService *services.MessageService, userService *services.UserService) *Handler {
	return &Handler{
		messageService,
		userService,
	}
}
