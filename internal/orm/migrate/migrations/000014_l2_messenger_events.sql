-- +goose Up
-- +goose L2MessengerBegin
create table l2_messenger_events
(
    number    bigint,
    type      smallint,
    msg_nonce bigserial
        primary key,
    msg_hash  text,
    msg_proof text,
    confirm   boolean
);

comment on column l2_messenger_events.number is ' block number';

comment on column l2_messenger_events.type is ' tx type';

alter table l2_messenger_events
    owner to maskpp;

create index idx_l2_messenger_events_number
    on l2_messenger_events (number);

create index idx_l2_messenger_events_confirm
    on l2_messenger_events (confirm);

create index idx_l2_messenger_events_msg_hash
    on l2_messenger_events (msg_hash);

create index idx_l2_messenger_events_type
    on l2_messenger_events (type);
-- +goose L2MessengerEnd

-- +goose Down
-- +goose L2MessengerBegin
drop table if exists l2_messenger_events
-- +goose L2MessengerEnd
