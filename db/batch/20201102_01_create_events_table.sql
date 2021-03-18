create table events (
    id SERIAL NOT NULL,
    title VARCHAR(255),
    body text,
    host_user_id SERIAL,
    place VARCHAR(255),
    date DATE,
    start_time TIME,
    end_time TIME,
    created_at timestamp,
    updated_at timestamp,
    status SERIAL,
    PRIMARY KEY (id),
    foreign key (host_user_id) references users(id)
);