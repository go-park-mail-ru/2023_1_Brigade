package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/messages"
	"project/internal/model"
	"time"
)

type messageHandler struct {
	messageUsecase messages.Usecase
	upgrader       websocket.Upgrader
	clients        map[uint64]*websocket.Conn
}

func (u *messageHandler) SendMessagesHandler(ctx echo.Context) error {
	ws, err := u.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Error(err)
		return err
	}

	session := ctx.Get("session").(model.Session)
	u.clients[session.UserId] = ws
	log.Warn(session.UserId)

	defer func() {
		err := ws.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	go func() {
		for {
			msg, err := u.messageUsecase.ReceiveMessage(ctx)
			if err != nil {
				log.Error(err)
			}

			var producerMessage model.ProducerMessage
			err = json.Unmarshal(msg, &producerMessage)
			if err != nil {
				log.Error(err)
			}

			client := u.clients[producerMessage.ReceiverID]
			if client == nil {
				continue
			}

			err = client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Error(err)
			}
		}
	}()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		err = u.messageUsecase.SendMessage(ctx, message)
		if err != nil {
			log.Error(err)
		}
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
			HandshakeTimeout: time.Second * 3600,
		},
		clients: make(map[uint64]*websocket.Conn),
	}

	sendMessagesUrl := "/message/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendMessagesHandler)

	return handler
}
