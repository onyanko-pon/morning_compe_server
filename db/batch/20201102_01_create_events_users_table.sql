create table events_users (
    user_id SERIAL,
    event_id SERIAL,
    status INTEGER default 0,
    created_at timestamp,
    updated_at timestamp,
    foreign key (user_id) references users(id),
    foreign key (event_id) references events(id),
    unique (user_id, event_id)
);