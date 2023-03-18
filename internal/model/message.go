package model

type Message struct {
	Id        string `json:"id"         db:"id"`
	Body      string `json:"body"       db:"body"`
	AuthorID  string `json:"author_id"  db:"author_id"`
	IsRead    bool   `json:"is_read"    db:"is_read"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}
