DROP TABLE IF EXISTS user_id_seq CASCADE;
DROP TABLE IF EXISTS chat_id_seq CASCADE;
DROP TABLE IF EXISTS message_id_seq CASCADE;
DROP TABLE IF EXISTS profile CASCADE;
DROP TABLE IF EXISTS session CASCADE;
DROP TABLE IF EXISTS chat CASCADE;
DROP TABLE IF EXISTS message CASCADE;
DROP TABLE IF EXISTS chat_members CASCADE;
DROP TABLE IF EXISTS message_receiver CASCADE;
DROP TABLE IF EXISTS user_contacts CASCADE;
DROP TABLE IF EXISTS images_urls CASCADE;
DROP TABLE IF EXISTS users_avatar CASCADE;

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
                                       id         SERIAL UNIQUE PRIMARY KEY,
                                       body       VARCHAR(1024), -- валидация, со стороны приложения
    id_chat    INTEGER,
    id_author  INTEGER,
    FOREIGN KEY (id_author) REFERENCES profile(id),
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
                                            id_message INTEGER,
                                            FOREIGN KEY (id_chat)   REFERENCES chat(id),
    FOREIGN KEY (id_message) REFERENCES message(id)
    );

