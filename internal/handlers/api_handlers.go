package handlers

import (
	"2task/internal/messagesService"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Service *messagesService.MessageService
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

func (h *Handler) PatchMessageHandler(w http.ResponseWriter, r *http.Request) {
	req := new(struct {
		Id      int    `json:"id"`
		Message string `json:"message"`
	})

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Service.UpdateMessageByID(req.Id, messagesService.Message{Text: req.Message})

	if err != nil {
		log.Fatal("[DB] Ошибка обновления записи ", err)
		json.NewEncoder(w).Encode(struct{ err string }{"Ошибка обновления записи"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Updated bool
		Text    string
	}{true, res.Text})
}

func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id int `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteMessageByID(req.Id)

	if err != nil {
		log.Fatal("[DB] Ошибка обновления записи ", err)
		json.NewEncoder(w).Encode(struct{ err string }{"Ошибка обновления записи"})
		return
	}

	json.NewEncoder(w).Encode(struct{ Deleted bool }{true})
}
