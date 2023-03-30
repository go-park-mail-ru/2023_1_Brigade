package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/messages"
	"project/internal/model"
)

type messageHandler struct {
	messageUsecase messages.Usecase
	upgrader       websocket.Upgrader
	clients        map[uint64]*websocket.Conn
	//clients        map[*websocket.Conn]bool
	//broadcast      chan []byte
	tmp_counter uint64
}

func (u *messageHandler) SendMessagesHandler(ctx echo.Context) error {
	ws, err := u.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}

	//session := ctx.Get("session").(model.Session)
	//u.clients[session.UserId] = ws

	// заглушка
	u.clients[u.tmp_counter+1] = ws
	u.tmp_counter++

	defer func() {
		err := ws.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	//{ "body": "string", "author_id": 1, "chat_id": 0 }
	// {"message":{"id":0,"body":"hello world!","author_id":0,"chat_id":0,"is_read":false},"receiver_id":0}

	for {
		_, message, err := ws.ReadMessage() // блокирующая
		if err != nil {
			log.Error(err)
		}

		err = u.messageUsecase.SendMessage(ctx, message)
		if err != nil {
			log.Error(err)
		}

		go func() {
			for {
				msg, err := u.messageUsecase.ReceiveMessage(ctx)
				if err != nil {
					log.Error(err)
					return
				}

				var producerMessage model.ProducerMessage
				err = json.Unmarshal(msg, &producerMessage)
				if err != nil {
					log.Error(err)
					return
				}

				client := u.clients[producerMessage.ReceiverID]
				if client == nil {
					log.Error("nil client")
					return
				}

				err = client.WriteMessage(websocket.BinaryMessage, msg)
				if err != nil {
					log.Error(err)
					err = client.Close()
					if err != nil {
						log.Error(err)
					}
					delete(u.clients, producerMessage.ReceiverID)
					return
				}
			}
		}()
	}
}

func NewMessagesHandler(e *echo.Echo, messageUsecase messages.Usecase) messageHandler {
	handler := messageHandler{
		messageUsecase: messageUsecase,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients: make(map[uint64]*websocket.Conn),
	}

	sendMessagesUrl := "/message/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendMessagesHandler)

	return handler
}
