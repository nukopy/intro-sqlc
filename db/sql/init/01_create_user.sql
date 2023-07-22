-- create role
create role "myuser" with password 'mypassword' login;

-- grant privileges on database
grant all privileges on database "intro_sqlc_db" to "myuser";

-- grant privileges on all tables in schema public
grant all privileges on all tables in schema public to "myuser";

-- 新しく作成されるテーブルに対する権限を設定
alter default privileges in schema public
grant all on tables to myuser;