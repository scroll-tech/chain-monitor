CREATE TABLE message_match
(
    id                       BIGSERIAL       PRIMARY KEY,
    message_hash             VARCHAR         NOT NULL,
    token_type               INTEGER         NOT NULL,

    -- l1 event info
    l1_event_type            INTEGER         NOT NULL,
    l1_block_number          BIGINT          NOT NULL,
    l1_tx_hash               VARCHAR         NOT NULL,
    l1_token_ids             VARCHAR         NOT NULL,
    l1_amounts               VARCHAR         NOT NULL,

    -- l2 event info
    l2_event_type            INTEGER         NOT NULL,
    l2_block_number          BIGINT          NOT NULL,
    l2_tx_hash               VARCHAR         NOT NULL,
    l2_token_ids             VARCHAR         NOT NULL,
    l2_amounts               VARCHAR         NOT NULL,

    -- eth info
    l1_messenger_eth_balance DECIMAL(78, 0)  NOT NULL,
    l2_messenger_eth_balance DECIMAL(78, 0)  NOT NULL,
    l1_eth_balance_status    INTEGER         NOT NULL,
    l2_eth_balance_status    INTEGER         NOT NULL,

    -- status
    check_status             INTEGER         NOT NULL,
    l1_block_status          INTEGER         NOT NULL,
    l2_block_status          INTEGER         NOT NULL,
    l1_cross_chain_status    INTEGER         NOT NULL,
    l2_cross_chain_status    INTEGER         NOT NULL,
    message_proof            BYTEA           DEFAULT NULl,
    message_nonce            BIGINT          DEFAULT NULL,

    created_at               TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at               TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at               TIMESTAMP(0)    DEFAULT NULL
);

CREATE UNIQUE INDEX idx_message_match_message_hash ON message_match (message_hash);
CREATE INDEX if not exists idx_check_status_deleted_at ON message_match (check_status, deleted_at);
CREATE INDEX if not exists idx_l1_l2_block_status_deleted_id_desc ON message_match (l1_block_status, l2_block_status, deleted_at, id DESC);
CREATE INDEX if not exists idx_message_l1_block_number ON message_match (l1_block_number);
CREATE INDEX if not exists idx_message_l2_block_number ON message_match (l2_block_number);
CREATE UNIQUE INDEX idx_message_match_message_nonce ON message_match (message_nonce);
