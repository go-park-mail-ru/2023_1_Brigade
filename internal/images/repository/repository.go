package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"io"
	"mime/multipart"
	"os"
	"project/internal/images"
)

func NewImagesMemoryRepository(db *sqlx.DB) images.Repostiory {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r repository) LoadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (string, error) {
	hash := uuid.New().String()
	fileOnDisk, err := os.Create("/home/ubuntu/avatars/" + hash + ".jpg")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(fileOnDisk, file)
	if err != nil {
		return "", err
	}

	url := "https://technogramm.ru/avatars/" + hash + ".jpg"
	_, err = r.db.Query("UPDATE profile SET avatar=$1 WHERE id=$2", url, userID)

	return url, nil
}
