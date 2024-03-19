package model

import (
	"time"
)

type CreateChat struct {
	Title     string
	Usernames []string
}

type SendMessageRequest struct {
	ChatID   int64
	UserID   int64
	Text     string
	TimeSend time.Time
}
