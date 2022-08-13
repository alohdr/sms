CREATE TABLE IF NOT EXISTS users (
    id varchar(100) NOT NULL,
    username varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    is_deleted bool DEFAULT false,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);