package film

import "database/sql"

type RepoPostgre struct {
	DB *sql.DB
}

func NewPostgreRepository(db *sql.DB) *RepoPostgre {
	return &RepoPostgre{DB: db}
}

func (repo *RepoPostgre) GetFilmsByGenre(genre string, start int, end int) ([]*FilmItem, error) {
	films := make([]*FilmItem, 0, end-start)

	rows, err := repo.DB.Query(
		"SELECT film.id, film.title, poster FROM film"+
			"JOIN films_genre ON film.id = films_genre.id_film"+
			"JOIN genre ON films_genre.id_genre = genre.id"+
			"WHERE genre.title = $1'"+
			"OFFSET $2 LIMIT $3",
		genre, start, end)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &FilmItem{}
		err := rows.Scan(&post.Id, &post.Title, &post.Poster)
		if err != nil {
			return nil, err
		}
		films = append(films, post)
	}

	return films, nil
}
