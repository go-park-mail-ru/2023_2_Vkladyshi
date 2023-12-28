CREATE TABLE IF NOT EXISTS users_comment (
	id_user SERIAL NOT NULL
	id_film SERIAL NOT NULL
	comment TEXT NOT NULL DEFAULT '',
	rating INTEGER NOT NULL DEFAULT '5'
		CONSTRAINT rating_is_positive CHECK (rating >= 0),
	
	PRIMARY KEY(id_user, id_film)
);
