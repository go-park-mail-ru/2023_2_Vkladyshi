package film

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"

	_ "github.com/jackc/pgx/stdlib"
)

type IFilmsRepo interface {
	GetFilmsByGenre(genre uint64, start uint64, end uint64) ([]models.FilmItem, error)
	GetFilms(start uint64, end uint64) ([]models.FilmItem, error)
	GetFilm(filmId uint64) (*models.FilmItem, error)
	GetFilmRating(filmId uint64) (float64, uint64, error)
	FindFilm(title string, dateFrom string, dateTo string,
		ratingFrom float32, ratingTo float32, mpaa string, genres []string, actors []string) ([]models.FilmItem, error)
}

type RepoPostgre struct {
	db *sql.DB
}

func GetFilmRepo(config configs.DbDsnCfg, lg *slog.Logger) (*RepoPostgre, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.DbName, config.Password, config.Host, config.Port, config.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		lg.Error("sql open error", "err", err.Error())
		return nil, fmt.Errorf("get film repo: %w", err)
	}
	err = db.Ping()
	if err != nil {
		lg.Error("sql ping error", "err", err.Error())
		return nil, fmt.Errorf("get film repo: %w", err)
	}
	db.SetMaxOpenConns(config.MaxOpenConns)

	postgreDb := RepoPostgre{db: db}

	go postgreDb.pingDb(config.Timer, lg)
	return &postgreDb, nil
}

func (repo *RepoPostgre) pingDb(timer uint32, lg *slog.Logger) {
	for {
		err := repo.db.Ping()
		if err != nil {
			lg.Error("Repo Film db ping error", "err", err.Error())
		}

		time.Sleep(time.Duration(timer) * time.Second)
	}
}

func (repo *RepoPostgre) GetFilmsByGenre(genre uint64, start uint64, end uint64) ([]models.FilmItem, error) {
	films := make([]models.FilmItem, 0, end-start)

	rows, err := repo.db.Query(
		"SELECT film.id, film.title, poster FROM film "+
			"JOIN films_genre ON film.id = films_genre.id_film "+
			"WHERE id_genre = $1 "+
			"ORDER BY release_date DESC "+
			"OFFSET $2 LIMIT $3",
		genre, start, end)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("GetFilmsByGenre err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		post := models.FilmItem{}
		err := rows.Scan(&post.Id, &post.Title, &post.Poster)
		if err != nil {
			return nil, fmt.Errorf("GetFilmsByGenre scan err: %w", err)
		}
		films = append(films, post)
	}

	return films, nil
}

func (repo *RepoPostgre) GetFilms(start uint64, end uint64) ([]models.FilmItem, error) {
	films := make([]models.FilmItem, 0, end-start)

	rows, err := repo.db.Query(
		"SELECT film.id, film.title, poster FROM film "+
			"ORDER BY release_date DESC "+
			"OFFSET $1 LIMIT $2",
		start, end)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("GetFilms err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		post := models.FilmItem{}
		err := rows.Scan(&post.Id, &post.Title, &post.Poster)
		if err != nil {
			return nil, fmt.Errorf("GetFilms scan err: %w", err)
		}
		films = append(films, post)
	}

	return films, nil
}

func (repo *RepoPostgre) GetFilm(filmId uint64) (*models.FilmItem, error) {
	film := &models.FilmItem{}
	err := repo.db.QueryRow(
		"SELECT * FROM film "+
			"WHERE id = $1", filmId).
		Scan(&film.Id, &film.Title, &film.Info, &film.Poster, &film.ReleaseDate, &film.Country, &film.Mpaa)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return film, nil
		}

		return nil, fmt.Errorf("GetFilm err: %w", err)
	}

	return film, nil
}

func (repo *RepoPostgre) GetFilmRating(filmId uint64) (float64, uint64, error) {
	var rating sql.NullFloat64
	var number sql.NullInt64
	err := repo.db.QueryRow(
		"SELECT AVG(rating), COUNT(rating) FROM users_comment "+
			"WHERE id_film = $1", filmId).Scan(&rating, &number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, 0, nil
		}
		return 0, 0, fmt.Errorf("GetFilmRating err: %w", err)
	}

	return rating.Float64, uint64(number.Int64), nil
}

func (repo *RepoPostgre) FindFilm(title string, dateFrom string, dateTo string,
	ratingFrom float32, ratingTo float32, mpaa string, genres []string, actors []string) ([]models.FilmItem, error) {

	films := []models.FilmItem{}

	var s strings.Builder
	s.WriteString(
		"SELECT DISTINCT film.title, film.id, film.poster, AVG(users_comment.rating) FROM film " +
			"JOIN films_genre ON film.id = films_genre.id_film " +
			"JOIN genre ON genre.id = films_genre.id_genre " +
			"JOIN users_comment ON film.id = users_comment.id_film " +
			"JOIN person_in_film ON film.id = person_in_film.id_film " +
			"JOIN crew ON person_in_film.id_person = crew.id WHERE ")
	if title != "" {
		s.WriteString("fts @@ to_tsquery($1) AND ")
	}
	if dateFrom != "" {
		s.WriteString("release_date >= '$2' AND ")
	}
	if dateTo != "" {
		s.WriteString("release_date <= '$3' AND ")
	}
	if mpaa != "" {
		s.WriteString("mpaa = $8 AND ")
	}
	s.WriteString(
		"(CASE WHEN array_length($4::varchar[], 1)> 0 " +
			"THEN genre.title = ANY ($4::varchar[]) ELSE TRUE END) AND (CASE " +
			"WHEN array_length($5::varchar[], 1)> 0 " +
			"THEN crew.name = ANY ($5::varchar[]) ELSE TRUE END) " +

			"GROUP BY film.title, film.id, genre.title " +
			"HAVING AVG(users_comment.rating) > $6 AND AVG(users_comment.rating) < $7 " +
			"ORDER BY film.title")

	rows, err := repo.db.Query(s.String(), title, dateFrom, dateTo, genres, actors, ratingFrom, ratingTo, mpaa)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("find film err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		post := models.FilmItem{}
		err := rows.Scan(&post.Id, &post.Title, &post.Poster)
		if err != nil {
			return nil, fmt.Errorf("find film scan err: %w", err)
		}
		films = append(films, post)
	}

	return films, nil
}

func checkTitle(title string) string {
	if title != "" {
		return ""
	}

	return ""
}
