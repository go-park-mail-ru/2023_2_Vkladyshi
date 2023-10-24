CREATE TABLE crew ( 
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL DEFAULT '',
	birth_date DATE NOT NULL DEFAULT CURRENT_DATE,
	photo VARCHAR(100) NOT NULL DEFAULT ''
);

CREATE TABLE profession (
	id SERIAL PRIMARY KEY,
	title VARCHAR(20) NOT NULL DEFAULT ''
);

CREATE TABLE film (
	id SERIAL PRIMARY KEY,
	title VARCHAR(50) NOT NULL DEFAULT '',
	info TEXT NOT NULL DEFAULT '',
	poster VARCHAR(100) NOT NULL DEFAULT '',
	out_date DATE NOT NULL DEFAULT CURRENT_DATE,
	country VARCHAR(20) NOT NULL DEFAULT '',
	mpa VARCHAR(10) NOT NULL DEFAULT ''
);

CREATE TABLE person_in_film (
	id_person SERIAL NOT NULL REFERENCES crew(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	id_film SERIAL NOT NULL REFERENCES film(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	id_profession SERIAL NOT NULL REFERENCES profession(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	character_name VARCHAR(30) DEFAULT '',
	
	PRIMARY KEY(id_person, id_film, id_profession)
);

CREATE TABLE genre (
	id SERIAL PRIMARY KEY,
	title VARCHAR(20) NOT NULL DEFAULT ''
);

CREATE TABLE films_genre (
	id_film SERIAL NOT NULL REFERENCES film(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	id_genre SERIAL NOT NULL REFERENCES genre(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	
	PRIMARY KEY(id_film, id_genre)
);

CREATE TABLE profiles (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL DEFAULT '',
	birth_date DATE NOT NULL DEFAULT CURRENT_DATE,
	photo VARCHAR(100) NOT NULL DEFAULT '',
	login VARCHAR(100) NOT NULL UNIQUE DEFAULT '',
	password VARCHAR(255) NOT NULL DEFAULT '', --Unique?
	registration_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_comment (
	id_user SERIAL NOT NULL REFERENCES profiles(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	id_film SERIAL NOT NULL REFERENCES film(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	comment TEXT NOT NULL DEFAULT '',
	rating INTEGER NOT NULL DEFAULT '5'
		CONSTRAINT rating_is_positive CHECK (rating >= 0),
	
	PRIMARY KEY(id_user, id_film)
);
