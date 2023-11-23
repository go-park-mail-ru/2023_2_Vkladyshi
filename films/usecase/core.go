package usecase

import (
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/crew"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/film"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/genre"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/repository/profession"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"
)

type ICore interface {
	GetFilmsByGenre(genre uint64, start uint64, end uint64) ([]models.FilmItem, error)
	GetFilms(start uint64, end uint64) ([]models.FilmItem, error)
	GetFilm(filmId uint64) (*models.FilmItem, error)
	GetFilmGenres(filmId uint64) ([]models.GenreItem, error)
	GetFilmRating(filmId uint64) (float64, uint64, error)
	GetFilmDirectors(filmId uint64) ([]models.CrewItem, error)
	GetFilmScenarists(filmId uint64) ([]models.CrewItem, error)
	GetFilmCharacters(filmId uint64) ([]models.Character, error)
	GetActor(actorId uint64) (*models.CrewItem, error)
	GetActorsCareer(actorId uint64) ([]models.ProfessionItem, error)
	GetGenre(genreId uint64) (string, error)
}

type Core struct {
	lg         *slog.Logger
	films      film.IFilmsRepo
	genres     genre.IGenreRepo
	crew       crew.ICrewRepo
	profession profession.IProfessionRepo
}

func GetCore(cfg_sql configs.DbDsnCfg, lg *slog.Logger) (*Core, error) {
	films, err := film.GetFilmRepo(cfg_sql, lg)
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}
	genres, err := genre.GetGenreRepo(cfg_sql, lg)
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}
	crew, err := crew.GetCrewRepo(cfg_sql, lg)
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}
	professions, err := profession.GetProfessionRepo(cfg_sql, lg)
	if err != nil {
		lg.Error("cant create repo")
		return nil, err
	}
	core := Core{
		lg:         lg.With("module", "core"),
		films:      films,
		genres:     genres,
		crew:       crew,
		profession: professions,
	}
	return &core, nil
}

func (core *Core) GetFilmsByGenre(genre uint64, start uint64, end uint64) ([]models.FilmItem, error) {
	films, err := core.films.GetFilmsByGenre(genre, start, end)
	if err != nil {
		core.lg.Error("failed to get films from db", "err", err.Error())
		return nil, fmt.Errorf("GetFilmsByGenre err: %w", err)
	}

	return films, nil
}

func (core *Core) GetFilms(start uint64, end uint64) ([]models.FilmItem, error) {
	films, err := core.films.GetFilms(start, end)
	if err != nil {
		core.lg.Error("failed to get films from db", "err", err.Error())
		return nil, fmt.Errorf("GetFilms err: %w", err)
	}

	return films, nil
}

func (core *Core) GetFilm(filmId uint64) (*models.FilmItem, error) {
	film, err := core.films.GetFilm(filmId)
	if err != nil {
		core.lg.Error("Get Film error", "err", err.Error())
		return nil, fmt.Errorf("GetFilm err: %w", err)
	}

	return film, nil
}

func (core *Core) GetFilmGenres(filmId uint64) ([]models.GenreItem, error) {
	genres, err := core.genres.GetFilmGenres(filmId)
	if err != nil {
		core.lg.Error("Get Film Genres error", "err", err.Error())
		return nil, fmt.Errorf("GetFilmGenres err: %w", err)
	}

	return genres, nil
}

func (core *Core) GetFilmRating(filmId uint64) (float64, uint64, error) {
	rating, number, err := core.films.GetFilmRating(filmId)
	if err != nil {
		core.lg.Error("Get Film Rating error", "err", err.Error())
		return 0, 0, fmt.Errorf("GetFilmRating err: %w", err)
	}

	return rating, number, nil
}

func (core *Core) GetFilmDirectors(filmId uint64) ([]models.CrewItem, error) {
	directors, err := core.crew.GetFilmDirectors(filmId)
	if err != nil {
		core.lg.Error("Get Film Directors error", "err", err.Error())
		return nil, fmt.Errorf("GetFilmDirectors err: %w", err)
	}

	return directors, nil
}

func (core *Core) GetFilmScenarists(filmId uint64) ([]models.CrewItem, error) {
	scenarists, err := core.crew.GetFilmScenarists(filmId)
	if err != nil {
		core.lg.Error("Get Film Scenarists error", "err", err.Error())
		return nil, fmt.Errorf("GetFilmScenarists err: %w", err)
	}

	return scenarists, nil
}

func (core *Core) GetFilmCharacters(filmId uint64) ([]models.Character, error) {
	characters, err := core.crew.GetFilmCharacters(filmId)
	if err != nil {
		core.lg.Error("Get Film Characters error", "err", err.Error())
		return nil, fmt.Errorf("GetFilmCharacters err: %w", err)
	}

	return characters, nil
}

func (core *Core) GetActor(actorId uint64) (*models.CrewItem, error) {
	actor, err := core.crew.GetActor(actorId)
	if err != nil {
		core.lg.Error("Get Actor error", "err", err.Error())
		return nil, fmt.Errorf("GetActor err: %w", err)
	}

	return actor, nil
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
