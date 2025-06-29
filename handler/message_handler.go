package handler

import (
	"encoding/json"
	"go-logecho/model"
	"go-logecho/service"
	"net/http"
)

type MessageHandler struct {
	svc service.MessageService
}

func NewMessageHandler(s service.MessageService) MessageHandler {
	return MessageHandler{s}
}

func (m *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message model.Message
	json.NewDecoder(r.Body).Decode(&message)

	m.svc.CreateMessage(message)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}
