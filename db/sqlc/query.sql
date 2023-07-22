-- name: GetAuthor :one
select *
from authors
where id = $1
limit 1;

-- name: ListAuthors :many
select *
from authors
order by id;

-- name: CreateAuthor :one
insert into authors (id, name)
values ($1, $2)
returning *;

-- name: DeleteAuthor :exec
delete from authors
where id = $1;

-- name: ListBookOverPrice :many
select b.title,
    a.name,
    b.price
from books b
    left join authors a on 1 = 1
    and b.author_id = a.id
where price > $1
order by b.title;