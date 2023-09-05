-- +goose Up
-- +goose L1ERC20Begin
create table l1_erc20_events
(
    number   bigint,
    tx_hash  varchar(66) not null
        primary key,
    msg_hash text,
    type     smallint,
    l1_token text,
    l2_token text,
    amount   numeric(32)
);

comment on column l1_erc20_events.number is ' block number';

comment on column l1_erc20_events.tx_hash is ' tx hash';

comment on column l1_erc20_events.type is ' tx type';

comment on column l1_erc20_events.l1_token is ' input token contract address';

comment on column l1_erc20_events.l2_token is ' output token contract address';

create index idx_l1_erc20_events_type
    on l1_erc20_events (type);

create index idx_l1_erc20_events_msg_hash
    on l1_erc20_events (msg_hash);

create index idx_l1_erc20_events_number
    on l1_erc20_events (number);
-- +goose L1ERC20End

-- +goose Down
-- +goose L1ERC20Begin
drop table if exists l1_erc20_events;
-- +goose L1ERC20End
