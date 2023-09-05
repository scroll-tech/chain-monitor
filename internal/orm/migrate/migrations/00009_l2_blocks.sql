-- +goose Up
-- +goose L2BlocksBegin
create table l2_blocks
(
    number      bigserial
        primary key,
    event_count bigint
);

comment on column l2_blocks.number is ' block number';

alter table l2_blocks
    owner to maskpp;
-- +goose L2BlocksEnd

-- +goose Down
-- +goose L2BlocksBegin
drop table if exists l2_blocks
-- +goose L2BlocksEnd
