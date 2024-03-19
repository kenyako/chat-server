package chat

import (
	"github.com/kenyako/chat-server/internal/service"
	desc "github.com/kenyako/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatAPIServer
	chatService service.ChatAPIService
}

func NewImplementation(chatService service.ChatAPIService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
