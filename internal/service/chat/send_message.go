package chat

import (
	"context"

	"github.com/kenyako/chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, info *model.SendMessageRequest) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.chatRepository.SendMessage(ctx, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
