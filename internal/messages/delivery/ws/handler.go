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

func (u messageHandler) SendMessagesHandler(ctx echo.Context) error {
	log.Warn("aaaaaaaaaaaaaaaaa")
	ws, err := u.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Error(err)
		return err
	}

	session := ctx.Get("session").(model.Session)
	u.clients[session.UserId] = ws

	defer func() {
		err := ws.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	//'{ "body": "string", "author_id": 1, "chat_id": 1 }'

	go func() {
		for {
			msg, err := u.messageUsecase.ReceiveMessage(ctx)
			if err != nil {
				log.Error(err)
				//return
			}
			log.Warn("ПРИНЯЛ СООБЩЕНИЕ")

			var producerMessage model.ProducerMessage
			err = json.Unmarshal(msg, &producerMessage)
			log.Warn(producerMessage)
			if err != nil {
				log.Error(err)
				//return
			}

			client := u.clients[producerMessage.ReceiverID]
			//client := u.clients[producerMessage.AuthorId]
			log.Warn(producerMessage.ReceiverID)
			if client == nil {
				log.Error("nil client")
				continue
				//return
			}

			err = client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Error(err)
				continue
				//log.Error(err)
				//err = client.Close()
				//if err != nil {
				//	log.Error(err)
				//}
				//delete(u.clients, producerMessage.ReceiverID)
				//return
			}
			log.Warn("ОТПРАВИЛ КЛИЕНТАМ")
		}
	}()

	//go func() {
	for {
		_, message, err := ws.ReadMessage() // блокирующая
		if err != nil {
			log.Error(err)
			continue
		}

		err = u.messageUsecase.SendMessage(ctx, message)
		if err != nil {
			log.Error(err)
		}
	}
	//}()
	//}
	//{ "body": "string", "author_id": 1, "chat_id": 1 }
	//return nil
}

func NewMessagesHandler(e *echo.Echo, messageUsecase messages.Usecase) messageHandler {
	log.Warn("messages зашел")
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

	log.Warn("okkkk navesil")

	return handler
}
