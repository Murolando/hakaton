
CREATE TABLE "role"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(25) not null
);
INSERT INTO "role" (name) values('child'); 
INSERT INTO "role" (name) values('teacher'); 
INSERT INTO "role" (name) values('admin'); 
CREATE TABLE "user"
(
    id                    serial PRIMARY KEY not null unique,
    name                  VARCHAR(64) not null,
    login                 VARCHAR(40) unique not null,
    image_src             varchar(500),
    password_hash         varchar(255) not null,
    role_id               int references "role"(id) on delete cascade,
    refresh               VARCHAR(64),
    expiredAt             TIMESTAMP(0) 
);

CREATE TABLE "class"
(
    id                    serial PRIMARY KEY not null unique,
    name                  VARCHAR(64) not null,
    description           VARCHAR(500),
    code                  VARCHAR(10) not null,
    created_at            TIMESTAMP(0) NOT NULL,
    comments_access       bool
);

CREATE TABLE "user_class"
(
    id                    serial PRIMARY KEY not null unique,
    user_id               int references "user"(id) on delete cascade,
    class_id              int references "class"(id) on delete cascade
);
CREATE TABLE "lesson_type"
(
    id                    serial PRIMARY KEY not null unique,
    name                  VARCHAR(100) not null
);
INSERT INTO "lesson_type" (name) values('theory');
INSERT INTO "lesson_type" (name) values('kontur');
CREATE TABLE "lesson"
(
    id                    serial PRIMARY KEY not null unique,
    name                  VARCHAR(100) not null,
    video                 VARCHAR(200),
    lesson_type_id        int references "lesson_type"(id) on delete cascade,
    class_id              int references "class"(id) on delete cascade,
    created_at            TIMESTAMP(0) NOT NULL,
    expired_at            TIMESTAMP(0) NOT NULL,
    lesson_access         bool
);
CREATE TABLE "user_kontur_result"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_id             int references "lesson"(id) on delete cascade,
    user_id               int references "user"(id) on delete cascade,
    max_count             int,
    count                 int,
    result                int
);
CREATE TABLE "comment"
(
    id                   serial PRIMARY KEY not null unique,
    name                 VARCHAR(100) not null,
    lesson_id            int references "lesson"(id) on delete cascade,
    author_id            int references "user"(id) on delete cascade,
    created_at           TIMESTAMP(0) NOT NULL
);
CREATE TABLE "kontur"
(
    id                   serial PRIMARY KEY not null unique,
    name                 VARCHAR(100) not null,
    image_src            VARCHAR(400) not null
);
