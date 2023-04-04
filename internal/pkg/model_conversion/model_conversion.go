package model_conversion

import (
	"project/internal/model"
)

func FromAuthorizedUserArrayToUserArray(authorizedUsers []model.AuthorizedUser) []model.User {
	var users []model.User
	for _, user := range authorizedUsers {
		users = append(users, model.User{
			Id:       user.Id,
			Avatar:   user.Avatar,
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
			Status:   user.Status,
		})
	}

	return users
}

func FromAuthorizedUserToUser(user model.AuthorizedUser) model.User {
	return model.User{
		Id:       user.Id,
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
	}
}
