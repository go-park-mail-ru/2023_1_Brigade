package centrifugo

import (
	"context"
	"github.com/centrifugal/centrifuge-go"
)

type Centrifugo interface {
	Publish(ctx context.Context, channel string, data []byte) (centrifuge.PublishResult, error)
	OnPublication(handler centrifuge.ServerPublicationHandler)
	GetSubscription(channel string) (*centrifuge.Subscription, bool)
	Close()
}
