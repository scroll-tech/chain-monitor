CREATE TABLE message_match
(
    id                       BIGSERIAL       PRIMARY KEY,
    message_hash             VARCHAR         NOT NULL,
    token_type               INTEGER         NOT NULL,

    -- l1 event info
    l1_event_type            INTEGER         NOT NULL,
    l1_block_number          BIGINT          NOT NULL,
    l1_tx_hash               VARCHAR         NOT NULL,
    l1_token_id              VARCHAR         NOT NULL,
    l1_amount                DECIMAL(78, 0)  NOT NULL,

    -- l1 event info
    l2_event_type            INTEGER         NOT NULL,
    l2_block_number          BIGINT          NOT NULL,
    l2_tx_hash               VARCHAR         NOT NULL,
    l2_token_id              VARCHAR         NOT NULL,
    l2_amount                DECIMAL(78, 0)  NOT NULL,

    -- status
    check_status             INTEGER         NOT NULL,
    withdraw_root_status     INTEGER         NOT NULL,
    l1_gateway_status        INTEGER         NOT NULL,
    l2_gateway_status        INTEGER         NOT NULL,
    l1_cross_chain_status    INTEGER         NOT NULL,
    l2_cross_chain_status    INTEGER         NOT NULL,
    message_proof            BYTEA           DEFAULT NULl,
    message_nonce            BIGINT          DEFAULT NULL,

    created_at               TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at               TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at               TIMESTAMP(0)    DEFAULT NULL
);

CREATE UNIQUE INDEX idx_message_match_message_hash ON message_match (message_hash);
CREATE INDEX idx_message_l2_block_number ON message_match (l2_block_number);
