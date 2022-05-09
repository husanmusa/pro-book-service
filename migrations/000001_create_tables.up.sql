create table if not exists users
(
    id          uuid not null primary key,
    first_name  varchar,
    last_name   varchar,
    middle_name varchar,
    phone       varchar,
    email       varchar
);

create table if not exists books
(
    id             uuid not null primary key,
    name           varchar,
    count          integer,
    author         varchar,
    published_date date,
    created_at     timestamp,
    updated_at     timestamp,
    deleted_at     timestamp
);

create table if not exists user_books
(
    id      uuid not null primary key,
    book_id uuid references books (id),
    user_id uuid references users (id)
);