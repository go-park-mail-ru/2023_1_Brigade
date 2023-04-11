INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', 'marcussss1', 'marcussss1', 'marcussss1@gmail.com', 'marcussss1', '123');

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', 'marcussss2', 'marcussss2', 'marcussss2@gmail.com', 'marcussss2', '123');

INSERT INTO profile (avatar, username, nickname, email, status, password)
VALUES ('', 'marcussss3', 'marcussss3', 'marcussss3@gmail.com', 'marcussss3', '123');

-- { "body": "string", "author_id": 1, "chat_id": 1 }
-- { "body": "string", "author_id": 1, "chat_id": 1 }
--   // {"message":{"id":0,"body":"hello world!","author_id":0,"chat_id":0,"is_read":false},"receiver_id":0}

INSERT INTO chat (type, avatar, title)
VALUES (1, 'avatarka', 'title');

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


INSERT INTO user_contacts (id_user, id_contact)
VALUES (
           (SELECT id FROM chat
            WHERE id = 5),
           (SELECT id FROM profile
            WHERE id = 1)
       );

INSERT INTO profile (id, username, email, status, password)
VALUES (1, 'marcussss1', 'marcussss1@gmail.com', 'marcussss1', '123');

INSERT INTO profile (id, username, email, status, password)
VALUES (2, 'marcussss2', 'marcussss2@gmail.com', 'marcussss2', '123');

INSERT INTO profile (id, username, email, status, password)
VALUES (3, 'marcussss3', 'marcussss3@gmail.com', 'marcussss3', '123');

INSERT INTO chat (id, title)
VALUES (1, 'my_chat');

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
