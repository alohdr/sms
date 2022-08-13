CREATE TABLE IF NOT EXISTS provider (
    id varchar(100) NOT NULL,
    name varchar(100) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);