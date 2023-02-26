package model

type Chat struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt string    `json:"created_at"`
	Members   []User    `json:"participants"`
	Messages  []Message `json:"messages"`
}
