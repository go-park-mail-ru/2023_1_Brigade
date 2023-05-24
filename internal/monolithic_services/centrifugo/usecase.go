package centrifugo

import (
	"context"
	"github.com/centrifugal/centrifuge-go"
)

type Centrifugo interface {
	Publish(ctx context.Context, channel string, data []byte) (centrifuge.PublishResult, error)
	GetSubscription(channel string) (CentrifugoSubscription, bool)
	Close()
}

type CentrifugoSubscription interface {
	OnPublication(handler centrifuge.PublicationHandler)
}
