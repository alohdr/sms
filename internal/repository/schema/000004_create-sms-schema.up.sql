CREATE TABLE IF NOT EXISTS sms_history (
    id varchar(100) NOT NULL,
    sender varchar(255) NOT NULL,
    phone_number varchar(100) NOT NULL, 
    type varchar(255) NOT NULL ,
    status varchar(255) NOT NULL ,
  	is_deleted boolean NOT NULL DEFAULT false,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  	PRIMARY KEY(id)
);