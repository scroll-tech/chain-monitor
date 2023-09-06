-- +goose Up
-- +goose L2ERC721Begin
create table l2_erc721_events
(
    number   bigint,
    tx_hash  varchar(66) not null
        primary key,
    msg_hash text,
    type     smallint,
    l1_token text,
    l2_token text,
    token_id serial
);

comment on column l2_erc721_events.number is ' block number';

comment on column l2_erc721_events.tx_hash is ' tx hash';

comment on column l2_erc721_events.type is ' tx type';

create index idx_l2_erc721_events_msg_hash
    on l2_erc721_events (msg_hash);

create index idx_l2_erc721_events_number
    on l2_erc721_events (number);

create index idx_l2_erc721_events_type
    on l2_erc721_events (type);
-- +goose L2ERC721End

-- +goose Down
-- +goose L2ERC721Begin
drop table if exists l2_erc721_events;
-- +goose L2ERC721End
