package ws

import (
	"context"
	"net/http"
	"project/internal/config"
	"project/internal/microservices/chat"
	"project/internal/microservices/user"
	"project/internal/model"
	"project/internal/monolithic_services/centrifugo"
	httpUtils "project/internal/pkg/http_utils"
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
	centrifugo  centrifugo.Centrifugo
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

		producerMessage = httpUtils.SanitizeStruct(producerMessage).(model.ProducerMessage)

		if producerMessage.Action != config.Create {
			log.Error("action don't create")
			return
		}

		client := u.clients[producerMessage.ReceiverID]
		if client == nil {
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
			AuthorID:       producerMessage.AuthorId,
			ChatName:       chat.Title,
			ChatAvatar:     chat.Avatar,
			AuthorNickname: userAuthor.Nickname,
			Body:           producerMessage.Body,
		}

		if chat.Type == config.Chat {
			notification.ChatName = notification.AuthorNickname

			if len(chat.Members) > 0 {
				if notification.AuthorNickname == chat.Members[0].Nickname {
					notification.ChatAvatar = chat.Members[0].Avatar
				} else {
					notification.ChatAvatar = chat.Members[1].Avatar
				}
			}
		}

		if len(producerMessage.Attachments) != 0 {
			notification.Body = "Картинка"
		}

		if producerMessage.Type == config.Sticker {
			notification.Body = "Стикер"
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
	defer func() {
		err = ws.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	u.clients[session.UserId] = ws

	for {
		_, _, err = ws.ReadMessage()
		if err != nil {
			return err
		}
	}
}

func NewNotificationsHandler(e *echo.Echo, chatUsecase chat.Usecase, userUsecase user.Usecase, centrifugo centrifugo.Centrifugo, channelName string) (notificationsHandler, error) {
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
		centrifugo:  centrifugo,
		channelName: channelName,
	}

	sendMessagesUrl := "/notification/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendNotificationsHandler)

	return handler, nil
}
