package server

import (
	"fmt"
	"net/http"
	"time"

	"example.com/cli-chat/chat"
)

type ServerWrapper struct {
	// Other global variables (Logs)
	Server *http.Server
}

func NewChatServer() (*ServerWrapper, error) {
	fmt.Printf("Starting server...\n")

	muxWrapper := NewMux()

	ChatRouter := chat.NewChatRouter(muxWrapper.Mux)
	err := ChatRouter.InitializeRoutes()

	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      muxWrapper.Mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return &ServerWrapper{
		Server: server,
	}, nil
}
