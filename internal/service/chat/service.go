package chat

import (
	"github.com/kenyako/chat-server/internal/repository"
	"github.com/kenyako/chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatAPIRepo
}

func NewServiceChat(chatRepository repository.ChatAPIRepo) service.ChatAPIService {

	return &serv{
		chatRepository: chatRepository,
	}
}
