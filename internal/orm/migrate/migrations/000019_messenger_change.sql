-- +goose Up
-- +goose MessengerChangeBegin
alter table l1_messenger_events
    rename column confirm to from_gateway;
alter table l1_messenger_events
    alter column from_gateway set default true;

alter table l2_messenger_events
    rename column confirm to from_gateway;
alter table l2_messenger_events
    alter column from_gateway set default true;

alter table l2_messenger_events
    add tx_hash varchar(66) default '' not null;
-- +goose MessengerChangeEnd

-- +goose Down
-- +goose MessengerChangeBegin
alter table l1_messenger_events
    rename column from_gateway to confirm;
alter table l1_messenger_events
    alter column confirm drop default;

alter table l2_messenger_events
    rename column from_gateway to confirm;
alter table l2_messenger_events
    alter column confirm drop default;

alter table l2_messenger_events
drop column tx_hash;
-- +goose MessengerChangeEnd