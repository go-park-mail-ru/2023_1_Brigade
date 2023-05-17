package ws

import (
	"context"
	"encoding/json"
	"github.com/centrifugal/centrifuge-go"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/microservices/messages"
	"project/internal/model"
	"time"
)

type messageHandler struct {
	messageUsecase messages.Usecase
	upgrader       websocket.Upgrader
	clients        map[uint64]*websocket.Conn
	centrifugo     *centrifuge.Client
	channelName    string
}

func (u *messageHandler) SendMessagesHandler(ctx echo.Context) error {
	sub, _ := u.centrifugo.GetSubscription(u.channelName)

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		var producerMessage model.ProducerMessage
		err := json.Unmarshal(e.Data, &producerMessage)
		if err != nil {
			log.Error(err)
			return
		}

		client := u.clients[producerMessage.ReceiverID]
		if client == nil {
			log.Error("nil client")
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
	defer ws.Close()

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

func NewMessagesHandler(e *echo.Echo, messageUsecase messages.Usecase, centrifugo config.Centrifugo) (messageHandler, error) {
	c := centrifuge.NewJsonClient(centrifugo.ConnAddr, centrifuge.Config{})

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		c.Close()
		log.Fatal()
	}()

	err := c.Connect()
	if err != nil {
		return messageHandler{}, err
	}

	sub, err := c.NewSubscription(centrifugo.ChannelName, centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		return messageHandler{}, err
	}

	err = sub.Subscribe()
	if err != nil {
		return messageHandler{}, err
	}

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
		centrifugo:  c,
		channelName: centrifugo.ChannelName,
	}

	sendMessagesUrl := "/message/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendMessagesHandler)

	return handler, nil
}
