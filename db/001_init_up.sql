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
    id         SERIAL UNIQUE,
    type       INTEGER,
    name       VARCHAR(255),
    created_at VARCHAR(255),
    masters    INTEGER REFERENCES profile(id)
);

CREATE TABLE IF NOT EXISTS chat_members (
    id_member INTEGER,
    id_chat INTEGER,
    FOREIGN KEY (id_chat)    REFERENCES chat(id),
    FOREIGN KEY (id_member) REFERENCES profile(id)
);
