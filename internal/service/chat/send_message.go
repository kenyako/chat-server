package chat

import (
	"context"

	"github.com/kenyako/chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, info *model.SendMessageRequest) error {

	err := s.chatRepository.SendMessage(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
