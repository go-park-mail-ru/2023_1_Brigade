CREATE TABLE IF NOT EXISTS profile (
    id       SERIAL UNIQUE PRIMARY KEY,
    username VARCHAR(255),
    email    VARCHAR(255),
    status   VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS chat (
    id    SERIAL UNIQUE PRIMARY KEY,
    title VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS message (
    id         SERIAL UNIQUE PRIMARY KEY,
    body       VARCHAR(1024), -- валидация, со стороны приложения
    id_author  INTEGER,
    id_chat    INTEGER,
    is_read    BOOL,
    created_at TIMESTAMP, -- google.com
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
    FOREIGN KEY (id_message)  REFERENCES message(id),
    FOREIGN KEY (id_receiver) REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS user_contacts (
    id_user    INTEGER,
    id_contact INTEGER,
    FOREIGN KEY (id_user)    REFERENCES profile(id),
    FOREIGN KEY (id_contact) REFERENCES profile(id)
);
