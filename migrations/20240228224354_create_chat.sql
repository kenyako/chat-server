-- +goose Up
create table chats (
    id bigserial primary key,
    title text not null
);

create table users_chats (
    id bigserial primary key,
    name text not null,
    password text not null,
    email text not null
);

create table chat_messages (
    id bigserial primary key,
    chat_id bigint references chats(id),
    user_id bigint references users_chats(id),
    text_mes text not null,
    time_sent timestamp not null
);

-- +goose Down
drop table chats;
drop table users_chats;
drop table chat_messages;