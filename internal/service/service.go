package service

import (
	"context"

	"github.com/kenyako/chat-server/internal/model"
)

type ChatAPIService interface {
	Create(ctx context.Context, info *model.CreateChat) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, info *model.SendMessageRequest) error
}
