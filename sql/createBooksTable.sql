CREATE TABLE IF NOT EXISTS books
(
	id SERIAL,
	title TEXT NOT NULL,
	author TEXT NOT NULL,
	publisher TEXT NOT NULL,
	publish_date DATE,
	rating INT NOT NULL,
	status TEXT NOT NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
)