package repository

import (
	"database/sql"
	"project/internal/chat"
)

func NewChatMemoryRepository(db *sql.DB) chat.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}
