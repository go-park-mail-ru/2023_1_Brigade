package model

type Chat struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt string    `json:"created_at"`
	Members   []User    `json:"participants"`
	Messages  []Message `json:"messages"`
}
