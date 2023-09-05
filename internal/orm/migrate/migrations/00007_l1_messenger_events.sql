-- +goose Up
-- +goose L1MessengerBegin
create table l1_messenger_events
(
    number   bigint,
    tx_hash  varchar(66) not null
        primary key,
    msg_hash text,
    type     smallint,
    confirm  boolean
);

comment on column l1_messenger_events.number is ' block number';

comment on column l1_messenger_events.tx_hash is ' tx hash';

comment on column l1_messenger_events.type is ' tx type';

alter table l1_messenger_events
    owner to maskpp;

create index idx_l1_messenger_events_msg_hash
    on l1_messenger_events (msg_hash);

create index idx_l1_messenger_events_number
    on l1_messenger_events (number);

create index idx_l1_messenger_events_confirm
    on l1_messenger_events (confirm);

create index idx_l1_messenger_events_type
    on l1_messenger_events (type);
-- +goose L1MessengerEnd

-- +goose Down
-- +goose L1MessengerBegin
drop table if exists l1_messenger_events
-- +goose L1MessengerEnd
