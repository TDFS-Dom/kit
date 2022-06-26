package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-saas/kit/pkg/event"
	"github.com/go-saas/kit/pkg/lazy"
	"sync"
)

var (
	_ event.Producer = (*Producer)(nil)
	_ event.Consumer = (*Consumer)(nil)
)

func init() {
	event.RegisterConsumer("pulsar", func(ctx context.Context, cfg *event.Config) (event.Consumer, error) {
		opt := pulsar.ClientOptions{URL: cfg.Addr}
		if cfg.Pulsar != nil {
			if cfg.Pulsar.OperationTimeout != nil {
				opt.OperationTimeout = cfg.Pulsar.OperationTimeout.AsDuration()
			}
			if cfg.Pulsar.ConnectionTimeout != nil {
				opt.ConnectionTimeout = cfg.Pulsar.ConnectionTimeout.AsDuration()
			}
		}
		return NewConsumer(cfg.Topic, cfg.Group, opt)
	})

	event.RegisterProducer("pulsar", func(cfg *event.Config) (*event.ProducerMux, error) {
		opt := pulsar.ClientOptions{URL: cfg.Addr}
		if cfg.Pulsar != nil {
			if cfg.Pulsar.OperationTimeout != nil {
				opt.OperationTimeout = cfg.Pulsar.OperationTimeout.AsDuration()
			}
			if cfg.Pulsar.ConnectionTimeout != nil {
				opt.ConnectionTimeout = cfg.Pulsar.ConnectionTimeout.AsDuration()
			}
		}
		sender, err := NewProducer(cfg.Topic, opt)
		if err != nil {
			return nil, err
		}
		res := event.NewProducer(sender)
		return res, nil
	})
}

type Producer struct {
	Client   pulsar.Client
	producer lazy.Of[pulsar.Producer]
	topic    string
}

func NewProducer(topic string, opt pulsar.ClientOptions) (*Producer, error) {
	client, err := pulsar.NewClient(opt)
	if err != nil {
		return nil, err
	}
	p := lazy.New(func(ctx context.Context) (pulsar.Producer, error) {
		return client.CreateProducer(pulsar.ProducerOptions{
			Topic: topic,
		})
	})
	if err != nil {
		return nil, err
	}
	return &Producer{Client: client, producer: p, topic: topic}, nil
}

func (p *Producer) Close() error {
	defer func() {
		//close Client
		p.Client.Close()
	}()
	//close Producer
	if p.producer.Initialized() {
		i, _ := p.producer.Value(context.Background())
		i.Close()
	}
	return nil
}

func (p *Producer) Send(ctx context.Context, msg event.Event) error {
	i, err := p.producer.Value(ctx)
	if err != nil {
		return err
	}
	_, err = i.Send(ctx, toProducerMsg(msg))
	return err
}

func (p *Producer) BatchSend(ctx context.Context, msg []event.Event) error {
	i, err := p.producer.Value(ctx)
	if err != nil {
		return err
	}
	for _, e := range msg {
		_, err = i.Send(ctx, toProducerMsg(e))
		return err
	}
	return nil
}

func toProducerMsg(msg event.Event) *pulsar.ProducerMessage {
	ret := &pulsar.ProducerMessage{
		Payload: msg.Value(),
		Key:     msg.Key(),
	}
	for _, k := range msg.Header().Keys() {
		if ret.Properties == nil {
			ret.Properties = map[string]string{}
		}
		ret.Properties[k] = msg.Header().Get(k)
	}
	return ret
}

type Consumer struct {
	Client   pulsar.Client
	subsName string
	topic    string
	Consumer pulsar.Consumer
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

func NewConsumer(topic, subsName string, opt pulsar.ClientOptions) (*Consumer, error) {
	client, err := pulsar.NewClient(opt)
	if err != nil {
		return nil, err
	}
	return &Consumer{Client: client, subsName: subsName, topic: topic}, nil
}

func (c *Consumer) Process(ctx context.Context, handler event.ConsumerHandler) error {

	ctx, cancel := context.WithCancel(ctx)
	c.cancel = cancel

	//create consumer
	consumer, err := c.Client.Subscribe(pulsar.ConsumerOptions{
		Topic:            c.topic,
		SubscriptionName: c.subsName,
		Type:             pulsar.KeyShared,
	})
	if err != nil {
		return err
	}
	c.Consumer = consumer

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			if msg, err := consumer.Receive(ctx); err != nil {
				log.Error(err)
			} else {
				err := handler.Process(ctx, toMsg(msg))
				if err == nil {
					consumer.Ack(msg)
				} else {
					consumer.Nack(msg)
				}
			}
			// check if context was cancelled, signaling that the Consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()
	return nil

}

func toMsg(msg pulsar.Message) event.Event {
	ret := event.NewMessage(msg.Key(), msg.Payload())
	if msg.Properties() != nil {
		for k, v := range msg.Properties() {
			ret.Header().Set(k, v)
		}
	}
	return ret
}

func (c *Consumer) Close() error {
	c.cancel()
	c.wg.Wait()

	defer func() {
		if c.Consumer != nil {
			c.Consumer.Close()
		}
		//close Client
		c.Client.Close()
	}()

	if c.Consumer != nil {
		err := c.Consumer.Unsubscribe()
		if err != nil {
			return err
		}
	}

	return nil
}
