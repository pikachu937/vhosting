DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id            SERIAL       NOT NULL UNIQUE,
    username      VARCHAR(30) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);
