-- +goose Up
-- +goose L2ScrollMessengerBegin
alter table l2_messenger_events
drop constraint l2_messenger_events_pkey;
alter table l2_messenger_events
    add primary key (msg_hash);
-- +goose L2ScrollMessengerEnd

-- +goose Down
-- +goose L2ScrollMessengerBegin
alter table l2_messenger_events
drop constraint l2_messenger_events_pkey;
alter table l2_messenger_events
    add primary key (msg_nonce);
-- +goose L2ScrollMessengerEnd
