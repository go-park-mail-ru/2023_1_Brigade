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
