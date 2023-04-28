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
	"project/internal/messages"
	"project/internal/model"
	"project/internal/qaas/send_messages/consumer"
	"project/internal/qaas/send_messages/producer"
	"time"
)

type messageHandler struct {
	messageUsecase  messages.Usecase
	consumerUsecase consumer.Usecase
	producerUsecase producer.Usecase
	upgrader        websocket.Upgrader
	clients         map[uint64]*websocket.Conn
	centrifugo      *centrifuge.Client
}

func (u *messageHandler) SendMessagesHandler(ctx echo.Context) error {
	sub, _ := u.centrifugo.GetSubscription("channel")

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		msg, err := u.messageUsecase.PullFromConsumer(context.TODO())
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

		err = client.WriteMessage(websocket.TextMessage, msg)
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
	log.Warn(session)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		err = u.messageUsecase.SwitchMessageType(context.TODO(), message)
		if err != nil {
			log.Error(err)
		}
	}
}

func NewMessagesHandler(e *echo.Echo, messageUsecase messages.Usecase) messageHandler {
	c := centrifuge.NewJsonClient("ws://centrifugo:8900/connection/websocket", centrifuge.Config{})

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		_ = c.Close
		log.Fatal()
	}()

	err := c.Connect()
	if err != nil {
		log.Error(err)
	}

	sub, err := c.NewSubscription("channel", centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Error(err)
	}

	err = sub.Subscribe()
	if err != nil {
		log.Error(err)
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
		clients:    make(map[uint64]*websocket.Conn),
		centrifugo: c,
	}

	sendMessagesUrl := "/message/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendMessagesHandler)

	return handler
}
