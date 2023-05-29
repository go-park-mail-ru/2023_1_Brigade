package model

const (
	Group   = 0
	Dialog  = 1
	Channel = 2
)

type Chat struct {
	Id          uint64    `json:"id"          db:"id"`
	MasterID    uint64    `json:"master_id  " db:"master_id"`
	Type        uint64    `json:"type"        db:"type"`
	Description string    `json:"description" db:"description"`
	Title       string    `json:"title"       db:"title"`
	Avatar      string    `json:"avatar"      db:"avatar"`
	Members     []User    `json:"members"     db:"members"`
	Messages    []Message `json:"messages"    db:"messages"`
}

type Chats struct {
	Chats []ChatInListUser `json:"chats"`
}

type ChatInListUser struct {
	Id                uint64  `json:"id"                  db:"id"`
	Type              uint64  `json:"type"                db:"type"`
	Title             string  `json:"title"               db:"title"`
	Avatar            string  `json:"avatar"              db:"avatar"`
	Members           []User  `json:"members"             db:"members"`
	LastMessage       Message `json:"last_message"        db:"last_message"`
	LastMessageAuthor User    `json:"last_message_author" db:"last_message_author"`
}

type DBChat struct {
	Id          uint64 `json:"id"          db:"id"`
	MasterID    uint64 `json:"master_id"   db:"master_id"`
	Type        uint64 `json:"type"        db:"type"`
	Description string `json:"description" db:"description"`
	Title       string `json:"title"       db:"title"`
	Avatar      string `json:"avatar"      db:"avatar"`
}

type EditChat struct {
	Id      uint64   `json:"id"      db:"id"`
	Type    uint64   `json:"type"    db:"type"`
	Title   string   `json:"title"   db:"title"`
	Members []uint64 `json:"members" db:"members"`
}

type CreateChat struct {
	Type        uint64   `json:"type"        db:"type"`
	Description string   `json:"description" db:"description"`
	Title       string   `json:"title"       db:"title"`
	Avatar      string   `json:"avatar"      db:"avatar"`
	Members     []uint64 `json:"members"     db:"members"`
}

type ChatMembers struct {
	ChatId   uint64 `json:"id_chat"   db:"id_chat"`
	MemberId uint64 `json:"id_member" db:"id_member"`
}

type ChatMessages struct {
	ChatId    uint64 `json:"id_chat"    db:"id_chat"`
	MessageId string `json:"id_message" db:"id_message"`
}

type FoundedChatsMessagesChannels struct {
	FoundedChats    []ChatInListUser `json:"founded_chats"    db:"founded_chats"`
	FoundedMessages []ChatInListUser `json:"founded_messages" db:"founded_messages"`
	FoundedChannels []ChatInListUser `json:"founded_channels" db:"founded_channels"`
	FoundedContacts []User           `json:"founded_contacts" db:"founded_contacts"`
}
