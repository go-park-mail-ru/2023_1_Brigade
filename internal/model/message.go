package model

import "time"

type Message struct {
	Id        string    `json:"id"         db:"id"`
	Body      string    `json:"body"       db:"body"`
	AuthorId  uint64    `json:"author_id"  db:"author_id"`
	ChatId    uint64    `json:"id_chat"    db:"id_chat"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type WebSocketMessage struct {
	Id       string `json:"id"        db:"id"`
	Type     uint64 `json:"type"      db:"type"`
	Body     string `json:"body"      db:"body"`
	AuthorID uint64 `json:"author_id" db:"author_id"`
	ChatID   uint64 `json:"chat_id"   db:"chat_id"`
}

type ProducerMessage struct {
	Id         string    `json:"id"          db:"id"`
	Body       string    `json:"body"        db:"body"`
	AuthorId   uint64    `json:"author_id"   db:"author_id"`
	ChatID     uint64    `json:"chat_id"     db:"chat_id"`
	ReceiverID uint64    `json:"receiver_id" db:"receiver_id"`
	CreatedAt  time.Time `json:"created_at"  db:"created_at"`
}
