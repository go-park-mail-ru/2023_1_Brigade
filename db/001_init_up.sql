CREATE TABLE IF NOT EXISTS profile (
    id       SERIAL UNIQUE PRIMARY KEY,
    avatar   VARCHAR(1024),
    username VARCHAR(255),
    nickname VARCHAR(255),
    email    VARCHAR(255),
    status   VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS chat (
    id    SERIAL UNIQUE PRIMARY KEY,
    type INTEGER,
    avatar VARCHAR(255),
    title VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS message (
    id         VARCHAR(255) UNIQUE,
    body       VARCHAR(1024), -- валидация, со стороны приложения
    id_chat    INTEGER,
    author_id  INTEGER,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (author_id) REFERENCES profile(id),
    FOREIGN KEY (id_chat)   REFERENCES chat(id)
);

CREATE TABLE IF NOT EXISTS chat_members (
    id_chat   INTEGER,
    id_member INTEGER,
    FOREIGN KEY (id_chat)   REFERENCES chat(id),
    FOREIGN KEY (id_member) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS user_contacts (
    id_user    INTEGER,
    id_contact INTEGER,
    FOREIGN KEY (id_user)    REFERENCES profile(id),
    FOREIGN KEY (id_contact) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS chat_messages (
    id_chat    INTEGER,
    id_message VARCHAR(255),
    FOREIGN KEY (id_chat)   REFERENCES chat(id),
    FOREIGN KEY (id_message) REFERENCES message(id)
);
