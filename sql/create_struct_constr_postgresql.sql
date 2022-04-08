DROP TABLE IF EXISTS user_groups;
DROP TABLE IF EXISTS group_perms;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS perms;


CREATE TABLE IF NOT EXISTS perms (
    id        SERIAL       NOT NULL UNIQUE,
    name      VARCHAR(100) NOT NULL UNIQUE,
    code_name VARCHAR(30)  NOT NULL UNIQUE,
    PRIMARY KEY (id)
);
INSERT INTO perms (id, name, code_name) VALUES
(0, 'Can create User',           'post_user'),
(1, 'Can get User info',         'get_user'),
(2, 'Can get All-Users info',    'get_all_users'),
(3, 'Can partially update User', 'patch_user'),
(4, 'Can delete User',           'delete_user');
ALTER SEQUENCE perms_id_seq RESTART WITH 5;


CREATE TABLE IF NOT EXISTS groups (
    id   SERIAL      NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);
INSERT INTO groups (id, name) VALUES
(0, 'Super Users'),
(1, 'Staff'),
(2, 'Users');
ALTER SEQUENCE groups_id_seq RESTART WITH 3;


CREATE TABLE IF NOT EXISTS users (
    id            SERIAL                   NOT NULL UNIQUE,
    username      VARCHAR(50)              NOT NULL UNIQUE,
    password_hash VARCHAR(255)             NOT NULL,
    is_active     BOOLEAN                  NOT NULL,
    is_superuser  BOOLEAN                  NOT NULL,
    is_staff      BOOLEAN                  NOT NULL,
    first_name    VARCHAR(50)              NOT NULL,
    last_name     VARCHAR(50)              NOT NULL,
    email         VARCHAR(255)             NOT NULL,
    date_joined   TIMESTAMP WITH TIME ZONE NOT NULL,
    last_login    TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (id)
);
INSERT INTO users (id, username, password_hash, is_active, is_superuser, is_staff, first_name, last_name, email, date_joined, last_login) VALUES
(0, 'superuser', '6a4b40733133447655336f33482365304e376a39474068394b377223507320f3765880a5c269b747e1e906054a4b4a3a991259f1e16b5dde4742cec2319a', 'true',  'true',  'false', '-', '-', 'superuser@superuser.com', '1970-01-01 03:00:00.000000+03', '1970-01-01 03:00:00.000000+03'),
(1, 'ivanov',    '6a4b40733133447655336f33482365304e376a39474068394b37722350735994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', 'true',  'true',  'false', 'Ivan', 'Ivanov', 'i.ivanov@google.com', '2022-04-07 12:55:10.017647+03', '2022-04-07 12:55:10.017647+03'),
(2, 'petrov',    '6a4b40733133447655336f33482365304e376a39474068394b37722350735994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', 'true',  'false', 'true',  'Petr', 'Petrov', 'p.petrov@yandex.ru', '2022-04-07 12:56:10.017647+03', '2022-04-07 12:56:10.017647+03'),
(3, 'sidorov',   '6a4b40733133447655336f33482365304e376a39474068394b37722350735994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', 'true',  'false', 'false', 'Sidor', 'Sidorov', 's.sidorov@mail.ru', '2022-04-07 12:57:10.017647+03', '2022-04-07 12:57:10.017647+03');
ALTER SEQUENCE users_id_seq RESTART WITH 4;


CREATE TABLE IF NOT EXISTS group_perms (
    id       SERIAL  NOT NULL UNIQUE,
    group_id INTEGER NOT NULL,
    perm_id  INTEGER NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_group_id
        FOREIGN KEY (group_id)
        REFERENCES groups (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_perm_id
        FOREIGN KEY (perm_id)
        REFERENCES perms (id)
        ON DELETE CASCADE
);
INSERT INTO group_perms (id, group_id, perm_id) VALUES
(0, 0, 0),
(1, 0, 1),
(2, 0, 2),
(3, 0, 3),
(4, 0, 4);
ALTER SEQUENCE group_perms_id_seq RESTART WITH 5;


CREATE TABLE IF NOT EXISTS user_groups (
    id       SERIAL  NOT NULL UNIQUE,
    user_id  INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_group_id
        FOREIGN KEY (group_id)
        REFERENCES groups (id)
        ON DELETE CASCADE
);
INSERT INTO user_groups (id, user_id, group_id) VALUES
(0, 0, 0),
(1, 1, 0),
(2, 2, 1),
(3, 3, 2);
ALTER SEQUENCE user_groups_id_seq RESTART WITH 4;
