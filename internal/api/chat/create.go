package chat

import (
	"context"

	"github.com/kenyako/chat-server/internal/converter"
	desc "github.com/kenyako/chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, info *desc.CreateRequest) (*desc.CreateResponse, error) {

	id, err := i.chatService.Create(ctx, converter.ToChatCreateFromDesc(info))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
