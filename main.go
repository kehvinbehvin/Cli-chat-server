package main

import (
	"fmt"

	"example.com/cli-chat/server"
)

func main() {
	s, err := server.NewChatServer()
	if err != nil {
		return
	}

	fmt.Printf("Server initialised, listening on port 8080...\n")

	s.Server.ListenAndServe()
}
