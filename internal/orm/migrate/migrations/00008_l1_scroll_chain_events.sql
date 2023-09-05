-- +goose Up
-- +goose L1ScrollChainBegin
create table l1_scroll_chain_events
(
    number          bigint,
    batch_index     bigint,
    batch_hash      text not null
        primary key,
    commit_number   bigint,
    commit_hash     text,
    finalize_number bigint,
    finalize_hash   text,
    revert_number   bigint,
    revert_hash     text,
    l2_start_number bigint,
    l2_end_number   bigint,
    status          smallint
);

comment on column l1_scroll_chain_events.batch_index is ' batch index, increase one by one';

comment on column l1_scroll_chain_events.batch_hash is ' batch content hash, it is unique';

comment on column l1_scroll_chain_events.commit_number is ' commit l1chain number';

comment on column l1_scroll_chain_events.finalize_number is ' finalize l1chain number';

comment on column l1_scroll_chain_events.revert_number is ' revert batch l1chain number';

comment on column l1_scroll_chain_events.l2_start_number is ' l2chain block start number contained in this batch';

comment on column l1_scroll_chain_events.l2_end_number is ' l2chain block end number contained in this batch';

alter table l1_scroll_chain_events
    owner to maskpp;

create index idx_l1_scroll_chain_events_status
    on l1_scroll_chain_events (status);

create index idx_l1_scroll_chain_events_revert_number
    on l1_scroll_chain_events (revert_number);

create index idx_l1_scroll_chain_events_finalize_number
    on l1_scroll_chain_events (finalize_number);

create index idx_l1_scroll_chain_events_commit_number
    on l1_scroll_chain_events (commit_number);

create index idx_l1_scroll_chain_events_batch_index
    on l1_scroll_chain_events (batch_index);

create index idx_l1_scroll_chain_events_number
    on l1_scroll_chain_events (number);
-- +goose L1ScrollChainEnd

-- +goose Down
-- +goose L1ScrollChainBegin
drop table if exists l1_scroll_chain_events
-- +goose L1ScrollChainEnd
