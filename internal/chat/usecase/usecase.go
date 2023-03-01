package usecase

import (
	"context"
	"project/internal/chat"
	"project/internal/pkg/http_utils"
)

type usecaseImpl struct {
	repo chat.Repository
}

func NewChatUsecase(chatRepo chat.Repository) chat.Usecase {
	return &usecaseImpl{repo: chatRepo}
}

func (u *usecaseImpl) CreateChat(ctx context.Context, jsonChatData []byte) http_utils.Response {
	//chat := model.Chat{}
	//err := json.Unmarshal(jsonChatData, &chat)
	//fmt.Println("POST USECASE")
	//err = u.repo.InsertChatInDB(ctx, chat)
	//
	//if err != nil {
	//	return err
	//}
	//
	//return nil
	return http_utils.Response{}
}

func (u *usecaseImpl) GetChatById(ctx context.Context, chatID int) http_utils.Response {
	return u.repo.GetChatInDB(ctx, chatID)
}

func (u *usecaseImpl) GetAllChats(ctx context.Context) http_utils.Response {
	return u.repo.GetAllChatsInDB(ctx)
}

func (u *usecaseImpl) DeleteChatById(ctx context.Context, chatID int) http_utils.Response {
	return u.repo.DeleteChatInDB(ctx, chatID)
}
