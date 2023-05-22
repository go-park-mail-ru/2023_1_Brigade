package model

type Message struct {
	Id        string `json:"id"         db:"id"`
	ImageUrl  string `json:"image_url"  db:"image_url"`
	Type      uint64 `json:"type"       db:"type"`
	Body      string `json:"body"       db:"body"`
	AuthorId  uint64 `json:"author_id"  db:"author_id"`
	ChatId    uint64 `json:"id_chat"    db:"id_chat"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type WebSocketMessage struct {
	Id       string `json:"id"        db:"id"`
	ImageUrl string `json:"image_url" db:"image_url"`
	Action   uint64 `json:"action"    db:"action"`
	Type     uint64 `json:"type"      db:"type"`
	Body     string `json:"body"      db:"body"`
	AuthorID uint64 `json:"author_id" db:"author_id"`
	ChatID   uint64 `json:"chat_id"   db:"chat_id"`
}

type ProducerMessage struct {
	Id         string `json:"id"          db:"id"`
	ImageUrl   string `json:"image_url"   db:"image_url"`
	Action     uint64 `json:"action"      db:"action"`
	Type       uint64 `json:"type"        db:"type"`
	Body       string `json:"body"        db:"body"`
	AuthorId   uint64 `json:"author_id"   db:"author_id"`
	ChatID     uint64 `json:"chat_id"     db:"chat_id"`
	ReceiverID uint64 `json:"receiver_id" db:"receiver_id"`
	CreatedAt  string `json:"created_at"  db:"created_at"`
}

type Notification struct {
	AuthorID       uint64 `json:"author_id"  db:"author_id"`
	ChatName       string `json:"chat_name" db:"chat_name"`
	ChatAvatar     string `json:"chat_avatar" db:"chat_avatar"`
	AuthorNickname string `json:"author_nickname" db:"author_nickname"`
	Body           string `json:"body" db:"body"`
}
