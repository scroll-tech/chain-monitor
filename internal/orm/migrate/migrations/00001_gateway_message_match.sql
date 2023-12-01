-- +goose Up
-- +goose GatewayMessageMatchBegin
CREATE TABLE gateway_message_match
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

    -- status
    l1_block_status                  INTEGER         NOT NULL,
    l2_block_status                  INTEGER         NOT NULL,
    l1_cross_chain_status            INTEGER         NOT NULL,
    l2_cross_chain_status            INTEGER         NOT NULL,

    l1_block_status_updated_at       TIMESTAMP(0)    DEFAULT NULL,
    l2_block_status_updated_at       TIMESTAMP(0)    DEFAULT NULL,
    l1_cross_chain_status_updated_at TIMESTAMP(0)    DEFAULT NULL,
    l2_cross_chain_status_updated_at TIMESTAMP(0)    DEFAULT NULL,
    created_at                       TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                       TIMESTAMP(0)    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at                       TIMESTAMP(0)    DEFAULT NULL
);

CREATE UNIQUE INDEX if not exists idx_gmm_message_hash ON gateway_message_match (message_hash);
CREATE INDEX if not exists idx_gmm_l1_l2_l1cc_id ON gateway_message_match (l1_block_status, l2_block_status, l1_cross_chain_status, id);
CREATE INDEX if not exists idx_gmm_l1_l2_l2cc_id ON gateway_message_match (l1_block_status, l2_block_status, l2_cross_chain_status,  id);
CREATE INDEX if not exists idx_gmm_l1block_l1blocknum_id ON gateway_message_match (l1_block_status, l1_block_number desc, id DESC);
CREATE INDEX if not exists idx_gmm_l2block_l2blocknum_id ON gateway_message_match (l2_block_status, l2_block_number desc, id DESC);
-- +goose GatewayMessageMatchEnd

-- +goose Down
-- +goose GatewayMessageMatchBegin
drop table if exists gateway_message_match;
-- +goose GatewayMessageMatchEnd
