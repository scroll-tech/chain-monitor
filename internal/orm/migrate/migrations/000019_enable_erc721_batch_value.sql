-- +goose Up
-- +goose EnableERC721BatchValueBegin
alter table l1_erc721_events
    rename column token_id to token_id_list;
alter table l1_erc721_events
alter column token_id_list type varchar(32) using token_id_list::varchar(32);
alter table l1_erc721_events
    alter column token_id_list set default '';

alter table l2_erc721_events
    rename column token_id to token_id_list;
alter table l2_erc721_events
alter column token_id_list type varchar(32) using token_id_list::varchar(32);
alter table l2_erc721_events
    alter column token_id_list set default '';
-- +goose EnableERC721BatchValueEnd

-- +goose Down
-- +goose EnableERC721BatchValueBegin
alter table l2_erc721_events
    rename column token_id_list to token_id;
alter table l2_erc721_events
alter column token_id type numeric(32) using token_id::numeric(32);

alter table l1_erc721_events
    rename column token_id_list to token_id;
alter table l1_erc721_events
alter column token_id type numeric(32) using token_id::numeric(32);
-- +goose EnableERC721BatchValueEnd