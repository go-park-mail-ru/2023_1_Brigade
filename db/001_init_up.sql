DROP TABLE IF EXISTS user_id_seq CASCADE;
DROP TABLE IF EXISTS chat_id_seq CASCADE;
DROP TABLE IF EXISTS message_id_seq CASCADE;
DROP TABLE IF EXISTS profile CASCADE;
DROP TABLE IF EXISTS session CASCADE;
DROP TABLE IF EXISTS chat CASCADE;
DROP TABLE IF EXISTS message CASCADE;
DROP TABLE IF EXISTS chat_members CASCADE;
-- DROP TABLE IF EXISTS users_chats CASCADE;
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
    id     SERIAL UNIQUE PRIMARY KEY,
    type   INTEGER,
    title  VARCHAR(255),
    avatar VARCHAR(1024)
);

CREATE TABLE IF NOT EXISTS message (
    id        SERIAL UNIQUE PRIMARY KEY,
    body      VARCHAR(1024), -- валидация, со стороны приложения
    id_author INTEGER,
    id_chat   INTEGER,
--     is_read    BOOL,
--     created_at TIMESTAMP, -- google.com
    FOREIGN KEY (id_author) REFERENCES profile(id),
    FOREIGN KEY (id_chat)   REFERENCES chat(id)
);

CREATE TABLE IF NOT EXISTS users_chats (
    id_user INTEGER,
    id_chat   INTEGER,
    FOREIGN KEY (id_user)   REFERENCES profile(id),
    FOREIGN KEY (id_chat) REFERENCES   chat(id)
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

-- curl -X 'POST' 'http://localhost:8081/api/v1/chats/' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{ "title": "string", "members": [1,2,3]}'

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', '', 'marcussss1', 'marcussss1@gmail.com', 'marcussss1', '123');

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', '', 'marcussss2', 'marcussss2@gmail.com', 'marcussss2', '123');

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', '', 'marcussss3', 'marcussss3@gmail.com', 'marcussss3', '123');

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', '', 'marcussss4', 'marcussss4@gmail.com', 'marcussss4', '123');

-- INSERT INTO profile (avatar, username, nickname, email, status, password)
-- VALUES ('', '', 'marcussss5', 'marcussss5@gmail.com', 'marcussss5', '123');

INSERT INTO user_contacts (id_user, id_contact)
VALUES (
           (SELECT id FROM profile
            WHERE id = 1),
           (SELECT id FROM profile
            WHERE id = 2)
       );

INSERT INTO user_contacts (id_user, id_contact)
VALUES (
           (SELECT id FROM profile
            WHERE id = 1),
           (SELECT id FROM profile
            WHERE id = 3)
       );

INSERT INTO message (body, id_author, id_chat)
VALUES (   'HI',
           (SELECT id FROM profile
            WHERE id = 1),
           (SELECT id FROM chat
            WHERE id = 1)
       );

-- CREATE TABLE IF NOT EXISTS images_urls (
--     id_image SERIAL UNIQUE PRIMARY KEY,
--     image_url       VARCHAR(2048)
-- );
--
-- CREATE TABLE IF NOT EXISTS users_avatar (
--     id_user  INTEGER,
--     id_image INTEGER,
--     FOREIGN KEY (id_user)  REFERENCES profile(id),
--     FOREIGN KEY (id_image) REFERENCES images_urls(id_image)
-- );
--
-- INSERT INTO images_urls (id_image, image_url)
-- VALUES (
--            1,
--            'http://minio:9000/avatars/1avatara_ru_3D001.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=minio%2F20230403%2Fus-east-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20230403T123817Z\u0026X-Amz-Expires=604800\u0026X-Amz-SignedHeaders=host\u0026X-Amz-Signature=734c3e423869827332ab623cf3abebe6e8ba541ca9fc805b1acf2a836c934938'
--        );
--
-- INSERT INTO images_urls (id_image, image_url)
-- VALUES (
--            2,
--            'http://minio:9000/avatars/1avatara_ru_3D001.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=minio%2F20230403%2Fus-east-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20230403T123817Z\u0026X-Amz-Expires=604800\u0026X-Amz-SignedHeaders=host\u0026X-Amz-Signature=734c3e423869827332ab623cf3abebe6e8ba541ca9fc805b1acf2a836c934938'
--        );
--
-- INSERT INTO images_urls (id_image, image_url)
-- VALUES (
--            3,
--            'http://minio:9000/avatars/1avatara_ru_3D001.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=minio%2F20230403%2Fus-east-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20230403T123817Z\u0026X-Amz-Expires=604800\u0026X-Amz-SignedHeaders=host\u0026X-Amz-Signature=734c3e423869827332ab623cf3abebe6e8ba541ca9fc805b1acf2a836c934938'
--        );
--
-- INSERT INTO users_avatar (id_user, id_image)
-- VALUES (
--            1,
--            1
--         );
--
-- INSERT INTO users_avatar (id_user, id_image)
-- VALUES (
--            2,
--            2
--        );
--
-- INSERT INTO users_avatar (id_user, id_image)
-- VALUES (
--            3,
--            3
--        );
-- INSERT INTO user_contacts (id_user, id_contact)
-- VALUES (
--            (SELECT id FROM profile
--             WHERE id = 1),
--            (SELECT id FROM profile
--             WHERE id = 2)
--        );
--
-- INSERT INTO user_contacts (id_user, id_contact)
-- VALUES (
--            (SELECT id FROM profile
--             WHERE id = 1),
--            (SELECT id FROM profile
--             WHERE id = 3)
--        );
