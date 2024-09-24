package handlers

import (
	"2task/internal/messagesService"
	"2task/internal/web/messages"
	"context"
	"log"
)

type Handler struct {
	Service *messagesService.MessageService
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages200JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}

	return response, nil
}

func (h *Handler) DeleteMessages(_ context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	messageID := request.Body.Id

	err := h.Service.DeleteMessageByID(int(*messageID))

	if err != nil {
		log.Fatal("[DB] Ошибка обновления записи ", err)
		var res messages.DeleteMessages400TextResponse = "err"
		return res, err
	}

	var res messages.DeleteMessages200TextResponse = "deleted"
	return res, nil

}

func (h *Handler) PatchMessages(_ context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	req := request.Body
	messageToUpdate := messagesService.Message{Text: *req.Message}
	createdMessage, err := h.Service.UpdateMessageByID(int(*req.Id), messageToUpdate)

	if err != nil {
		return nil, err
	}

	response := messages.PatchMessages200JSONResponse{
		Id:      req.Id,
		Message: &createdMessage.Text,
	}

	return response, nil
}
