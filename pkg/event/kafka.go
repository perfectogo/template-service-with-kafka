package event

import (
	"context"
	"fmt"
	"log"

	// "go_boilerplate/pkg/logger"
	"sync"

	"github.com/Shopify/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/perfectogo/template-service-with-kafka/config"
	"github.com/perfectogo/template-service-with-kafka/pkg/logger"
)

type Kafka struct {
	log          logger.Logger
	cfg          config.Config
	consumers    map[string]*Consumer
	publishers   map[string]*Publisher
	saramaConfig *sarama.Config
}

func NewKafka(cfg config.Config, log logger.Logger) (*Kafka, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	kafka := &Kafka{
		log:          log,
		cfg:          cfg,
		consumers:    make(map[string]*Consumer),
		publishers:   make(map[string]*Publisher),
		saramaConfig: saramaConfig,
	}

	return kafka, nil
}

// RunConsumers ...
func (r *Kafka) RunConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, consumer := range r.consumers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Consumer) {
			defer wg.Done()

			err := c.cloudEventClient.StartReceiver(ctx, func(ctx context.Context, event cloudevents.Event) {
				resp := c.handler(ctx, event)
				err := event.SetData(cloudevents.ApplicationJSON, resp)
				if err != nil {
					r.log.Error("Failed to set data")
				}

				if !resp.NoResponse {
					err = r.Push("v1.websocket_service.response", event)
					if err != nil {
						r.log.Error("Failed to push")
					}
				}
			})

			log.Panic("Failed to start consumer", err)
		}(&wg, consumer)
		fmt.Println("Key:", consumer.topic, "=>", "consumer:", consumer)
	}

	wg.Wait()
}
