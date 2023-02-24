package model

type Chat struct {
	Id       uint64    `json:"id"`
	Members  []User    `json:"members"`
	Messages []Message `json:"messages"`
}
