-- +goose Up
-- +goose L2ERC1155Begin
create table l2_erc1155_events
(
    number   bigint,
    tx_hash  varchar(66) not null
        primary key,
    msg_hash text,
    type     smallint,
    l1_token text,
    l2_token text,
    token_id serial,
    amount   numeric(32)
);

comment on column l2_erc1155_events.number is ' block number';

comment on column l2_erc1155_events.tx_hash is ' tx hash';

comment on column l2_erc1155_events.type is ' tx type';

create index idx_l2_erc1155_events_type
    on l2_erc1155_events (type);

create index idx_l2_erc1155_events_msg_hash
    on l2_erc1155_events (msg_hash);

create index idx_l2_erc1155_events_number
    on l2_erc1155_events (number);
-- +goose L2ERC1155End

-- +goose Down
-- +goose L2ERC1155Begin
drop table if exists l2_erc1155_events;
-- +goose L2ERC1155End
