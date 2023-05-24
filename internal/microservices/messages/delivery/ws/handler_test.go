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

//type Centrifugo interface {
//	Publish(ctx context.Context, channel string, data []byte) (centrifuge.PublishResult, error)
//	OnPublication(handler centrifuge.ServerPublicationHandler)
//}
//
//type centrifugo struct{}
//
//func (m *centrifugo) foo() {}
//
//type mockCentrifugo struct{}
//
//func (m *mockCentrifugo) foo() {}
//
//func NewMock() CentrifugoInt {
//	return &mockCentrifugo{}
//}
//
//func newMock() *centrifuge.Client {
//	return &mock{}
//}
//
//func (m *mock) Connect() error
//func (m *mock) Disconnect() error
//func (m *mock) Close()
//func (m *mock) State() State
//func (m *mock) NewSubscription(channel string, config ...SubscriptionConfig) (*Subscription, error)
//func (m *mock) RemoveSubscription(sub *Subscription) error
//func (m *mock) GetSubscription(channel string) (*Subscription, bool)
//func (m *mock) Subscriptions() map[string]*Subscription
//func (m *mock) Send(ctx context.Context, data []byte) error
//func (m *mock) RPC(ctx context.Context, method string, data []byte) (RPCResult, error)
//func (m *mock) nextCmdID() uint32
//func (m *mock) isConnected() bool
//func (m *mock) isClosed() bool
//func (m *mock) isSubscribed(channel string) bool
//func (m *mock) sendRPC(ctx context.Context, method string, data []byte, fn func(RPCResult, error))
//func (m *mock) moveToDisconnected(code uint32, reason string)
//func (m *mock) moveToConnecting(code uint32, reason string)
//func (m *mock) moveToClosed()
//func (m *mock) handleError(err error)
//func (m *mock) clearConnectedState()
//func (m *mock) handleDisconnect(d *disconnect)
//func (m *mock) waitServerPing(disconnectCh chan struct{}, pingInterval uint32)
//func (m *mock) readOnce(t transport) error
//func (m *mock) reader(t transport, disconnectCh chan struct{})
//func (m *mock) runHandlerSync(fn func())
//func (m *mock) runHandlerAsync(fn func())
//func (m *mock) handle(reply *protocol.Reply)
//func (m *mock) handleMessage(msg *protocol.Message) error
//func (m *mock) handlePush(push *protocol.Push)
//func (m *mock) handleServerPublication(channel string, pub *protocol.Publication)
//func (m *mock) handleServerJoin(channel string, join *protocol.Join)
//func (m *mock) handleServerLeave(channel string, leave *protocol.Leave)
//func (m *mock) handleServerSub(channel string, sub *protocol.Subscribe)
//func (m *mock) handleServerUnsub(channel string, _ *protocol.Unsubscribe)
//func (m *mock) getReconnectDelay() time.Duration
//func (m *mock) startReconnecting() error
//func (m *mock) startConnecting() error
//func (m *mock) resubscribe()
//func (m *mock) refreshToken() (string, error)
//func (m *mock) sendRefresh()
//func (m *mock) handleRefreshError(err error)
//func (m *mock) sendSubRefresh(channel string, token string, fn func(*protocol.SubRefreshResult, error))
//func (m *mock) sendConnect(fn func(*protocol.ConnectResult, error)) error
//func (m *mock) sendSubscribe(channel string, data []byte, recover bool, streamPos StreamPosition, token string, positioned bool, recoverable bool, joinLeave bool, fn func(res *protocol.SubscribeResult, err error)) error
//func (m *mock) nextFutureID() uint64
//func (m *mock) resolveConnectFutures(err error)
//func (m *mock) onConnect(fn func(err error))
//func (m *mock) Publish(ctx context.Context, channel string, data []byte) (PublishResult, error)
//func (m *mock) publish(ctx context.Context, channel string, data []byte, fn func(PublishResult, error))
//func (m *mock) sendPublish(channel string, data []byte, fn func(PublishResult, error))
//func (m *mock) History(ctx context.Context, channel string, opts ...HistoryOption) (HistoryResult, error)
//func (m *mock) history(ctx context.Context, channel string, opts HistoryOptions, fn func(HistoryResult, error))
//func (m *mock) sendHistory(channel string, opts HistoryOptions, fn func(HistoryResult, error))
//func (m *mock) Presence(ctx context.Context, channel string) (PresenceResult, error)
//func (m *mock) presence(ctx context.Context, channel string, fn func(PresenceResult, error))
//func (m *mock) sendPresence(channel string, fn func(PresenceResult, error))
//func (m *mock) PresenceStats(ctx context.Context, channel string) (PresenceStatsResult, error)
//func (m *mock) presenceStats(ctx context.Context, channel string, fn func(PresenceStatsResult, error))
//func (m *mock) sendPresenceStats(channel string, fn func(PresenceStatsResult, error))
//func (m *mock) unsubscribe(channel string, fn func(UnsubscribeResult, error))
//func (m *mock) sendUnsubscribe(channel string, fn func(UnsubscribeResult, error))
//func (m *mock) sendAsync(cmd *protocol.Command, cb func(*protocol.Reply, error)) error
//func (m *mock) send(cmd *protocol.Command) error
//func (m *mock) addRequest(id uint32, cb func(*protocol.Reply, error))
//func (m *mock) removeRequest(id uint32)
//func (m *mock) OnConnected(handler ConnectedHandler)
//func (m *mock) OnConnecting(handler ConnectingHandler)
//func (m *mock) OnDisconnected(handler DisconnectHandler)
//func (m *mock) OnError(handler ErrorHandler)
//func (m *mock) OnMessage(handler MessageHandler)
//func (m *mock) OnPublication(handler ServerPublicationHandler)
//func (m *mock) OnSubscribed(handler ServerSubscribedHandler)
//func (m *mock) OnSubscribing(handler ServerSubscribingHandler)
//func (m *mock) OnUnsubscribed(handler ServerUnsubscribedHandler)
//func (m *mock) OnJoin(handler ServerJoinHandler)
//func (m *mock) OnLeave(handler ServerLeaveHandler)

//func (a *scas)

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

	//rrr := &mock{}
	//var rc centrifuge.Client
	//rc = &mock{}

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
