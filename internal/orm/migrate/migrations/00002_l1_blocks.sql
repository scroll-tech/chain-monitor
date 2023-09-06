-- +goose Up
-- +goose L1BlocksBegin
create table l1_blocks
(
    number      bigserial
        primary key,
    event_count bigint
);
-- +goose L1BlocksEnd

-- +goose Down
-- +goose L1BlocksBegin
drop table if exists l1_blocks;
-- +goose L1BlocksEnd
