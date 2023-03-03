DROP TABLE profile;
DROP SEQUENCE profileSeq;

CREATE SEQUENCE profileSeq
   START 1
   INCREMENT 1;

CREATE TABLE Profile (
	id       integer primary key not null DEFAULT nextval('profileSeq'),
	username    varchar(255),
	name varchar(255),
	email    varchar(255),
	status   varchar(255),
	password varchar(255)
);

CREATE TABLE Session (
    user_id   integer,
    cookie    varchar(255),
);
