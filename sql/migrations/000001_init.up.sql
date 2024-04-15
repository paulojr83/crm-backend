create table customers(
  id            varchar(255) not null primary key,
  name          varchar(255)        not null,
  role          varchar(16)        not null,
  email         varchar(32)        not null,
  phone         varchar(32),
  contacted     bool
)