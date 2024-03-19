package chat

import (
	"context"

	desc "github.com/kenyako/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, info *desc.DeleteRequest) (*emptypb.Empty, error) {

	err := i.chatService.Delete(ctx, info.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
