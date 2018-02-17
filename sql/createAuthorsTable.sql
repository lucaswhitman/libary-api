CREATE TABLE IF NOT EXISTS authors
(
	id BIGINT,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	CONSTRAINT authors_pkey PRIMARY KEY (id)
)