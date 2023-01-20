package chat

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

type ChatController struct {
	service ChatService
}

func NewChatController(service *ChatService) *ChatController {
	fmt.Printf("Starting chat controller...\n")

	return &ChatController{
		service: *service,
	}
}

func (cc *ChatController) Subscribe(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		defer c.Close(websocket.StatusInternalError, "[Subscribe]")
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*3600)
	defer cancel()

	currentSubscriber, err := cc.service.addSubscriber(c)
	if err != nil {
		return
	}

	go cc.listen(ctx, currentSubscriber)

	<-ctx.Done()

	c.Close(websocket.StatusNormalClosure, "Server connection closing")
}

func (cc *ChatController) listen(ctx context.Context, currentSubscriber *Subscriber) {
	for {
		c := currentSubscriber.connection
		_, b, err := c.Read(ctx)
		if err != nil {
			return
		}

		fmt.Println(string(b))

		cc.publish(b, currentSubscriber)
	}
}

func (cc *ChatController) publish(m []byte, messageSender *Subscriber) {
	for index, subscriber := range cc.service.subscribers {
		if index != messageSender.id {
			cc.writeTimeout(time.Second*5, subscriber.connection, m)
		}
	}
}

func (cc *ChatController) writeTimeout(timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
