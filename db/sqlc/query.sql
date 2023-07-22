-- name: GetAuthor :one
select *
from author
where id = $1
limit 1;

-- name: ListAuthors :many
select *
from author
order by id;

-- name: CreateAuthor :one
insert into author (id, name)
values ($1, $2)
returning *;

-- name: DeleteAuthor :exec
delete from author
where id = $1;

-- name: ListBookOverPrice :many
select b.title,
    a.name,
    b.price
from book b
    left join author a on 1 = 1
    and b.author_id = a.id
where price > $1
order by b.title;