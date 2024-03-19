package converter

import (
	"github.com/kenyako/chat-server/internal/model"
	desc "github.com/kenyako/chat-server/pkg/chat_v1"
)

func ToChatCreateFromDesc(info *desc.CreateRequest) *model.CreateChat {

	return &model.CreateChat{
		Title:     info.Title,
		Usernames: info.Usernames,
	}
}

func ToSendRequestFromDesc(info *desc.SendRequest) *model.SendMessageRequest {

	return &model.SendMessageRequest{
		ChatID:   info.ChatID,
		UserID:   info.From,
		Text:     info.Text,
		TimeSend: info.Timestamp.AsTime(),
	}
}
