CREATE DATABASE puppet
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

\c "puppet";

CREATE TABLE reports (
    ID SERIAL PRIMARY KEY,
    certname varchar(40) NOT NULL,
    environment varchar(40),
    status varchar(40) NOT NULL,
    time varchar(40),
    transaction_uuid varchar(50) NOT NULL
);

CREATE TABLE users (
    username varchar(40) NOT NULL,
    password varchar(40) NOT NULL
);

INSERT INTO users (username, password) VALUES ('admin', 'ch@ngem3');