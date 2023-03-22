package ws

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/messages"
)

type messageHandler struct {
	messageUsecase messages.Usecase
	upgrader       websocket.Upgrader
}

func (u *messageHandler) SendMessagesHandler(ctx echo.Context) error {
	ws, err := u.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	//defer ws.Close()
	//{ "body": "string", "author_id": 12, "chat_id": 131 }
	// {"message":{"id":0,"body":"hello world!","author_id":0,"chat_id":0,"is_read":false},"receiver_id":0}
	go func() {
		defer ws.Close()

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Error(err)
			}

			err = u.messageUsecase.SendMessage(ctx, message)
			if err != nil {
				log.Error(err)
			}

			msg, err := u.messageUsecase.ReceiveMessage(ctx)
			if err != nil {
				log.Error(err)
			}

			err = ws.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				log.Error(err)
			}
		}
	}()

	return nil
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
	}

	sendMessagesUrl := "/message/"
	api := e.Group("api/v1")
	sendMessages := api.Group(sendMessagesUrl)
	sendMessages.GET("", handler.SendMessagesHandler)

	return handler
}
