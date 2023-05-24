-- CREATE TABLE IF NOT EXISTS profile (
--     id       SERIAL UNIQUE PRIMARY KEY,
--     avatar   VARCHAR(1024),
--     username VARCHAR(255),
--     nickname VARCHAR(255),
--     email    VARCHAR(255),
--     status   VARCHAR(255),
--     password VARCHAR(255)
-- );
--
-- CREATE TABLE IF NOT EXISTS session (
--     cookie     VARCHAR(255),
--     profile_id INTEGER,
--     FOREIGN KEY (profile_id) REFERENCES profile(id)
-- );
--
-- CREATE TABLE IF NOT EXISTS chat (
--     id    SERIAL UNIQUE PRIMARY KEY,
--     master_id    INTEGER,
--     type INTEGER,
--     avatar VARCHAR(1024),
--     title VARCHAR(255),
--     FOREIGN KEY (master_id) REFERENCES profile(id)
-- );
--
-- CREATE TABLE IF NOT EXISTS message (
--     id         VARCHAR(255) UNIQUE,
--     image_url  VARCHAR(1024),
--     type       INTEGER,
--     body       VARCHAR(1024), -- валидация, со стороны приложения
--     id_chat    INTEGER,
--     author_id  INTEGER,
--     created_at VARCHAR(255),
--     FOREIGN KEY (author_id) REFERENCES profile(id),
--     FOREIGN KEY (id_chat)   REFERENCES chat(id)
-- );
--
-- CREATE TABLE IF NOT EXISTS chat_members (
--     id_chat   INTEGER,
--     id_member INTEGER,
--     FOREIGN KEY (id_chat)   REFERENCES chat(id),
--     FOREIGN KEY (id_member) REFERENCES profile(id)
-- );
--
-- CREATE TABLE IF NOT EXISTS user_contacts (
--     id_user    INTEGER,
--     id_contact INTEGER,
--     FOREIGN KEY (id_user)    REFERENCES profile(id),
--     FOREIGN KEY (id_contact) REFERENCES profile(id)
-- );
--
-- CREATE TABLE IF NOT EXISTS chat_messages (
--     id_chat    INTEGER,
--     id_message VARCHAR(255),
--     FOREIGN KEY (id_chat)   REFERENCES chat(id),
--     FOREIGN KEY (id_message) REFERENCES message(id)
-- );
--
-- INSERT INTO profile (id, avatar, username, nickname, email, status, password)
-- VALUES (0, 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm', 'Technogramm', '', 'Technogramm', '123');



CREATE TABLE IF NOT EXISTS profile (
                                       id       SERIAL UNIQUE PRIMARY KEY,
                                       avatar   VARCHAR(1024),
    username VARCHAR(255),
    nickname VARCHAR(255),
    email    VARCHAR(255),
    status   VARCHAR(255),
    password VARCHAR(255)
    );

CREATE TABLE IF NOT EXISTS session (
    cookie     VARCHAR(255),
    profile_id INTEGER,
    FOREIGN KEY (profile_id) REFERENCES profile(id)
    );

CREATE TABLE IF NOT EXISTS chat (
                                    id    SERIAL UNIQUE PRIMARY KEY,
                                    master_id    INTEGER,
                                    type INTEGER,
                                    avatar VARCHAR(1024),
    title VARCHAR(255),
    FOREIGN KEY (master_id) REFERENCES profile(id)
    );

CREATE TABLE IF NOT EXISTS message (
    id         VARCHAR(255) UNIQUE,
    type       INTEGER,
    body       VARCHAR(1024), -- валидация, со стороны приложения
    id_chat    INTEGER,
    author_id  INTEGER,
    created_at VARCHAR(255),
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

CREATE TABLE IF NOT EXISTS attachments (
                                           id_message VARCHAR(255),
                                           url        VARCHAR(1024),
    name       VARCHAR(255),
    FOREIGN KEY (id_message) REFERENCES message(id)
    );

INSERT INTO profile (id, avatar, username, nickname, email, status, password)
VALUES (0, 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm', 'Technogramm', '', 'Technogramm', '123');
