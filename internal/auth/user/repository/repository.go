//package repository
//
//import (
//	"context"
//	"database/sql"
//	"github.com/jmoiron/sqlx"
//	auth "project/internal/auth/user"
//	"project/internal/model"
//	myErrors "project/internal/pkg/errors"
//)
//
//func NewAuthUserMemoryRepository(db *sqlx.DB) auth.Repository {
//	return &repository{db: db}
//}
//
//type repository struct {
//	db *sqlx.DB
//}
//
//func (r repository) createTechnogrammChat(ctx context.Context, user model.AuthorizedUser) error {
//	tx, err := r.db.BeginTxx(ctx, nil)
//	if err != nil {
//		return err
//	}
//
//	var chat model.DBChat
//	err = tx.QueryRowContext(ctx, `INSERT INTO chat (type, avatar, title)
//    VALUES (0, 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm') RETURNING id;`).Scan(&chat.Id)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	_, err = tx.ExecContext(ctx, `INSERT INTO message (id, body, id_chat, author_id, created_at)
//    VALUES ('1337', 'Привет, это технограмм!', (SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = $2), now() at time zone 'Europe/Moscow');`, chat.Id, 0)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	_, err = tx.ExecContext(ctx, `INSERT INTO chat_messages (id_chat, id_message)
//    VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM message WHERE id ='1337'));`, chat.Id)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	_, err = tx.ExecContext(ctx, `INSERT INTO chat_members (id_chat, id_member)
//    VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = $2));`, chat.Id, user.Id)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	_, err = tx.ExecContext(ctx, `INSERT INTO chat_members (id_chat, id_member)
//    VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = 0));`, chat.Id)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r repository) CreateUser(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
//	err := r.db.QueryRowContext(ctx, `INSERT INTO profile (avatar, username, nickname, email, status, password) `+
//		`VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
//		"", user.Username, user.Nickname, user.Email, user.Status, user.Password).Scan(&user.Id)
//	if err != nil {
//		return model.AuthorizedUser{}, err
//	}
//
//	err = r.createTechnogrammChat(ctx, user)
//	if err != nil {
//		return model.AuthorizedUser{}, err
//	}
//
//	return user, nil
//}
//
//func (r repository) UpdateUserAvatar(ctx context.Context, url string, userID uint64) (model.AuthorizedUser, error) {
//	var user model.AuthorizedUser
//	err := r.db.GetContext(ctx, &user, `UPDATE profile SET avatar=$1 WHERE id=$2 RETURNING *`, url, userID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return model.AuthorizedUser{}, myErrors.ErrUserNotFound
//		}
//		return model.AuthorizedUser{}, err
//	}
//
//	return user, nil
//}
//
//func (r repository) CheckCorrectPassword(ctx context.Context, email string, password string) error {
//	var exists bool
//	err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1 AND password=$2)", email, password)
//
//	if err != nil {
//		return err
//	}
//	if !exists {
//		return myErrors.ErrIncorrectPassword
//	}
//
//	return nil
//}
//
//func (r repository) CheckExistEmail(ctx context.Context, email string) error {
//	var exists bool
//	err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1)", email)
//
//	if err != nil {
//		return err
//	}
//	if !exists {
//		return myErrors.ErrEmailNotFound
//	}
//
//	return nil
//}
//
//func (r repository) CheckExistUsername(ctx context.Context, username string) error {
//	var exists bool
//	err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE username=$1)", username)
//
//	if err != nil {
//		return err
//	}
//	if !exists {
//		return myErrors.ErrUsernameNotFound
//	}
//
//	return nil
//}

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	auth "project/internal/auth/user"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

func NewAuthUserMemoryRepository(db *sqlx.DB) auth.Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r repository) createTechnogrammChat(user model.AuthorizedUser) {
	var chat model.DBChat
	err := r.db.QueryRow(`INSERT INTO chat (type, avatar, title) VALUES (0, 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm') RETURNING id;`).Scan(&chat.Id)
	if err != nil {
		log.Error(err)
	}

	_, err = r.db.Exec(`INSERT INTO message (id, body, id_chat, author_id, created_at)
	VALUES ('1337', 'Привет, это технограмм!', (SELECT id FROM chat
	WHERE id = $1), (SELECT id FROM profile
	WHERE id = $2), now() at time zone 'Europe/Moscow');`, chat.Id, 0)
	if err != nil {
		log.Error(err)
	}

	_, err = r.db.Exec(`INSERT INTO chat_messages (id_chat, id_message) VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM message WHERE id ='1337'));`, chat.Id)
	if err != nil {
		log.Error(err)
	}

	_, err = r.db.Exec(`INSERT INTO chat_members (id_chat, id_member) VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = $2));`, chat.Id, user.Id)
	if err != nil {
		log.Error(err)
	}

	_, err = r.db.Exec(`INSERT INTO chat_members (id_chat, id_member) VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = 0));`, chat.Id)
	if err != nil {
		log.Error(err)
	}
}

func (r repository) CreateUser(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	row, err := r.db.Query(`INSERT INTO profile (avatar, username, nickname, email, status, password) `+
		`VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		"", user.Username, user.Nickname, user.Email, user.Status, user.Password)
	defer row.Close()

	if err != nil {
		return model.AuthorizedUser{}, err
	}
	if row.Next() {
		err = row.Scan(&user.Id)
		if err != nil {
			return model.AuthorizedUser{}, err
		}
	}

	r.createTechnogrammChat(user)

	return user, nil
}

func (r repository) UpdateUserAvatar(ctx context.Context, url string, userID uint64) (model.AuthorizedUser, error) {
	result, err := r.db.Exec("UPDATE profile SET avatar=$1 WHERE id=$2", url, userID)
	if err != nil {
		return model.AuthorizedUser{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.AuthorizedUser{}, err
	}

	if rowsAffected == 0 {
		return model.AuthorizedUser{}, myErrors.ErrUserNotFound
	}

	var user model.AuthorizedUser
	err = r.db.Get(&user, "SELECT * FROM profile WHERE id=$1", userID)
	if err != nil {
		return model.AuthorizedUser{}, err
	}

	return user, nil
}

func (r repository) CheckCorrectPassword(ctx context.Context, email string, password string) error {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1 AND password=$2)", email, password)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrIncorrectPassword
	}

	return nil
}

func (r repository) CheckExistEmail(ctx context.Context, email string) error {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1)", email)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrEmailNotFound
	}

	return nil
}

func (r repository) CheckExistUsername(ctx context.Context, username string) error {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM profile WHERE username=$1)", username)

	if err != nil {
		return err
	}
	if !exists {
		return myErrors.ErrUsernameNotFound
	}

	return nil
}
