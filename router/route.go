package router

import (
	"go-logecho/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h handler.MessageHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/messages", h.CreateMessage).Methods("POST")

	return r
}
