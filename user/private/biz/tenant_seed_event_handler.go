package biz

import (
	"context"
	"fmt"
	"github.com/goxiaoy/go-saas-kit/pkg/event/event"
	v1 "github.com/goxiaoy/go-saas-kit/saas/event/v1"
	"github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas/seed"
	"github.com/hibiken/asynq"
	"google.golang.org/protobuf/encoding/protojson"
)

type TenantSeedEventHandler event.Handler

func NewTenantSeedEventHandler(client *asynq.Client) TenantSeedEventHandler {
	msg := &v1.TenantCreatedEvent{}
	return TenantSeedEventHandler(event.ProtoHandler[*v1.TenantCreatedEvent](msg, func(ctx context.Context, msg *v1.TenantCreatedEvent) error {
		t, err := NewUserMigrationTask(msg)
		if err != nil {
			return err
		}
		_, err = client.EnqueueContext(ctx, t)
		return err
	}))
}

func NewUserMigrationTask(msg *v1.TenantCreatedEvent) (*asynq.Task, error) {
	payload, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(JobTypeUserMigration, payload, asynq.Queue(string(ConnName))), nil
}

type UserMigrationTaskHandler func(ctx context.Context, t *asynq.Task) error

func NewUserMigrationTaskHandler(seeder seed.Seeder, sender event.Sender) UserMigrationTaskHandler {
	return func(ctx context.Context, t *asynq.Task) error {
		msg := &v1.TenantCreatedEvent{}
		if err := protojson.Unmarshal(t.Payload(), msg); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}
		extra := map[string]interface{}{}
		if len(msg.AdminEmail) > 0 {
			extra[AdminEmailKey] = msg.AdminEmail
		}
		if len(msg.AdminUsername) > 0 {
			extra[AdminUsernameKey] = msg.AdminUsername
		}
		if len(msg.AdminPassword) > 0 {
			extra[AdminPasswordKey] = msg.AdminPassword
		}
		if err := seeder.Seed(ctx, seed.NewSeedOption().WithTenantId(msg.Id).WithExtra(extra)); err != nil {
			return err
		}
		e := &v1.TenantReadyEvent{
			Id:          msg.Id,
			ServiceName: api.ServiceName,
		}
		ee, _ := event.NewMessageFromProto(e)
		return sender.Send(ctx, ee)
	}
}
