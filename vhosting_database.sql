DROP TABLE IF EXISTS public.videos;
DROP TABLE IF EXISTS public.infos;
DROP TABLE IF EXISTS public.user_groups;
DROP TABLE IF EXISTS public.user_perms;
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
( 0, 'Can create a User',                'post_user'),
( 1, 'Can get a User',                   'get_user'),
( 2, 'Can get all of the Users',         'get_all_users'),
( 3, 'Can partially update a User',      'patch_user'),
( 4, 'Can delete a User',                'delete_user'),

( 5, 'Can create a Group',               'post_group'),
( 6, 'Can get a Group',                  'get_group'),
( 7, 'Can get all of the Groups',        'get_all_groups'),
( 8, 'Can partially update a Group',     'patch_group'),
( 9, 'Can delete a Group',               'delete_group'),

(10, 'Can get all of the Permissions',   'get_all_perms'),
(11, 'Can set the user Permissions',     'set_user_perms'),
(12, 'Can get the user Permissions',     'get_user_perms'),
(13, 'Can delete the user Permissions',  'delete_user_perms'),
(14, 'Can set the group Permissions',    'set_group_perms'),
(15, 'Can get the group Permissions',    'get_group_perms'),
(16, 'Can delete the group Permissions', 'delete_group_perms'),

(17, 'Can create a Video',               'post_video'),
(18, 'Can get a Video',                  'get_video'),
(19, 'Can get all of the Videos',        'get_all_videos'),
(20, 'Can partially update a Video',     'patch_video'),
(21, 'Can delete a Video',               'delete_video'),

(22, 'Can create an Info',               'post_info'),
(23, 'Can get an Info',                  'get_info'),
(24, 'Can get all of the Infos',         'get_all_infos'),
(25, 'Can partially update an Info',     'patch_info'),
(26, 'Can delete an Info',               'delete_info');


CREATE TABLE IF NOT EXISTS public.groups (
    id   SERIAL      NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL UNIQUE,
	CONSTRAINT pk_groups PRIMARY KEY (id)
);
INSERT INTO public.groups (id, name) VALUES
(0, 'Video Creators'),
(1, 'Info Creators');
ALTER SEQUENCE groups_id_seq RESTART WITH 2;


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
    UNIQUE (group_id, perm_id),
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
( 0, 0, 17),
( 1, 0, 18),
( 2, 0, 19),
( 3, 0, 20),
( 4, 0, 21),

( 5, 1, 22),
( 6, 1, 23),
( 7, 1, 24),
( 8, 1, 25),
( 9, 1, 26);
ALTER SEQUENCE group_perms_id_seq RESTART WITH 10;


CREATE TABLE IF NOT EXISTS public.user_perms (
    id      SERIAL  NOT NULL UNIQUE,
    user_id INTEGER NOT NULL,
    perm_id INTEGER NOT NULL,
    UNIQUE (user_id, perm_id),
	CONSTRAINT pk_user_perms PRIMARY KEY (id),
	CONSTRAINT fk_user_perms_users FOREIGN KEY (user_id)
		REFERENCES public.users (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	CONSTRAINT fk_user_perms_perms FOREIGN KEY (perm_id)
		REFERENCES public.perms (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);


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


CREATE TABLE IF NOT EXISTS public.infos (
    id            SERIAL                   NOT NULL UNIQUE,
    stream        TEXT                     NOT NULL,
    start_period  TIMESTAMP WITH TIME ZONE NOT NULL,
    stop_period   TIMESTAMP WITH TIME ZONE NOT NULL,
    life_time     TIMESTAMP WITH TIME ZONE NOT NULL,
    user_id       INTEGER                  NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE NOT NULL,
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
    info_id       INTEGER                  NOT NULL,
    user_id       INTEGER                  NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE NOT NULL,
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
