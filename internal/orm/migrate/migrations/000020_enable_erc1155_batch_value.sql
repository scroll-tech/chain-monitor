-- +goose Up
-- +goose EnableERC1155BatchValueBegin
alter table l1_erc1155_events
    rename column token_id to token_id_list;
alter table l1_erc1155_events
alter column token_id_list type varchar(32) using token_id_list::varchar(32);
alter table l1_erc1155_events
    alter column token_id_list set default '';

alter table l1_erc1155_events
    rename column token_id to token_id_list;
alter table l1_erc1155_events
alter column token_id_list type varchar(32) using token_id_list::varchar(32);
alter table l1_erc1155_events
    alter column token_id_list set default '';

alter table l1_erc1155_events
    rename column amount to amount_list;
alter table l1_erc1155_events
alter column amount_list type varchar(32) using amount_list::varchar(32);
alter table l1_erc1155_events
    alter column amount_list set default '';

alter table l1_erc1155_events
    rename column amount to amount_list;
alter table l1_erc1155_events
alter column amount_list type varchar(32) using amount_list::varchar(32);
alter table l1_erc1155_events
    alter column amount_list set default '';
-- +goose EnableERC1155BatchValueEnd

-- +goose Down
-- +goose EnableERC1155BatchValueBegin
alter table l1_erc1155_events
    rename column token_id_list to token_id;
alter table l1_erc1155_events
alter column token_id type numeric(32) using token_id::numeric(32);

alter table l1_erc1155_events
    rename column token_id_list to token_id;
alter table l1_erc1155_events
alter column token_id type numeric(32) using token_id::numeric(32);
-- +goose EnableERC1155BatchValueEnd