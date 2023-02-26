package model

type IsRead struct {
	Member uint64 `json:"id"`
	IsRead bool   `json:"is_read"`
}

type Message struct {
	ID        uint64   `json:"id"`
	AuthorID  uint64   `json:"author_id"`
	Body      string   `json:"body"`
	Media     string   `json:"media"` // ??
	CreatedAt string   `json:"created_at"`
	IsRead    []IsRead `json:"is_read"`
}
