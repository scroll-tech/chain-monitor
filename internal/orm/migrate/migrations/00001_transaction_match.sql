-- +goose Up
-- +goose ChainMonitorBegin
create table chain_confirms
(
    number          bigserial
        primary key,
    withdraw_status boolean,
    deposit_status  boolean,
    confirm         boolean
);
-- +goose ChainMonitorEnd

-- +goose Down
-- +goose ChainMonitorBegin
drop table if exists chain_confirms;
-- +goose ChainMonitorEnd
