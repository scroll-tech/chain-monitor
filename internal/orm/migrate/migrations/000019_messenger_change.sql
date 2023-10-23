-- +goose Up
-- +goose MessengerChangeBegin
alter table l1_messenger_events
    add amount numeric(32);
alter table l2_messenger_events
    add amount numeric(32);
alter table l2_messenger_events
    add tx_hash varchar(66) default '' not null;
-- +goose MessengerChangeEnd

-- +goose Down
-- +goose MessengerChangeBegin
alter table l1_messenger_events
drop column amount;
alter table l2_messenger_events
drop column amount;
alter table l2_messenger_events
drop column tx_hash;
-- +goose MessengerChangeEnd