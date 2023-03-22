package model

const (
	Group  = 0
	Dialog = 1
)

type Chat struct {
	Id        uint64 `json:"id"         db:"id"`
	Type      int    `json:"type"       db:"type"`
	Name      string `json:"name"       db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Members   []User `json:"members"    db:"-"`
	Masters   []User `json:"masters"    db:"-"`
}
