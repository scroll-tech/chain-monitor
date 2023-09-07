-- +goose Up
-- +goose L1ChainConfirmBegin
create table l1_chain_confirms
(
    number          bigserial
        primary key,
    withdraw_status boolean,
    confirm         boolean
);
-- +goose L1ChainConfirmEnd

-- +goose Down
-- +goose L1ChainConfirmBegin
drop table if exists l1_chain_confirms;
-- +goose L1ChainConfirmEnd
