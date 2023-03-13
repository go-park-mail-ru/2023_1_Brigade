package usecase

import (
	"project/internal/chat"
)

type usecase struct {
	repo chat.Repository
}

func NewChatUsecase(chatRepo chat.Repository) chat.Usecase {
	return &usecase{repo: chatRepo}
}
