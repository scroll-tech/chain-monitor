-- +goose Up
-- +goose L2ETHBegin
create table l2_eth_events
(
    number   bigint,
    tx_hash  varchar(66) not null
        primary key,
    msg_hash text,
    type     smallint,
    amount   numeric(32)
);

comment on column l2_eth_events.number is ' block number';

comment on column l2_eth_events.tx_hash is ' tx hash';

comment on column l2_eth_events.type is ' tx type';

alter table l2_eth_events
    owner to maskpp;

create index idx_l2_eth_events_type
    on l2_eth_events (type);

create index idx_l2_eth_events_msg_hash
    on l2_eth_events (msg_hash);

create index idx_l2_eth_events_number
    on l2_eth_events (number);
-- +goose L2ETHEnd

-- +goose Down
-- +goose L2ETHBegin
drop table if exists l2_eth_events
-- +goose L2ETHEnd
