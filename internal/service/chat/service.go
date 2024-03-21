package chat

import (
	"github.com/kenyako/chat-server/internal/client/db"
	"github.com/kenyako/chat-server/internal/repository"
	"github.com/kenyako/chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatAPIRepo
	txManager      db.TxManager
}

func NewServiceChat(chatRepository repository.ChatAPIRepo, txManager db.TxManager) service.ChatAPIService {

	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
