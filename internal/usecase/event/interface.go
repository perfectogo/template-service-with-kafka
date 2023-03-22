package event

import (
	"context"

	"github.com/perfectogo/template-service-with-kafka/internal/entity"
)

type ConsumerConfig interface {
	GetBrokers() []string
	GetTopic() string
	GetGroupID() string
	GetHandler() func(ctx context.Context, key, value []byte) error
}

type BrokerConsumer interface {
	Run()
	RegisterConsumer(config ConsumerConfig)
	Close()
}

type BrokerProducer interface {
	ProduceTodo(ctx context.Context, key string, value *entity.Investment) error
	Close()
}
