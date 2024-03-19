package chat

import (
	"context"

	"github.com/kenyako/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.CreateChat) (int64, error) {

	id, err := s.chatRepository.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
