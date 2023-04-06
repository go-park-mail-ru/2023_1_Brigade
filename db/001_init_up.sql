-- CREATE TABLE user_id_seq;
-- CREATE TABLE chat_id_seq;
-- CREATE TABLE message_id_seq;
-- CREATE TABLE profile;
-- CREATE TABLE session;
-- CREATE TABLE chat;
-- CREATE TABLE message;
-- CREATE TABLE chat_members;
-- CREATE TABLE message_receiver;
-- CREATE TABLE user_contacts;
-- CREATE TABLE images_urls;
-- CREATE TABLE users_avatar;



-- curl -X 'POST' 'http://localhost:8081/api/v1/chats/' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{ "title": "string", "members": [1,2,3]}'

-- INSERT INTO profile (avatar, username, nickname, email, status, password)
-- VALUES ('', '', 'marcussss1', 'marcussss1@gmail.com', 'marcussss1', '123');
--
-- INSERT INTO profile (avatar, username, nickname, email, status, password)
-- VALUES ('', '', 'marcussss2', 'marcussss2@gmail.com', 'marcussss2', '123');
--
-- INSERT INTO profile (avatar, username, nickname, email, status, password)
-- VALUES ('', '', 'marcussss3', 'marcussss3@gmail.com', 'marcussss3', '123');

-- INSERT INTO profile (avatar, username, nickname, email, status, password)
-- VALUES ('', '', 'marcussss4', 'marcussss4@gmail.com', 'marcussss4', '123');
--
-- -- INSERT INTO profile (avatar, username, nickname, email, status, password)
-- -- VALUES ('', '', 'marcussss5', 'marcussss5@gmail.com', 'marcussss5', '123');
--
-- CREATE TABLE IF NOT EXISTS chat_members (
--                                             id_chat   INTEGER,
--                                             id_member INTEGER,
--                                             FOREIGN KEY (id_chat)   REFERENCES chat(id),
--     FOREIGN KEY (id_member) REFERENCES profile(id)
--     );

INSERT INTO chat_members (id_chat, id_member)
VALUES (
           (SELECT id FROM chat
            WHERE id = 1),
           (SELECT id FROM profile
            WHERE id = 2)
       );



--
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

INSERT INTO user_contacts (id_user, id_contact)
VALUES (
           (SELECT id FROM profile
            WHERE id = 1),
           (SELECT id FROM profile
            WHERE id = 4)
       );

INSERT INTO message (body, id_author, id_chat)
VALUES (   'HI',
           (SELECT id FROM profile
            WHERE id = 1),
           (SELECT id FROM chat
            WHERE id = 3)
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
