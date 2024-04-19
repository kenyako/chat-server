package chat

import (
	"github.com/kenyako/chat-server/internal/repository"
	"github.com/kenyako/chat-server/internal/service"
	"github.com/kenyako/platform_common/pkg/postgres"
)

type serv struct {
	chatRepository repository.ChatAPIRepo
	txManager      postgres.TxManager
}

func NewServiceChat(chatRepository repository.ChatAPIRepo, txManager postgres.TxManager) service.ChatAPIService {

	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
