-- таблица пользователей
-- 2 НФ, т.к. по email однозначно определяется username (и наоборот)
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
    cookie     TEXT PRIMARY KEY,
    profile_id INTEGER REFERENCES profile(id)
);

-- таблица чатов
-- 3 НФ
CREATE TABLE IF NOT EXISTS chat (
    id    SERIAL PRIMARY KEY,
    master_id    INTEGER REFERENCES profile(id), -- айди автора
    type INTEGER, -- тип чата (диалог/группа/канал)
    description TEXT,
    avatar TEXT, -- url аватарки
    title TEXT
);

-- таблица сообщений
-- 3 НФ
CREATE TABLE IF NOT EXISTS message (
    id         TEXT PRIMARY KEY,
    type       INTEGER, -- тип сообщения (обычное/стикер/картинка)
    body       TEXT,
    id_chat    INTEGER REFERENCES chat(id),
    author_id  INTEGER REFERENCES profile(id),
    created_at TEXT
);

-- таблица связи M:N чатов и пользователей (участников чата)
CREATE TABLE IF NOT EXISTS chat_members (
    id_chat   INTEGER REFERENCES chat(id),
    id_member INTEGER REFERENCES profile(id),
    -- используется в 5+ запросах (в т.ч. с AND), поэтому решили добавить индекс (primary key) здесь
    PRIMARY KEY (id_chat, id_member)
);

-- таблица связи M:N пользователей и пользователей в списке контактов
CREATE TABLE IF NOT EXISTS user_contacts (
    id_user    INTEGER REFERENCES profile(id),
    id_contact INTEGER REFERENCES profile(id)
);

-- таблица связи 1:M чатов и сообщений
CREATE TABLE IF NOT EXISTS chat_messages (
    id_chat    INTEGER REFERENCES chat(id),
    id_message TEXT REFERENCES message(id)
    -- в принципе, здесь можно было использовать PRIMARY KEY для составного ключа
    -- но мы решили так не делать, поскольку нет ни одного запроса, где используется AND
    -- поэтому мы вынесли в отдельные индексы id_chat и id_message
);

-- таблица ссылок на файлы, прикрепленные к сообщениям
CREATE TABLE IF NOT EXISTS attachments (
    id_message TEXT REFERENCES message(id),
    url        TEXT,
    name       TEXT
);

-- перед каждым индексом указаны ссылки на запросы, для которых они используются

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/auth/repository/repository.go#L46
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/auth/repository/repository.go#L60
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/user/repository/repository.go#L52
CREATE INDEX IF NOT EXISTS profile__email_username_idx
ON profile (email, username);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/auth/repository/repository.go#L74 
CREATE INDEX IF NOT EXISTS profile__username_idx
ON profile (username);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L107
CREATE INDEX IF NOT EXISTS chat_members__id_member_idx
ON chat_members (id_member);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L193
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L135
CREATE INDEX IF NOT EXISTS chat_messages__id_chat_idx
ON chat_messages (id_chat);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L64
CREATE INDEX IF NOT EXISTS chat_messages__id_message_idx
ON chat_messages (id_message);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L205
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L194
CREATE INDEX IF NOT EXISTS message__id_chat_created_at_idx
ON message (id_chat, created_at);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L258
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/chat/repository/repository.go#L276
CREATE INDEX IF NOT EXISTS chat)_title_idx 
ON chat USING GIN (title gin_trgm_ops);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L35
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L70
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L108
-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L205
CREATE INDEX IF NOT EXISTS attachments__id_message_idx
ON attachments (id_message);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/messages/repository/repository.go#L241
CREATE INDEX IF NOT EXISTS message__body_idx 
ON message USING GIN (body gin_trgm_ops);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/user/repository/repository.go#L72
CREATE INDEX IF NOT EXISTS user_contacts__id_user_idx
ON user_contacts (id_user);

-- https://github.com/go-park-mail-ru/2023_1_Brigade/blob/easyjson/internal/microservices/user/repository/repository.go#L205
CREATE INDEX IF NOT EXISTS profile__nickname_idx
ON profile USING GIN (nickname gin_trgm_ops);

INSERT INTO profile (id, avatar, username, nickname, email, status, password)
VALUES (0, 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm', 'Technogramm', 'technogramm@mail.ru', 'Служебный чат', 'Technogramm');
