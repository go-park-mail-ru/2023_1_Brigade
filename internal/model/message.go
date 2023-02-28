package model

type Message struct {
	Id        uint64 `json:"id"`
	AuthorId  uint64 `json:"author_id"`
	Body      string `json:"body"`
	Media     string `json:"media"` // ??
	CreatedAt string `json:"created_at"`
	IsRead    bool   `json:"is_read"`
}
