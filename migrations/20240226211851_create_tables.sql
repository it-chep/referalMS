-- +goose Up
-- +goose StatementBegin
create table if not exists integrators
(
    id   bigserial,
    name text
);

create table if not exists admins
(
    id                 bigserial,
    login              text unique,
    password           text,
    integrations_token text,
    salt               int,
    integrator_id      bigint,
    domain             text,
    active             bool,
    last_login_time    timestamp not null default current_timestamp,
    registration_time  timestamp not null default current_timestamp
);

create table if not exists referals
(
    id                        bigserial,
    admin_id                  bigint,
    tg_id                     bigint,
    id_in_integration_service bigint,
    name                      text,
    username                  text,
    referal_link              text,
    registration_time         timestamp not null default current_timestamp
);

create table if not exists users
(
    id                        bigserial,
    admin_id                  bigint    not null,
    tg_id                     bigint    not null,
    id_in_integration_service bigint,
    registration_time         timestamp not null default current_timestamp,
    referal_link              text,
    username                  text,
    referal_id                bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop table if exists referals;
drop table if exists admins;
drop table if exists integrators;

-- +goose StatementEnd
