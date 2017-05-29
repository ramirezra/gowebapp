drop table posts cascade if exists;
drop table comments if exists;

create table posts (
  id serial primary key,
  content text,
  author varchar(255)
);

create table comments (
  id serial primary key,
  content text,
  author varchar(255),
  post_id integer references posts(id)
);

GRANT USAGE, SELECT ON SEQUENCE posts_id_seq TO gwp;
GRANT USAGE, SELECT ON SEQUENCE comments_id_seq TO gwp;
