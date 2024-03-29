package ws

import (
	"context"
	"github.com/centrifugal/centrifuge-go"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"project/internal/config"
	messagesMock "project/internal/microservices/messages/usecase/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"strings"
	"testing"
	"time"
)

type testCase struct {
	name           string
	wsBody         []byte
	producerBody   []byte
	producerResult error
	consumerResult error
}

type WsHandler struct {
	handler echo.HandlerFunc
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	c := e.NewContext(r, w)
	c.Set("session", model.Session{UserId: 1})
	forever := make(chan struct{})
	_ = h.handler(c)
	<-forever
}

func TestHandlers_WSHandler(t *testing.T) {
	centrifugo := config.Centrifugo{
		ConnAddr:    "ws://localhost:8900/connection/websocket",
		ChannelName: "channel",
	}

	wsMessage := model.WebSocketMessage{
		Id:       "",
		Type:     config.Chat,
		Body:     "Hello world!",
		AuthorID: 1,
		ChatID:   1,
	}

	producerMessage := model.ProducerMessage{
		Id:         uuid.New().String(),
		Type:       config.Chat,
		Body:       "Hello world!",
		AuthorId:   1,
		ChatID:     1,
		ReceiverID: 1,
	}

	wsMessageJson, err := easyjson.Marshal(wsMessage)
	assert.NoError(t, err)

	producerMessageJson, err := easyjson.Marshal(producerMessage)
	assert.NoError(t, err)

	tests := []testCase{
		{
			name:           "handler ok worked",
			wsBody:         wsMessageJson,
			producerBody:   producerMessageJson,
			producerResult: nil,
			consumerResult: nil,
		},
		{
			name:           "producer return error",
			wsBody:         wsMessageJson,
			producerBody:   producerMessageJson,
			producerResult: myErrors.ErrInternal,
			consumerResult: nil,
		},
		{
			name:           "consumer return error",
			wsBody:         wsMessageJson,
			producerBody:   producerMessageJson,
			producerResult: myErrors.ErrInternal,
			consumerResult: nil,
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()

	messagesUsecase := messagesMock.NewMockUsecase(ctl)

	c := centrifuge.NewJsonClient(centrifugo.ConnAddr, centrifuge.Config{})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		c.Close()
		log.Fatal()
	}()

	err = c.Connect()
	if err != nil {
		log.Error(err)
	}

	sub, err := c.NewSubscription(centrifugo.ChannelName, centrifuge.SubscriptionConfig{
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

	handler, err := NewMessagesHandler(e, messagesUsecase, c, centrifugo.ChannelName)
	assert.NoError(t, err)

	h := WsHandler{handler: handler.SendMessagesHandler}
	server := httptest.NewServer(http.HandlerFunc(h.ServeHTTP))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/message/"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err, err)
	defer func() {
		err = ws.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	for _, test := range tests {
		messagesUsecase.EXPECT().PutInProducer(context.TODO(), test.wsBody).Return(test.producerResult).AnyTimes()

		err = ws.WriteMessage(websocket.TextMessage, test.wsBody)
		require.NoError(t, err, test.name)

		if test.producerResult != nil {
			continue
		}

		if test.consumerResult != nil {
			continue
		}

		sub, subscribed := c.GetSubscription(centrifugo.ChannelName)
		require.Equal(t, true, subscribed)

		_, err := sub.Publish(context.TODO(), test.producerBody)
		require.NoError(t, err)
		time.Sleep(250 * time.Millisecond)

		_, msg, err := ws.ReadMessage()
		require.Equal(t, test.producerBody, msg)
		require.NoError(t, err)
	}
}
