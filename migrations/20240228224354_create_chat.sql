-- +goose Up
create table chats (
    id bigserial primary key,
    title text not null
);

create table users_chats (
    chat_id bigserial references chats(id),
    user_id bigint,

    primary key (chat_id, user_id)
);

create table chat_messages (
    id bigserial primary key,
    chat_id bigint references chats(id),
    user_id bigint,
    text text not null,
    time_sent timestamp not null
);

-- +goose Down
drop table chats;
drop table users_chats;
drop table chat_messages;