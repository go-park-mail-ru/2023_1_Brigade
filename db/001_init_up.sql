-- таблица пользователей
-- 2 НФ, т.к. email определяется значением username
CREATE TABLE IF NOT EXISTS profile (
    id       SERIAL PRIMARY KEY,
    avatar   TEXT,
    username TEXT,
    nickname TEXT,
    email    TEXT,
    status   TEXT,
    password TEXT
);

-- таблица связи 1:1 куки и пользователя
CREATE TABLE IF NOT EXISTS session (
    cookie     TEXT,
    profile_id INTEGER,
    FOREIGN KEY (profile_id) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS chat (
    id    SERIAL PRIMARY KEY,
    master_id    INTEGER,
    type INTEGER,
    description TEXT,
    avatar TEXT,
    title TEXT,
    FOREIGN KEY (master_id) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS message (
    id         TEXT PRIMARY KEY,
    type       INTEGER,
    body       TEXT,
    id_chat    INTEGER,
    author_id  INTEGER,
    created_at TEXT,
    FOREIGN KEY (author_id) REFERENCES profile(id),
    FOREIGN KEY (id_chat)   REFERENCES chat(id)
);

-- таблица связи M:M чатов и пользователей (участников чата)
CREATE TABLE IF NOT EXISTS chat_members (
    id_chat   INTEGER,
    id_member INTEGER,
    FOREIGN KEY (id_chat)   REFERENCES chat(id),
    FOREIGN KEY (id_member) REFERENCES profile(id)
);

-- таблица связи M:M пользователей и пользователей в списке контактов
CREATE TABLE IF NOT EXISTS user_contacts (
    id_user    INTEGER,
    id_contact INTEGER,
    FOREIGN KEY (id_user)    REFERENCES profile(id),
    FOREIGN KEY (id_contact) REFERENCES profile(id)
);

-- таблица связи 1:M чатов и сообщений
CREATE TABLE IF NOT EXISTS chat_messages (
    id_chat    INTEGER,
    id_message TEXT,
    FOREIGN KEY (id_chat)   REFERENCES chat(id),
    FOREIGN KEY (id_message) REFERENCES message(id)
);

-- таблица ссылок на файлы, прикрепленные к сообщениям
CREATE TABLE IF NOT EXISTS attachments (
    id_message TEXT,
    url        TEXT,
    name       TEXT,
    FOREIGN KEY (id_message) REFERENCES message(id)
);

-- перед каждым индексом указаны ссылки на запросы, для которых они используются

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/auth/repository/repository.go#L46
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/auth/repository/repository.go#L60
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/user/repository/repository.go#L52
CREATE INDEX IF NOT EXISTS idx_profile_email_username
ON profile (email, username);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/auth/repository/repository.go#L74 
CREATE INDEX IF NOT EXISTS idx_profile_username
ON profile (username);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L81
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L122
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L199
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L259
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L277
CREATE INDEX IF NOT EXISTS idx_chat_members_id_chat_id_member
ON chat_members (id_chat, id_member);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L107
CREATE INDEX IF NOT EXISTS idx_chat_members_id_member
ON chat_members (id_member);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L193
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L135
CREATE INDEX IF NOT EXISTS idx_chat_messages_id_chat
ON chat_messages (id_chat);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L64
CREATE INDEX IF NOT EXISTS idx_chat_messages_id_message
ON chat_messages (id_message);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L205
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L194
CREATE INDEX IF NOT EXISTS idx_message_id_chat_created_at
ON message (id_chat, created_at);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L258
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L276
CREATE INDEX IF NOT EXISTS idx_chat_title 
ON chat USING GIN (title gin_trgm_ops);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L35
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L70
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L108
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L205
CREATE INDEX IF NOT EXISTS idx_attachments_id_message
ON attachments (id_message);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L241
CREATE INDEX IF NOT EXISTS idx_message_body 
ON message USING GIN (body gin_trgm_ops);

INSERT INTO profile (id, avatar, username, nickname, email, status, password)
VALUES (0, 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm', 'Technogramm', 'technogramm@mail.ru', 'Служебный чат', 'Technogramm');
