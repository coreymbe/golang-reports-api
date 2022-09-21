CREATE DATABASE puppet
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

\c "puppet";

CREATE TABLE reports (
    certname character varying(40) NOT NULL,
    environment character varying(40),
    status character varying(40) NOT NULL,
    "time" character varying(40),
    transaction_uuid character varying(40) NOT NULL,
    ID SERIAL PRIMARY KEY
);