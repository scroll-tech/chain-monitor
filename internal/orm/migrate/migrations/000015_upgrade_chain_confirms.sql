-- +goose Up
-- +goose UpgradeChainMonitorBegin
alter table chain_confirms
    rename column withdraw_status to withdraw_root_status;

alter table chain_confirms
    rename column confirm to deposit_confirm;

alter table chain_confirms
    add withdraw_status boolean;

alter table chain_confirms
    add withdraw_confirm boolean;
-- +goose UpgradeChainMonitorEnd

-- +goose Down
-- +goose UpgradeChainMonitorBegin

alter table chain_confirms
drop column withdraw_status;

alter table chain_confirms
drop column withdraw_confirm;

alter table chain_confirms
    rename column withdraw_root_status to withdraw_status;

alter table chain_confirms
    rename column deposit_confirm to confirm;
-- +goose UpgradeChainMonitorEnd
