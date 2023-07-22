create table authors (
    id integer primary key,
    name varchar(99) not null,
    created_at timestamp not null default now()
);

create table books (
    id integer primary key,
    title varchar(99) not null,
    price integer not null,
    author_id integer not null,
    created_at timestamp not null default now()
);

alter table books
add foreign key (author_id) references authors (id);