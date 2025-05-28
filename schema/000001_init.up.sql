CREATE TABLE users
(
  id serial not null unique,
  name varchar(255) not null,
  username varchar(255) not null unique,
  email varchar(255) not null unique,
  password_hash varchar(255) not null unique
);

CREATE TABLE chats
(
  id serial not null unique,
  chat_name varchar(255) not null unique
);

CREATE TABLE users_chats
(
  id serial not null unique,
  user_id int references users (id) on delete cascade not null,
  chat_id int references chats (id) on delete cascade not null
);

CREATE TABLE messages
(
  id serial not null unique,
  user_id int references users (id) on delete cascade not null,
  chat_id int references chats (id) on delete cascade not null,
  text_body text not null,
  send_time timestamp default null
);
