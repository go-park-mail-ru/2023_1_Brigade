package model

const (
	Group   = 0
	Dialog  = 1
	Channel = 2
)

type Chat struct {
	Id       uint64    `json:"id"       db:"id"`
	Type     uint64    `json:"type"     db:"type"`
	Title    string    `json:"title"    db:"title"`
	Avatar   string    `json:"avatar"   db:"avatar"`
	Members  []User    `json:"members"  db:"members"`
	Messages []Message `json:"messages" db:"messages"`
}

type ChatInListUser struct {
	Id                uint64  `json:"id"                  db:"id"`
	Type              uint64  `json:"type"                db:"type"`
	Title             string  `json:"title"               db:"title"`
	Avatar            string  `json:"avatar"              db:"avatar"`
	LastMessage       Message `json:"last_message"        db:"last_message"`
	LastMessageAuthor User    `json:"last_message_author" db:"last_message_author"`
}

type CreateChat struct {
	Type    uint64   `json:"type"     db:"type"`
	Title   string   `json:"title"    db:"title"`
	Members []uint64 `json:"members"  db:"members"`
}

type ChatMembers struct {
	ChatId   uint64 `json:"id_chat"   db:"id_chat"`
	MemberId uint64 `json:"id_member" db:"id_member"`
}

type UsersChats struct {
	UserId uint64 `json:"id_user" db:"id_user"`
	ChatId uint64 `json:"id_chat" db:"id_chat"`
}
