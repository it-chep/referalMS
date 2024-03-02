-- +goose NO TRANSACTION
-- +goose Up
create index concurrently if not exists users_tg_id on users (tg_id);
-- create index concurrently if not exists users_tg_id on users (tg_id);
create index concurrently if not exists referals_tg_id on referals (tg_id);
-- create index concurrently if not exists referals_tg_id on referals (tg_id);
create index concurrently if not exists admins_login on admins (login);
-- create index concurrently if not exists admins_login on admin (login);

-- +goose Down
drop index concurrently users_tg_id;
-- drop index concurrently users_tg_id;
drop index concurrently referals_tg_id;
-- drop index concurrently referals_tg_id;
drop index concurrently admins_login;
-- drop index concurrently admins_login;
