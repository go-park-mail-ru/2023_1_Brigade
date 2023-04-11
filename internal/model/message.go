package model

type Message struct {
	Id       uint64 `json:"id"        db:"id"`
	Body     string `json:"body"      db:"body"`
	AuthorId uint64 `json:"author_id" db:"author_id"`
	ChatId   uint64 `json:"id_chat"   db:"id_chat"`
}

type WebSocketMessage struct {
	Body     string `json:"body"      db:"body"`
	AuthorID uint64 `json:"author_id" db:"author_id"`
	ChatID   uint64 `json:"chat_id"   db:"chat_id"`
}

type ProducerMessage struct {
	Body       string `json:"body"        db:"body"`
	AuthorId   uint64 `json:"author_id"   db:"author_id"`
	ChatID     uint64 `json:"chat_id"     db:"chat_id"`
	ReceiverID uint64 `json:"receiver_id" db:"receiver_id"`
}
