CREATE TABLE Profile (
    id       serial,
    username varchar(255),
    name     varchar(255),
    email    varchar(255),
    status   varchar(255),
    password varchar(255)
);

CREATE TABLE Session (
    user_id integer,
    cookie  varchar(255)
);
