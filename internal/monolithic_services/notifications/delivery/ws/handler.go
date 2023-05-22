package ws

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/microservices/chat"
	"project/internal/microservices/user"
	"project/internal/model"
	"time"

	"github.com/centrifugal/centrifuge-go"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	log "github.com/sirupsen/logrus"
)

type notificationsHandler struct {
	chatUsecase chat.Usecase
	userUsecase user.Usecase
	upgrader    websocket.Upgrader
	clients     map[uint64]*websocket.Conn
	centrifugo  *centrifuge.Client
	channelName string
}

func (u *notificationsHandler) SendNotificationsHandler(ctx echo.Context) error {
	sub, _ := u.centrifugo.GetSubscription(u.channelName)
	session := ctx.Get("session").(model.Session)

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		var producerMessage model.ProducerMessage
		err := easyjson.Unmarshal(e.Data, &producerMessage)
		if err != nil {
			log.Error(err)
			return
		}

		if producerMessage.Action != config.Create {
			log.Error("action don't create")
			return
		}

		//if session.UserId == producerMessage.AuthorId {
		//	return
		//}

		client := u.clients[producerMessage.ReceiverID]
		if client == nil {
			log.Error("nil client")
			return
		}

		chat, err := u.chatUsecase.GetChatById(context.TODO(), producerMessage.ChatID, producerMessage.AuthorId)
		if err != nil {
			log.Error(err)
			return
		}

		userAuthor, err := u.userUsecase.GetUserById(context.TODO(), producerMessage.AuthorId)
		if err != nil {
			log.Error(err)
			return
		}

		notification := model.Notification{
			ChatName:       chat.Title,
			ChatAvatar:     chat.Avatar,
			AuthorNickname: userAuthor.Nickname,
			Body:           producerMessage.Body,
		}

		//if chat.Type == config.Chat {
		//	userReceiver, err := u.userUsecase.GetUserById(context.TODO(), producerMessage.ReceiverID)
		//	if err != nil {
		//		log.Error(err)
		//		return
		//	}
		//
		//	if session.UserId == userReceiver.Id {
		//		notification.ChatName = userAuthor.Nickname
		//		notification.ChatAvatar = userAuthor.Avatar
		//	} else {
		//		notification.ChatName = userReceiver.Nickname
		//		notification.ChatAvatar = userReceiver.Avatar
		//	}
		//}

		//if chat.Type == config.Chat {
		//	if session.UserId == producerMessage.ReceiverID {
		//		if len(chat.Members) > 0 {
		//			if
		//		}
		//	}
		//	//if len(chat.Members) > 0 {
		//	//	if chat.Members[0].Id == producerMessage.AuthorId {
		//	//		notification.ChatName = chat.Members[0].Nickname
		//	//		notification.AuthorNickname = chat.Members[0].Nickname
		//	//	} else {
		//	//		notification.ChatName = chat.Members[1].Nickname
		//	//		notification.AuthorNickname = chat.Members[1].Nickname
		//	//	}
		//	//	//if notification.ChatName == user.Nickname {
		//	//	//
		//	//	//}
		//	//}
		//}

		//if chat.Type == config.Chat {
		//
		//	//if len(chat.Members) > 0 {
		//	//	if chat.M
		//		//if chat.Members[0].Id == session.UserId {
		//		//	notification.ChatName = chat.Members[0].Nickname
		//		//	notification.ChatAvatar = chat.Members[0].Avatar
		//		//} else {
		//		//	notification.ChatName = chat.Members[1].Nickname
		//		//	notification.ChatAvatar = chat.Members[1].Avatar
		//		//}
		//	//}
		//
		//	//if user.Id == producerMessage.AuthorId {
		//	//
		//	//}
		//
		//}

		if producerMessage.Type == config.Sticker {
			notification.Body = "sticker"
		}

		if producerMessage.ImageUrl != "" {
			notification.Body = "image"
		}

		data, err := easyjson.Marshal(notification)
		if err != nil {
			log.Error(err)
			return
		}

		err = client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Error(err)
		}
	})

	ws, err := u.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	u.clients[session.UserId] = ws

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			return err
		}
	}
	//return nil
	//for {
	//	_, _, err := ws.ReadMessage()
	//	if err != nil {
	//		return err
	//	}
	//}
}

func NewNotificationsHandler(e *echo.Echo, chatUsecase chat.Usecase, userUsecase user.Usecase, centrifugo config.Centrifugo) (notificationsHandler, error) {
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
		return notificationsHandler{}, err
	}

	sub, err := c.NewSubscription(centrifugo.ChannelName, centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		return notificationsHandler{}, err
	}

	err = sub.Subscribe()
	if err != nil {
		return notificationsHandler{}, err
	}

	handler := notificationsHandler{
		chatUsecase: chatUsecase,
		userUsecase: userUsecase,
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

	sendMessagesUrl := "/notification/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendNotificationsHandler)

	return handler, nil
}
