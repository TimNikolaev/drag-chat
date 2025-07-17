CREATE TABLE users
(
  id serial primary key,
  name varchar(255) not null,
  username varchar(255) not null unique,
  email varchar(255) not null unique,
  password_hash varchar(255) not null
);

CREATE TABLE chats
(
  id serial primary key,
  chat_name varchar(255) not null unique
);

CREATE TABLE users_chats
(
  user_id int references users (id) on delete cascade not null,
  chat_id int references chats (id) on delete cascade not null
  primary key (user_id, chat_id)
);

CREATE TABLE messages
(
  id bigserial,
  chat_id int references chats (id) on delete cascade not null,
  sender_id int references users (id) on delete cascade not null,
  text_body text not null,
  is_edited boolean default false,
  send_time timestamp default now()
  primary key (id, chat_id)
);
