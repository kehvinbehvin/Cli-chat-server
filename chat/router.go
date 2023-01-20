package chat

import (
	"fmt"
	"net/http"
)

type ChatRouter struct {
	Mux        *http.ServeMux
	controller *ChatController
}

func NewChatRouter(m *http.ServeMux) *ChatRouter {
	fmt.Printf("Starting chat router...\n")

	chatService := NewChatService()
	chatController := NewChatController(chatService)

	return &ChatRouter{
		Mux:        m,
		controller: chatController,
	}
}

func (cr *ChatRouter) InitializeRoutes() error {
	cr.Mux.HandleFunc("/subscribe", cr.controller.Subscribe)
	return nil
}
