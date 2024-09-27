package handlers

import (
	"2task/internal/api"
	"2task/internal/services"
	"context"
	"fmt"
)

func (h *Handler) GetMessages(_ context.Context, _ api.GetMessagesRequestObject) (api.GetMessagesResponseObject, error) {
	allMessages, err := h.MessageService.GetAllMessages()
	if err != nil {
		return nil, err
	}

	response := api.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := api.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request api.PostMessagesRequestObject) (api.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := services.Message{Text: *messageRequest.Message}
	createdMessage, err := h.MessageService.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := api.PostMessages200JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}

	return response, nil
}

func (h *Handler) DeleteMessages(_ context.Context, request api.DeleteMessagesRequestObject) (api.DeleteMessagesResponseObject, error) {
	var res api.DeleteMessages400TextResponse
	messageID := request.Body.Id
	err := h.MessageService.DeleteMessageByID(int(*messageID))

	if err != nil {
		res = "err"
		fmt.Println("[DB] Ошибка обновления записи ", err)
		return res, err
	}

	res = "deleted"
	return res, nil

}

func (h *Handler) PatchMessages(_ context.Context, request api.PatchMessagesRequestObject) (api.PatchMessagesResponseObject, error) {
	req := request.Body
	messageToUpdate := services.Message{Text: *req.Message}
	createdMessage, err := h.MessageService.UpdateMessageByID(int(*req.Id), messageToUpdate)

	if err != nil {
		return nil, err
	}

	response := api.PatchMessages200JSONResponse{
		Id:      req.Id,
		Message: &createdMessage.Text,
	}

	return response, nil
}
