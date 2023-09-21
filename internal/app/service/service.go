package service

import (
	"context"
	"github.com/samarec1812/audit-service/internal/app/entity"
)

type EventRepository interface {
	Save(context.Context, entity.Event) error
}

type App interface {
	SaveEvent(context.Context, map[string]any, map[string]any) error
}

type app struct {
	eventRepo EventRepository
}

func NewApp(eventRepo EventRepository) App {
	return &app{
		eventRepo: eventRepo,
	}
}

func (a *app) SaveEvent(ctx context.Context, headers, body map[string]any) error {
	err := entity.Validate(headers, body)
	if err != nil {
		return err
	}
	event := entity.NewEvent(headers, body)

	return a.eventRepo.Save(ctx, *event)
}
