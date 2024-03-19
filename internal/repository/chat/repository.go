package chatrepo

import (
	"github.com/kenyako/chat-server/internal/client/db"
	"github.com/kenyako/chat-server/internal/repository"
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
	db db.Client
}

func NewRepository(db db.Client) repository.ChatAPIRepo {

	return &repo{
		db: db,
	}
}
