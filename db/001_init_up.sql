DROP TABLE IF EXISTS user_id_seq CASCADE;
DROP TABLE IF EXISTS chat_id_seq CASCADE;
DROP TABLE IF EXISTS message_id_seq CASCADE;
DROP TABLE IF EXISTS profile CASCADE;
DROP TABLE IF EXISTS chat CASCADE;
DROP TABLE IF EXISTS message CASCADE;
DROP TABLE IF EXISTS chat_members CASCADE;
DROP TABLE IF EXISTS user_contacts CASCADE;
DROP TABLE IF EXISTS chat_messages CASCADE;


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

CREATE TABLE IF NOT EXISTS chat_members (
    id_chat    INTEGER,
    id_message INTEGER,
    FOREIGN KEY (id_chat)   REFERENCES chat(id),
    FOREIGN KEY (id_member) REFERENCES message(id)
);

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', 'marcussss1', 'marcussss1', 'marcussss1@gmail.com', 'marcussss1', '123');

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', 'marcussss2', 'marcussss2', 'marcussss2@gmail.com', 'marcussss2', '123');

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', 'marcussss3', 'marcussss3', 'marcussss3@gmail.com', 'marcussss3', '123');

INSERT INTO chat (title)
VALUES ('cool');

INSERT INTO chat_members (id_chat, id_member)
VALUES (
           (SELECT id FROM chat
            WHERE id = 1),
           (SELECT id FROM profile
            WHERE id = 1)
       );

INSERT INTO chat_members (id_chat, id_member)
VALUES (
           (SELECT id FROM chat
            WHERE id = 1),
           (SELECT id FROM profile
            WHERE id = 2)
       );

INSERT INTO chat_members (id_chat, id_member)
VALUES (
           (SELECT id FROM chat
            WHERE id = 1),
           (SELECT id FROM profile
            WHERE id = 3)
       );
