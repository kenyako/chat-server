package chat

import (
	"context"

	"github.com/kenyako/chat-server/internal/converter"
	desc "github.com/kenyako/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, info *desc.SendRequest) (*emptypb.Empty, error) {

	err := i.chatService.SendMessage(ctx, converter.ToSendRequestFromDesc(info))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
