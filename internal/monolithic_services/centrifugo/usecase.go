package centrifugo

import (
	"context"
	"github.com/centrifugal/centrifuge-go"
)

type Centrifugo interface {
	Publish(ctx context.Context, channel string, data []byte) (centrifuge.PublishResult, error)
	GetSubscription(channel string) (*centrifuge.Subscription, bool)
	Close()
}

//type a struct {}
//
//func (r *a) OnPublication(handler centrifuge.PublicationHandler)  {
//
//}
//
//func (r *a) Publish(ctx context.Context, data []byte) (centrifuge.PublishResult, error) {

//}

//OnPublication(handler centrifuge.PublicationHandler)
//Publish(ctx context.Context, data []byte) (centrifuge.PublishResult, error)

type CentrifugoSubscription interface {
	//OnSubscribing(handler centrifuge.SubscribingHandler)
	//OnSubscribed(handler centrifuge.SubscribedHandler)
	//OnUnsubscribed(handler centrifuge.UnsubscribedHandler)
	//OnError(handler centrifuge.SubscriptionErrorHandler)
	//OnPublication(handler centrifuge.PublicationHandler)
	//OnJoin(handler centrifuge.JoinHandler)
	//OnLeave(handler centrifuge.LeaveHandler)
	//State() centrifuge.SubState
	//nextFutureID() uint64
	//resolveSubFutures(err error)
	//Publish(ctx context.Context, data []byte) (centrifuge.PublishResult, error)
	//History(ctx context.Context, opts ...centrifuge.HistoryOption) (centrifuge.HistoryResult, error)
	//Presence(ctx context.Context) (centrifuge.PresenceResult, error)
	//PresenceStats(ctx context.Context) (centrifuge.PresenceStatsResult, error)
	//onSubscribe(fn func(err error))
	//publish(ctx context.Context, data []byte, fn func(centrifuge.PublishResult, error))
	//history(ctx context.Context, opts centrifuge.HistoryOptions, fn func(centrifuge.HistoryResult, error))
	//presence(ctx context.Context, fn func(centrifuge.PresenceResult, error))
	//presenceStats(ctx context.Context, fn func(centrifuge.PresenceStatsResult, error))
	//Unsubscribe() error
	//unsubscribe(code uint32, reason string, sendUnsubscribe bool)
	//Subscribe() error
	//moveToUnsubscribed(code uint32, reason string)
	//moveToSubscribing(code uint32, reason string)
	//moveToSubscribed(res *protocol.SubscribeResult)
	//scheduleResubscribe()
	//subscribeError(err error)
	//emitError(err error)
	//handlePublication(pub *protocol.Publication)
	//handleJoin(info *protocol.ClientInfo)
	//handleLeave(info *protocol.ClientInfo)
	//handleUnsubscribe(unsubscribe *protocol.Unsubscribe)
	//resubscribe()
	//getSubscriptionToken(channel string) (string, error)
	//scheduleSubRefresh(ttl uint32)
	OnPublication(handler centrifuge.PublicationHandler)
	Publish(ctx context.Context, data []byte) (centrifuge.PublishResult, error)
}
