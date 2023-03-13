package repository

import (
	"database/sql"
	"project/internal/group"
)

func NewGroupMemoryRepository(db *sql.DB) group.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}
