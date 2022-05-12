DROP TABLE IF EXISTS public.videos;
DROP TABLE IF EXISTS public.infos;
DROP TABLE IF EXISTS public.user_groups;
DROP TABLE IF EXISTS public.group_perms;
DROP TABLE IF EXISTS public.logs;
DROP TABLE IF EXISTS public.sessions;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.groups;
DROP TABLE IF EXISTS public.perms;


CREATE TABLE IF NOT EXISTS public.perms (
    id        INTEGER      NOT NULL UNIQUE,
    name      VARCHAR(100) NOT NULL UNIQUE,
    code_name VARCHAR(30)  NOT NULL UNIQUE,
	CONSTRAINT pk_perms PRIMARY KEY (id)
);
INSERT INTO public.perms (id, name, code_name) VALUES
(0, 'Can create User',           'post_user'),
(1, 'Can get User info',         'get_user'),
(2, 'Can get All-Users info',    'get_all_users'),
(3, 'Can partially update User', 'patch_user'),
(4, 'Can delete User',           'delete_user');


CREATE TABLE IF NOT EXISTS public.groups (
    id   INTEGER     NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL UNIQUE,
	CONSTRAINT pk_groups PRIMARY KEY (id)
);
INSERT INTO public.groups (id, name) VALUES
(0, 'Super Users'),
(1, 'Staff'),
(2, 'Users');


CREATE TABLE IF NOT EXISTS public.users (
    id            SERIAL                   NOT NULL UNIQUE,
    username      VARCHAR(50)              NOT NULL UNIQUE,
    password_hash VARCHAR(255)             NOT NULL,
    is_active     BOOLEAN                  NOT NULL,
    is_superuser  BOOLEAN                  NOT NULL,
    is_staff      BOOLEAN                  NOT NULL,
    first_name    VARCHAR(50)              NOT NULL,
    last_name     VARCHAR(50)              NOT NULL,
    joining_date  TIMESTAMP WITH TIME ZONE NOT NULL,
    last_login    TIMESTAMP WITH TIME ZONE NOT NULL,
	CONSTRAINT pk_users PRIMARY KEY (id)
);
INSERT INTO public.users (id, username, password_hash, is_active, is_superuser, is_staff, first_name, last_name, joining_date, last_login) VALUES
(0, 'admin', '614240232425318c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', True, True, False, '', '', '2022-05-11 09:32:41.115644+03', '2022-05-11 09:32:41.115644+03');
ALTER SEQUENCE users_id_seq RESTART WITH 1;


CREATE TABLE IF NOT EXISTS public.sessions (
    id            SERIAL                   NOT NULL UNIQUE,
    content       VARCHAR(512)             NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT pk_sessions PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS public.logs (
    id             SERIAL                   NOT NULL UNIQUE,
    error_level    VARCHAR(7),                              -- "info", "warning", "error", "fatal"
    session_owner  VARCHAR(50),
    request_method VARCHAR(7),                              -- "POST", "GET", "PATCH", "DELETE"
    request_path   VARCHAR(100),
    status_code    INTEGER,
    error_code     INTEGER,
    message        VARCHAR(300),
    creation_date  TIMESTAMP WITH TIME ZONE
);


CREATE TABLE IF NOT EXISTS public.group_perms (
    id       SERIAL  NOT NULL UNIQUE,
    group_id INTEGER NOT NULL,
    perm_id  INTEGER NOT NULL,
	CONSTRAINT pk_group_perms PRIMARY KEY (id),
	CONSTRAINT fk_group_perms_groups FOREIGN KEY (group_id)
		REFERENCES public.groups (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_group_perms_perms FOREIGN KEY (perm_id)
		REFERENCES public.perms (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);
INSERT INTO public.group_perms (id, group_id, perm_id) VALUES
(0, 0, 0),
(1, 0, 1),
(2, 0, 2),
(3, 0, 3),
(4, 0, 4);
ALTER SEQUENCE group_perms_id_seq RESTART WITH 5;


CREATE TABLE IF NOT EXISTS public.user_groups (
    id       SERIAL  NOT NULL UNIQUE,
    user_id  INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
	CONSTRAINT pk_user_groups PRIMARY KEY (id),
	CONSTRAINT fk_user_groups_users FOREIGN KEY (user_id)
		REFERENCES public.users (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_user_groups_groups FOREIGN KEY (group_id)
		REFERENCES public.groups (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);
INSERT INTO public.user_groups (id, user_id, group_id) VALUES
(0, 0, 0);
ALTER SEQUENCE user_groups_id_seq RESTART WITH 1;


CREATE TABLE IF NOT EXISTS public.infos (
    id            SERIAL                   NOT NULL UNIQUE,
    creation_date TIMESTAMP WITH TIME ZONE NOT NULL,
    stream        TEXT                     NOT NULL,
    start_period  TIMESTAMP WITH TIME ZONE NOT NULL,
    stop_period   TIMESTAMP WITH TIME ZONE NOT NULL,
    time_life     TIMESTAMP WITH TIME ZONE NOT NULL,
    user_id       INTEGER                  NOT NULL,
    CONSTRAINT pk_infos PRIMARY KEY (id),
    CONSTRAINT fk_infos_users FOREIGN KEY (user_id)
		REFERENCES public.users (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS public.videos (
    id            SERIAL                   NOT NULL UNIQUE,
    url           VARCHAR(1024)            NOT NULL,
    file_name     VARCHAR(260)             NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE NOT NULL,
    info_id       INTEGER                  NOT NULL,
    user_id       INTEGER                  NOT NULL,
    CONSTRAINT pk_videos PRIMARY KEY (id),
    CONSTRAINT fk_videos_infos FOREIGN KEY (info_id)
		REFERENCES public.infos (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
    CONSTRAINT fk_videos_users FOREIGN KEY (user_id)
		REFERENCES public.users (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);
