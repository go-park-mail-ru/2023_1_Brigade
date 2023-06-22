CREATE USER brigade WITH PASSWORD 'brigade';
CREATE USER brigade_chats WITH PASSWORD 'brigade';

GRANT SELECT, UPDATE, INSERT, DELETE ON
    attachments,
    chat,
    chat_id_seq,
    chat_members,
    chat_messages,
    message,
    profile,
    profile_id_seq,
    session,
    user_contacts
    TO brigade;

GRANT SELECT ON
    attachments,
    chat,
    chat_messages,
    message
    TO brigade_chats;
