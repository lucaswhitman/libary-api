CREATE TABLE IF NOT EXISTS books
(
	id BIGINT,
	title TEXT NOT NULL,
	author_id BIGINT REFERENCES authors(id),
	publisher TEXT NOT NULL,
	publish_date DATE,
	rating INT NOT NULL,
	status TEXT NOT NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
)