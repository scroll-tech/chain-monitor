-- +goose Up
-- +goose UpgradeChainMonitorBegin
alter table chain_confirms
    rename to l2_chain_confirms;
alter table l2_chain_confirms
    rename column withdraw_status to withdraw_root_status;
-- +goose UpgradeChainMonitorEnd

-- +goose Down
-- +goose UpgradeChainMonitorBegin
alter table L2_chain_confirms
    rename to chain_confirms;
alter table chain_confirms
    rename column withdraw_root_status to withdraw_status;
-- +goose UpgradeChainMonitorEnd
