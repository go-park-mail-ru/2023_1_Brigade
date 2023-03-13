package usecase

import (
	"project/internal/group"
)

type usecase struct {
	repo group.Repository
}

func NewGroupUsecase(groupRepo group.Repository) group.Usecase {
	return &usecase{repo: groupRepo}
}
