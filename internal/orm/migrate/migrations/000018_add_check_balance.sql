-- +goose Up
-- +goose AddCheckBalanceBegin
alter table l2_chain_confirms
    add balance_status boolean default true;
alter table l1_chain_confirms
    add balance_status boolean default true;
-- +goose AddCheckBalanceEnd

-- +goose Down
-- +goose AddCheckBalanceBegin
alter table l1_chain_confirms drop column balance_status;
alter table l2_chain_confirms drop column balance_status;
-- +goose AddCheckBalanceEnd