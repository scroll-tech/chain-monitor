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

CREATE UNIQUE INDEX new2idx_message_match_message_hash ON message_match (message_hash);
CREATE INDEX if not exists new2idx_l1_l2_block_status_deleted_id_desc ON message_match (l1_block_status, l2_block_status, deleted_at, id DESC);
CREATE INDEX if not exists new2idx_message_l1_block_number ON message_match (l1_block_number);
CREATE INDEX if not exists new2idx_message_l2_block_number ON message_match (l2_block_number);
CREATE INDEX new2idxidx_message_match_next_message_nonce ON message_match (next_message_nonce);
CREATE INDEX new2idxidx_message_match_withdraw_root_status ON message_match (withdraw_root_status);
-- +goose MessageMatchEnd

-- +goose Down
-- +goose MessageMatchBegin
drop table if exists message_match;
-- +goose MessageMatchEnd
