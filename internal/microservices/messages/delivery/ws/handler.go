package ws

import (
	"context"
	"net/http"
	"project/internal/microservices/messages"
	"project/internal/model"
	"project/internal/monolithic_services/centrifugo"
	"time"

	"github.com/centrifugal/centrifuge-go"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	log "github.com/sirupsen/logrus"
)

type messageHandler struct {
	messageUsecase messages.Usecase
	upgrader       websocket.Upgrader
	clients        map[uint64]*websocket.Conn
	centrifugo     centrifugo.Centrifugo
	channelName    string
}

func (u *messageHandler) SendMessagesHandler(ctx echo.Context) error {
	sub, _ := u.centrifugo.GetSubscription(u.channelName)

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		var producerMessage model.ProducerMessage
		err := easyjson.Unmarshal(e.Data, &producerMessage)
		if err != nil {
			log.Error(err)
			return
		}

		client := u.clients[producerMessage.ReceiverID]
		if client == nil {
			return
		}

		err = client.WriteMessage(websocket.TextMessage, e.Data)
		if err != nil {
			log.Error(err)
		}
	})

	ws, err := u.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer func() {
		err = ws.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	session := ctx.Get("session").(model.Session)
	u.clients[session.UserId] = ws

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		err = u.messageUsecase.PutInProducer(context.TODO(), message)
		if err != nil {
			return err
		}
	}
}

func NewMessagesHandler(e *echo.Echo, messageUsecase messages.Usecase, centrifugo centrifugo.Centrifugo, channelName string) (messageHandler, error) {
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
		clients:     make(map[uint64]*websocket.Conn),
		centrifugo:  centrifugo,
		channelName: channelName,
	}

	sendMessagesUrl := "/message/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendMessagesHandler)

	return handler, nil
}
