CREATE DOMAIN email AS  VARCHAR(255) CHECK (VALUE ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$');

create table users (
    id SERIAL NOT NULL,
    username VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    email email UNIQUE,
    password_hash TEXT,
    description TEXT,
    image TEXT,
    status INTEGER DEFAULT 0,
    authorize_token_hash text,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id),
    check(LENGTH(username) > 2)
);