CREATE TABLE IF NOT EXISTS provider (
    id varchar(100) NOT NULL,
    name varchar(100) NOT NULL,
    is_selected bool,
    is_deleted bool DEFAULT false,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);