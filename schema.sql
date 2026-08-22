create table if not exists devices (
    id integer primary key,
    device_id int not null,
    name text default null,
    last_errored_conn datetime default current_timestamp,
    last_successful_conn datetime default current_timestamp,
    version double default 0.1
);

create table if not exists users (
    id integer primary key,
    name varchar(255),
    email varchar(255) default null
);

create table if not exists metrics (
    id integer primary key,
    name text not null,
    description text default null,
    value_type text default 'number'
);

create table if not exists user_attributes (
    id integer primary key,
    user_id integer not null,
    height not null,
    weight not null,
    body_type text default null,
    body_fat_pct double default 0.0,
    bmi double default 0.0
);

create table if not exists user_profile (
    id integer primary key,
    user_id integer not null,
    avatar text default null,
    position text default null,
    last_played datetime default null
);

create table if not exists field (
    id integer primary key,
    name text default null,
    width integer default 0,
    height integer default 0,
    total_width integer default 0,
    total_height integer default 0
);