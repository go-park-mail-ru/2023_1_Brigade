CREATE TABLE IF NOT EXISTS profile (
    id       SERIAL UNIQUE,
    username VARCHAR(255),
    email    VARCHAR(255),
    status   VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS session (
    profile_id INTEGER,
    cookie  VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS chat (
    id    SERIAL UNIQUE,
    title VARCHAR(255),
);

CREATE TABLE IF NOT EXISTS message (
    id        SERIAL UNIQUE,
    body      VARCHAR(255),
    id_author INTEGER,
    id_chat   INTEGER,
    is_read   BIT,
    FOREIGN KEY (id_author) REFERENCES profile(id),
    FOREIGN KEY (id_chat)   REFERENCES chat(id)
);

CREATE TABLE IF NOT EXISTS chat_members (
    id_chat   INTEGER,
    id_member INTEGER,
    FOREIGN KEY (id_chat)   REFERENCES chat(id),
    FOREIGN KEY (id_member) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS message_receiver (
    id_message  INTEGER,
    id_receiver INTEGER,
    FOREIGN KEY (id_message)  REFERENCES message(id)
    FOREIGN KEY (id_receiver) REFERENCES profile(id)
);

-- CREATE TABLE IF NOT EXISTS reciever_messages (
--      id_recevier INTEGER,
--      id_message  INTEGER,
--      FOREIGN KEY (id_recevier) REFERENCES profile(id),
--      FOREIGN KEY (id_message)  REFERENCES message(id)
-- );

-- CREATE TABLE IF NOT EXISTS message (
--     id_receiver  UNIQUE,
--     id_sender  INTEGER,--     body       VARCHAR(255),
--     created_at VARCHAR(255),
--     FOREIGN KEY (id_sender)   REFERENCES profile(id),
-- )

-- Таблица "Messages" (Сообщения):
--
-- message_id (идентификатор сообщения, PRIMARY KEY)
-- chat_id (идентификатор чата, FOREIGN KEY)
-- user_id (идентификатор пользователя, FOREIGN KEY)
-- message_text (текст сообщения)
-- sent_at (дата и время отправки сообщения)
