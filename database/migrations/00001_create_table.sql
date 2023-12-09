-- +goose Up
create table public."tasks"
(
    id      serial not null,
    status  integer not null,
    message text
);

create unique index "tasks_id_uindex"
    on public."tasks" (id);

create index "tasks_status_index"
    on public."tasks" (status);

alter table public."tasks"
    add constraint "tasks_pk"
        primary key (id);

-- +goose Down
DROP TABLE IF EXISTS public."tasks" CASCADE;