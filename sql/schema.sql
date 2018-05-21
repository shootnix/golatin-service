create table api_users (
	id                serial,
	description       text not null default '',
	api_key           varchar(64) unique,
	total_requests    integer  not null default 0,
	requests_allowed  integer  not null default 0,
	requests_count    integer  not null default 0,
	updated_at        timestamp not null default now(),
	disabled          integer  not null default 0,
	type_id	          integer foreign key references api_user_types(id) on delete cascade
);