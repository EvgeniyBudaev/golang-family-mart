CREATE TABLE users
(
    id       uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(255) UNIQUE NOT null,
    password varchar(512)        NOT null,
    email    varchar(255)  NULL  DEFAULT ''
);