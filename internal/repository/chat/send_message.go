package chatrepo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kenyako/chat-server/internal/model"
	"github.com/kenyako/platform_common/pkg/postgres"
)

func (r *repo) SendMessage(ctx context.Context, info *model.SendMessageRequest) error {

	builderInsert := sq.Insert(chatsMessages).
		Columns(chatMessagesChatIDColumn, chatMessagesFromColumn, chatMessagesTextColumn, chatMessagesTimeColumn).
		Values(info.ChatID, info.UserID, info.Text, time.Now())

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := postgres.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
