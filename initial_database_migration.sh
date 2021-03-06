#!/bin/sh
psql -U postgres << END_OF_SCRIPT

CREATE USER library WITH PASSWORD 'password';
CREATE USER library_test WITH PASSWORD 'password';
CREATE DATABASE library WITH OWNER = libary ENCODING = 'UTF8' TABLESPACE = pg_default LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1;
CREATE DATABASE library_test WITH OWNER = libary_test ENCODING = 'UTF8' TABLESPACE = pg_default LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1;

END_OF_SCRIPT