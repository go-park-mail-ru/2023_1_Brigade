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
	//config := centrifuge.Config{}
	//c := centrifuge.NewProtobufClient("ws://centrifugo:8900/connection/websocket", config)
	//c.OnError(func(e centrifuge.ErrorEvent) {
	//	log.Error(e.Error)
	//})
	//c.OnConnecting(func(e centrifuge.ConnectingEvent) {
	//	log.Warn("Connecting - %d (%s)", e.Code, e.Reason)
	//})
	//c.OnConnected(func(e centrifuge.ConnectedEvent) {
	//	log.Warn("Connected with ID %s", e.ClientID)
	//})
	//c.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
	//	log.Warn("Disconnected: %d (%s)", e.Code, e.Reason)
	//})
	//
	//c.OnError(func(e centrifuge.ErrorEvent) {
	//	log.Error("Error: %s", e.Error.Error())
	//})
	//
	//c.OnMessage(func(e centrifuge.MessageEvent) {
	//	log.Printf("Message from server: %s", string(e.Data))
	//})
	//
	//c.OnSubscribed(func(e centrifuge.ServerSubscribedEvent) {
	//	log.Warn("Subscribed to server-side channel %s: (was recovering: %v, recovered: %v)", e.Channel, e.WasRecovering, e.Recovered)
	//})
	//c.OnSubscribing(func(e centrifuge.ServerSubscribingEvent) {
	//	log.Warn("Subscribing to server-side channel %s", e.Channel)
	//})
	//c.OnUnsubscribed(func(e centrifuge.ServerUnsubscribedEvent) {
	//	log.Warn("Unsubscribed from server-side channel %s", e.Channel)
	//})
	//
	//c.OnPublication(func(e centrifuge.ServerPublicationEvent) {
	//	log.Warn("Publication from server-side channel %s: %s (offset %d)", e.Channel, e.Data, e.Offset)
	//})
	//c.OnJoin(func(e centrifuge.ServerJoinEvent) {
	//	log.Warn("Join to server-side channel %s: %s (%s)", e.Channel, e.User, e.Client)
	//})
	//c.OnLeave(func(e centrifuge.ServerLeaveEvent) {
	//	log.Warn("Leave from server-side channel %s: %s (%s)", e.Channel, e.User, e.Client)
	//})
	//defer c.Close()
	//
	//err := c.Connect()
	//if err != nil {
	//	log.Error(err)
	//}
	//
	//sub, err := c.NewSubscription("channel", centrifuge.SubscriptionConfig{
	//	Recoverable: true,
	//	JoinLeave:   true,
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//sub.OnSubscribing(func(e centrifuge.SubscribingEvent) {
	//	log.Warn("Subscribing on channel %s - %d (%s)", sub.Channel, e.Code, e.Reason)
	//})
	//sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
	//	log.Warn("Subscribed on channel %s, (was recovering: %v, recovered: %v)", sub.Channel, e.WasRecovering, e.Recovered)
	//})
	//sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
	//	log.Warn("Unsubscribed from channel %s - %d (%s)", sub.Channel, e.Code, e.Reason)
	//})
	//
	//sub.OnError(func(e centrifuge.SubscriptionErrorEvent) {
	//	log.Error("Subscription error %s: %s", sub.Channel, e.Error)
	//})

	//sub.OnPublication(func(e centrifuge.PublicationEvent) {
	go func() {
		log.Warn("получил публикацию")
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
			return
		}

		err = client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Error(err)
		}
	}()
	//})

	//sub.OnJoin(func(e centrifuge.JoinEvent) {
	//	log.Warn("Someone joined %s: user id %s, client id %s", sub.Channel, e.User, e.Client)
	//})
	//sub.OnLeave(func(e centrifuge.LeaveEvent) {
	//	log.Warn("Someone left %s: user id %s, client id %s", sub.Channel, e.User, e.Client)
	//})
	//
	//err = sub.Subscribe()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//res, err := sub.Publish(context.Background(), []byte{})
	//log.Error(res, err)

	ws, err := u.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Error(err)
		return err
	}

	session := ctx.Get("session").(model.Session)
	u.clients[session.UserId] = ws

	defer ws.Close()
	//defer func() {
	//	err := ws.Close()
	//	if err != nil {
	//		log.Error(err)
	//	}
	//}()

	//
	//go func() {
	//	for {
	//		msg, err := u.messageUsecase.ReceiveMessage(ctx)
	//		if err != nil {
	//			log.Error(err)
	//		}
	//
	//		var producerMessage model.ProducerMessage
	//		err = json.Unmarshal(msg, &producerMessage)
	//		if err != nil {
	//			log.Error(err)
	//		}
	//
	//		client := u.clients[producerMessage.ReceiverID]
	//		if client == nil {
	//			continue
	//		}
	//
	//		err = client.WriteMessage(websocket.TextMessage, msg)
	//		if err != nil {
	//			log.Error(err)
	//			continue
	//		}
	//	}
	//}()

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
