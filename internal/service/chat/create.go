package chat

import (
	"context"

	"github.com/kenyako/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.CreateChat) (int64, error) {

	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
