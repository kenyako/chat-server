package chatrepo

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/kenyako/chat-server/internal/client/db"
	"github.com/kenyako/chat-server/internal/model"
)

func (r *repo) Create(ctx context.Context, info *model.CreateChat) (int64, error) {

	builderInsert := sq.Insert(chats).
		Columns(chatTitleColumn).
		Values(info.Title).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var chatID int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	builderInsert = sq.Insert(usersChats).
		Columns(chatIDColumn, usersChatsUserIDColumn)

	// генерация ID пользователей, позже будет добавлена ручка для получения ID из базы auth
	for _, username := range info.Usernames {
		userID := int64(gofakeit.Uint8()) + int64(len(username))

		log.Printf("user with id: %d added in chat", userID)

		builderInsert = builderInsert.Values(chatID, userID)
	}

	query, args, err = builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q = db.Query{
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Fatalf("failed to added data in db: %v", err)
	}

	return chatID, nil
}
