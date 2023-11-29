UPDATE film 
SET fts = setweight(to_tsvector(title), 'A')
	|| setweight(to_tsvector(info), 'B');