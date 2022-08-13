CREATE TABLE IF NOT EXISTS template (
    id varchar(100) NOT NULL,
    type varchar(100) NOT NULL,
    apps_name varchar(255) NOT NULL,
    text varchar(255) NOT NULL,
    is_deleted bool DEFAULT false,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);