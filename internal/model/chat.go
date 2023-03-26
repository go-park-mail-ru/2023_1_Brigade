package model

const (
	Group  = 0
	Dialog = 1
)

type Chat struct {
	Id       uint64    `json:"id"       db:"id"`
	Title    string    `json:"title"    db:"title"`
	Members  []User    `json:"members"  db:"-"`
	Messages []Message `json:"messages" db:"-"`
}

type CreateChat struct {
	Title   string   `json:"title"    db:"title"`
	Members []uint64 `json:"members"  db:"-"`
}

//type ChatA struct {
//	Id    uint64 `json:"id"       db:"id"`
//	Title string `json:"title"    db:"title"`
//}

type ChatMembers struct {
	ChatId   uint64 `json:"id_chat"   db:"id_chat"`
	MemberId uint64 `json:"id_member" db:"id_member"`
}
