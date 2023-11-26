package usecase

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/crew"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/film"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/genre"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/profession"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/requests"
)

type ICore interface {
	GetFilmsAndGenreTitle(genreId uint64, start uint64, end uint64) ([]models.FilmItem, string, error)
	GetFilmInfo(filmId uint64) (*requests.FilmResponse, error)
	GetActorInfo(actorId uint64) (*requests.ActorResponse, error)
	GetActorsCareer(actorId uint64) ([]models.ProfessionItem, error)
	GetGenre(genreId uint64) (string, error)
	FindFilm(title string, dateFrom string, dateTo string,
		ratingFrom float32, ratingTo float32, mpaa string, genres []string, actors []string) ([]models.FilmItem, error)
}

type Core struct {
	lg         *slog.Logger
	films      film.IFilmsRepo
	genres     genre.IGenreRepo
	crew       crew.ICrewRepo
	profession profession.IProfessionRepo
}

func GetCore(cfg_sql *configs.DbDsnCfg, lg *slog.Logger) (*Core, error) {
	var films film.IFilmsRepo
	var genres genre.IGenreRepo
	var actors crew.ICrewRepo
	var professions profession.IProfessionRepo
	var err error

	switch cfg_sql.Films_db {
	case "postgres":
		films, err = film.GetFilmRepo(*cfg_sql, lg)
	}
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}

	switch cfg_sql.Genres_db {
	case "postgres":
		genres, err = genre.GetGenreRepo(*cfg_sql, lg)
	}
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}

	switch cfg_sql.Crew_db {
	case "postgres":
		actors, err = crew.GetCrewRepo(*cfg_sql, lg)
	}
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}

	switch cfg_sql.Profession_db {
	case "postgres":
		professions, err = profession.GetProfessionRepo(*cfg_sql, lg)
	}
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}
	core := Core{
		lg:         lg.With("module", "core"),
		films:      films,
		genres:     genres,
		crew:       actors,
		profession: professions,
	}
	return &core, nil
}

func (core *Core) GetFilmsAndGenreTitle(genreId uint64, start uint64, end uint64) ([]models.FilmItem, string, error) {
	var films []models.FilmItem
	var err error

	if genreId == 0 {
		films, err = core.films.GetFilms(start, end)
	} else {
		films, err = core.films.GetFilmsByGenre(genreId, start, end)
	}
	if err != nil {
		core.lg.Error("failed to get films from db", "err", err.Error())
		return nil, "", fmt.Errorf("GetFilms err: %w", err)
	}

	genre, err := core.genres.GetGenreById(genreId)
	if err != nil {
		core.lg.Error("failed to get genre by id", "err", err.Error())
		return nil, "", fmt.Errorf("GetFilms err: %w", err)
	}

	return films, genre, nil
}

func (core *Core) GetFilmInfo(filmId uint64) (*requests.FilmResponse, error) {
	film, err := core.films.GetFilm(filmId)
	if err != nil {
		core.lg.Error("get film error", "err", err.Error())
		return nil, fmt.Errorf("get film err: %w", err)
	}
	if film.Title == "" {
		return nil, errNotFound()
	}

	genres, err := core.genres.GetFilmGenres(filmId)
	if err != nil {
		core.lg.Error("get film genres error", "err", err.Error())
		return nil, fmt.Errorf("get film genres err: %w", err)
	}

	rating, number, err := core.films.GetFilmRating(filmId)
	if err != nil {
		core.lg.Error("get film rating error", "err", err.Error())
		return nil, fmt.Errorf("get film rating err: %w", err)
	}

	directors, err := core.crew.GetFilmDirectors(filmId)
	if err != nil {
		core.lg.Error("get film directors error", "err", err.Error())
		return nil, fmt.Errorf("get film directors err: %w", err)
	}

	scenarists, err := core.crew.GetFilmScenarists(filmId)
	if err != nil {
		core.lg.Error("get film scenarists error", "err", err.Error())
		return nil, fmt.Errorf("get film scenarists err: %w", err)
	}

	characters, err := core.crew.GetFilmCharacters(filmId)
	if err != nil {
		core.lg.Error("get film characters error", "err", err.Error())
		return nil, fmt.Errorf("get film scenarists err: %w", err)
	}

	result := requests.FilmResponse{
		Film:       *film,
		Genres:     genres,
		Rating:     rating,
		Number:     number,
		Directors:  directors,
		Scenarists: scenarists,
		Characters: characters,
	}

	return &result, nil
}

func (core *Core) GetActorInfo(actorId uint64) (*requests.ActorResponse, error) {
	actor, err := core.crew.GetActor(actorId)
	if err != nil {
		core.lg.Error("get actor error", "err", err.Error())
		return nil, fmt.Errorf("get actor err: %w", err)
	}
	if actor.Name == "" {
		return nil, errNotFound()
	}

	career, err := core.profession.GetActorsProfessions(actorId)
	if err != nil {
		core.lg.Error("get actor profession error", "err", err.Error())
		return nil, fmt.Errorf("get actor profession err: %w", err)
	}

	result := requests.ActorResponse{
		Name:      actor.Name,
		Photo:     actor.Photo,
		BirthDate: actor.Birthdate,
		Country:   actor.Country,
		Info:      actor.Info,
		Career:    career,
	}
	return &result, nil
}

func (core *Core) GetActorsCareer(actorId uint64) ([]models.ProfessionItem, error) {
	career, err := core.profession.GetActorsProfessions(actorId)
	if err != nil {
		core.lg.Error("Get Actors Career error", "err", err.Error())
		return nil, fmt.Errorf("GetActorsCareer err: %w", err)
	}

	return career, nil
}

func (core *Core) GetGenre(genreId uint64) (string, error) {
	genre, err := core.genres.GetGenreById(genreId)
	if err != nil {
		core.lg.Error("GetGenre error", "err", err.Error())
		return "", fmt.Errorf("GetGenre err: %w", err)
	}

	return genre, nil
}

func (core *Core) FindFilm(title string, dateFrom string, dateTo string,
	ratingFrom float32, ratingTo float32, mpaa string, genres []string, actors []string) ([]models.FilmItem, error) {

	films, err := core.films.FindFilm(title, dateFrom, dateTo, ratingFrom, ratingTo, mpaa, genres, actors)
	if err != nil {
		core.lg.Error("find film error", "err", err.Error())
		return nil, fmt.Errorf("find film err: %w", err)
	}

	return films, nil
}

func errNotFound() error {
	return errors.New("not found")
}
