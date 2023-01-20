package chat

import (
	"fmt"

	"nhooyr.io/websocket"
)

type ChatService struct {
	subscribers []*Subscriber
}

type Subscriber struct {
	connection *websocket.Conn
	id         int
}

func NewChatService() *ChatService {
	fmt.Printf("Starting chat service...\n")

	return &ChatService{
		subscribers: make([]*Subscriber, 0),
	}
}

func (cs *ChatService) addSubscriber(c *websocket.Conn) (*Subscriber, error) {
	subscriber := &Subscriber{
		connection: c,
		id:         len(cs.subscribers),
	}

	cs.subscribers = append(cs.subscribers, subscriber)

	fmt.Printf("New subscriber added\n")

	return subscriber, nil
}
