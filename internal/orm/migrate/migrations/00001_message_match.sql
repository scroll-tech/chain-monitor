-- +goose Up
-- +goose MessageMatchBegin
CREATE TABLE message_match
(
    id                               BIGSERIAL       PRIMARY KEY,
    message_hash                     VARCHAR         NOT NULL,
    token_type                       INTEGER         NOT NULL,

    -- l1 event info
    l1_event_type                    INTEGER         NOT NULL,
    l1_block_number                  BIGINT          NOT NULL,
    l1_tx_hash                       VARCHAR         NOT NULL,
    l1_token_ids                     VARCHAR         NOT NULL,
    l1_amounts                       VARCHAR         NOT NULL,

    -- l2 event info
    l2_event_type                    INTEGER         NOT NULL,
    l2_block_number                  BIGINT          NOT NULL,
    l2_tx_hash                       VARCHAR         NOT NULL,
    l2_token_ids                     VARCHAR         NOT NULL,
    l2_amounts                       VARCHAR         NOT NULL,

    -- eth info
    l1_messenger_eth_balance         DECIMAL(78, 0)  NOT NULL,
    l2_messenger_eth_balance         DECIMAL(78, 0)  NOT NULL,
    l1_eth_balance_status            INTEGER         NOT NULL,
    l2_eth_balance_status            INTEGER         NOT NULL,

    -- status
    l1_block_status                  INTEGER         NOT NULL,
    l2_block_status                  INTEGER         NOT NULL,
    l1_cross_chain_status            INTEGER         NOT NULL,
    l2_cross_chain_status            INTEGER         NOT NULL,
    withdraw_root_status             INTEGER         NOT NULL,
    message_proof                    BYTEA           DEFAULT NULl,
    next_message_nonce               BIGINT          DEFAULT NULL,

    l1_block_status_updated_at       TIMESTAMP(0)    DEFAULT NULL,
    l2_block_status_updated_at       TIMESTAMP(0)    DEFAULT NULL,
    l1_cross_chain_status_updated_at TIMESTAMP(0)    DEFAULT NULL,
    l2_cross_chain_status_updated_at TIMESTAMP(0)    DEFAULT NULL,
    l1_eth_balance_status_updated_at TIMESTAMP(0)    DEFAULT NULL,
    l2_eth_balance_status_updated_at TIMESTAMP(0)    DEFAULT NULL,
    message_proof_updated_at         TIMESTAMP(0)    DEFAULT NULL,
    created_at                       TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                       TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at                       TIMESTAMP(0)    DEFAULT NULL
);

CREATE UNIQUE INDEX if not exists idx_message_match_message_hash ON message_match (message_hash);
CREATE INDEX if not exists idx_l1_l2_l1cc_id ON message_match (l1_block_status, l2_block_status, l1_cross_chain_status, id);
CREATE INDEX if not exists idx_l1_l2_l2cc_id ON message_match (l1_block_status, l2_block_status, l2_cross_chain_status,  id);
CREATE INDEX if not exists idx_l1eth_tokentype_id ON message_match (l1_eth_balance_status, token_type,  id);
CREATE INDEX if not exists idx_l2eth_tokentype_id ON message_match (l2_eth_balance_status, token_type, id);
CREATE INDEX if not exists idx_l1block_l1blocknum_id ON message_match (l1_block_status, l1_block_number desc, id DESC);
CREATE INDEX if not exists idx_l2block_l2blocknum_id ON message_match (l2_block_status, l2_block_number desc, id DESC);
CREATE INDEX if not exists idx_withdrawroot_nextnonce_id ON message_match (withdraw_root_status, next_message_nonce DESC, id);
CREATE INDEX if not exists idx_withdrawroot_l2blocknumber_nextnonce ON message_match (withdraw_root_status, l2_block_number, next_message_nonce);
-- +goose MessageMatchEnd

-- +goose Down
-- +goose MessageMatchBegin
drop table if exists message_match;
-- +goose MessageMatchEnd
