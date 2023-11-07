package delivery

import (
	"log/slog"
	"math/rand"
	"regexp"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/errors"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/comment"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/crew"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/film"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/genre"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/profession"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/profile"
)

type Core struct {
	sessions   map[string]Session
	mutex      sync.RWMutex
	lg         *slog.Logger
	Films      film.IFilmsRepo
	Users      profile.IUserRepo
	Genres     genre.IGenreRepo
	Comments   comment.ICommentRepo
	Crew       crew.ICrewRepo
	Profession profession.IProfessionRepo
}

type Session struct {
	Login     string
	ExpiresAt time.Time
}

func GetCore(cfg configs.DbDsnCfg, lg *slog.Logger) *Core {
	core := Core{
		sessions:   make(map[string]Session),
		lg:         lg.With("module", "core"),
		Films:      film.GetFilmRepo(cfg, lg),
		Users:      profile.GetUserRepo(cfg, lg),
		Genres:     genre.GetGenreRepo(cfg, lg),
		Comments:   comment.GetCommentRepo(cfg, lg),
		Crew:       crew.GetCrewRepo(cfg, lg),
		Profession: profession.GetProfessionRepo(cfg, lg),
	}
	return &core
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (core *Core) CreateSession(login string) (string, Session, error) {
	SID := RandStringRunes(32)

	session := Session{
		Login:     login,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	core.mutex.Lock()
	core.sessions[SID] = session
	core.mutex.Unlock()

	return SID, session, nil
}

func (core *Core) KillSession(sid string) error {
	core.mutex.Lock()
	delete(core.sessions, sid)
	core.mutex.Unlock()
	return nil
}

func (core *Core) FindActiveSession(sid string) (bool, error) {
	core.mutex.RLock()
	_, found := core.sessions[sid]
	core.mutex.RUnlock()
	return found, nil
}

func (core *Core) CreateUserAccount(login string, password string, name string, birthDate string, email string) error {
	if matched, _ := regexp.MatchString(`^\w@\w$`, email); !matched {
		return errors.InvalideEmail
	}
	err := core.Users.CreateUser(login, password, name, birthDate, email)
	if err != nil {
		core.lg.Error("create user error", "err", err.Error())
		return err
	}

	return nil
}

func (core *Core) FindUserAccount(login string, password string) (*profile.UserItem, bool, error) {
	user, found, err := core.Users.GetUser(login, password)
	if err != nil {
		core.lg.Error("find user error", "err", err.Error())
		return nil, false, err
	}
	return user, found, nil
}

func (core *Core) FindUserByLogin(login string) (bool, error) {
	found, err := core.Users.FindUser(login)
	if err != nil {
		core.lg.Error("find user error", "err", err.Error())
		return false, err
	}

	return found, nil
}

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}

func (core *Core) GetFilmsByGenre(genre string, start uint64, end uint64) ([]film.FilmItem, error) {
	films, err := core.Films.GetFilmsByGenre(genre, start, end)
	if err != nil {
		core.lg.Error("failed to get films from db", "err", err.Error())
		return nil, err
	}

	return films, nil
}

func (core *Core) GetFilms(start uint64, end uint64) ([]film.FilmItem, error) {
	films, err := core.Films.GetFilms(start, end)
	if err != nil {
		core.lg.Error("failed to get films from db", "err", err.Error())
		return nil, err
	}

	return films, nil
}

func (core *Core) GetFilm(filmId uint64) (*film.FilmItem, error) {
	film, err := core.Films.GetFilm(filmId)
	if err != nil {
		core.lg.Error("Get Film error", "err", err.Error())
		return nil, err
	}

	return film, nil
}

func (core *Core) GetFilmGenres(filmId uint64) ([]genre.GenreItem, error) {
	genres, err := core.Genres.GetFilmGenres(filmId)
	if err != nil {
		core.lg.Error("Get Film Genres error", "err", err.Error())
		return nil, err
	}

	return genres, nil
}

func (core *Core) GetFilmRating(filmId uint64) (float64, uint64, error) {
	rating, number, err := core.Comments.GetFilmRating(filmId)
	if err != nil {
		core.lg.Error("Get Film Rating error", "err", err.Error())
		return 0, 0, err
	}

	return rating, number, nil
}

func (core *Core) GetFilmDirectors(filmId uint64) ([]crew.CrewItem, error) {
	directors, err := core.Crew.GetFilmDirectors(filmId)
	if err != nil {
		core.lg.Error("Get Film Directors error", "err", err.Error())
		return nil, err
	}

	return directors, nil
}

func (core *Core) GetFilmScenarists(filmId uint64) ([]crew.CrewItem, error) {
	scenarists, err := core.Crew.GetFilmScenarists(filmId)
	if err != nil {
		core.lg.Error("Get Film Scenarists error", "err", err.Error())
		return nil, err
	}

	return scenarists, nil
}

func (core *Core) GetFilmCharacters(filmId uint64) ([]crew.Character, error) {
	characters, err := core.Crew.GetFilmCharacters(filmId)
	if err != nil {
		core.lg.Error("Get Film Characters error", "err", err.Error())
		return nil, err
	}

	return characters, nil
}

func (core *Core) GetFilmComments(filmId uint64, first uint64, limit uint64) ([]comment.CommentItem, error) {
	comments, err := core.Comments.GetFilmComments(filmId, first, limit)
	if err != nil {
		core.lg.Error("Get Film Comments error", "err", err.Error())
		return nil, err
	}

	return comments, nil
}

func (core *Core) GetActor(actorId uint64) (*crew.CrewItem, error) {
	actor, err := core.Crew.GetActor(actorId)
	if err != nil {
		core.lg.Error("Get Actor error", "err", err.Error())
		return nil, err
	}

	return actor, nil
}

func (core *Core) GetActorsCareer(actorId uint64) ([]profession.ProfessionItem, error) {
	career, err := core.Profession.GetActorsProfessions(actorId)
	if err != nil {
		core.lg.Error("Get Actors Career error", "err", err.Error())
		return nil, err
	}

	return career, nil
}

func (core *Core) AddComment(filmId uint64, userId uint64, rating uint16, text string) (bool, error) {
	err := core.Comments.AddComment(filmId, userId, rating, text)
	if err != nil {
		core.lg.Error("Add Comment error", "err", err.Error())
		return false, err
	}

	return true, nil
}