package chatrepo

import (
	"github.com/kenyako/chat-server/internal/repository"
	"github.com/kenyako/platform_common/pkg/postgres"
)

const (
	chats         = "chats"
	chatsMessages = "chat_messages"
	usersChats    = "users_chats"
)

const (
	chatIDColumn    = "id"
	chatTitleColumn = "title"
)

const (
	chatMessagesChatIDColumn = "chat_id"
	chatMessagesTextColumn   = "text"
	chatMessagesFromColumn   = "user_id"
	chatMessagesTimeColumn   = "time_sent"
)

const (
	usersChatsUserIDColumn = "user_id"
	usersChatsChatIDColumn = "chat_id"
)

type repo struct {
	db postgres.Client
}

func NewRepository(db postgres.Client) repository.ChatAPIRepo {

	return &repo{
		db: db,
	}
}
