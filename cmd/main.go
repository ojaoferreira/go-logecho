package main

import (
	"go-logecho/config"
	"go-logecho/handler"
	"go-logecho/repository"
	"go-logecho/router"
	"go-logecho/service"
	"net/http"
	"time"
)

func main() {

	config.InitZAP()

	go keepAlive()

	// Dependency injection
	repo := repository.NewMessageRepository(config.Logger)
	srv := service.NewMessageService(repo)
	hdl := handler.NewMessageHandler(srv)

	r := router.NewRouter(hdl)

	config.Logger.Info("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func keepAlive() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		config.Logger.Info("I am alive!")
	}
}
