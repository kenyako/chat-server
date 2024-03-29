package chatrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kenyako/chat-server/internal/client/db"
)

func (r *repo) Delete(ctx context.Context, id int64) error {

	builderDelete := sq.Delete(usersChats).
		Where(sq.Eq{usersChatsChatIDColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	builderDelete = sq.Delete(chats).
		Where(sq.Eq{chatIDColumn: id})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	builderDelete = sq.Delete(chatsMessages).
		Where(sq.Eq{chatMessagesChatIDColumn: id})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
