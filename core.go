package main

import (
	"log/slog"
	"math/rand"
	"sync"
	"time"
)

type Core struct {
	sessions map[string]Session
	users    map[string]User
	Mutex    sync.RWMutex
	lg       *slog.Logger
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (core *Core) CreateSession(login string) (string, Session) {
	SID := RandStringRunes(32)

	session := Session{
		Login:     login,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	core.Mutex.Lock()
	core.sessions[SID] = session
	core.Mutex.Unlock()

	return SID, session
}

func (core *Core) KillSession(sid string) {
	core.Mutex.Lock()
	delete(core.sessions, sid)
	core.Mutex.Unlock()
}

func (core *Core) FindActiveSession(sid string) bool {
	core.Mutex.RLock()
	_, found := core.sessions[sid]
	core.Mutex.RUnlock()
	return found
}

func (core *Core) CreateUserAccount(request SignupRequest) {
	core.Mutex.Lock()
	core.users[request.Login] = User{Login: request.Login, Email: request.Email, Password: request.Password}
	core.Mutex.Unlock()
}

func (core *Core) FindUserAccount(login string) (User, bool) {
	core.Mutex.RLock()
	user, found := core.users[login]
	core.Mutex.RUnlock()
	return user, found
}

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}

func GetFilms() []Film {
	films := []Film{}

	films = append(films, Film{"Леди Баг и Супер-Кот: Пробуждение силы", "/static/img/ladybag.png", 7.5, []string{"комедия", "приключения", "фэнтези", "мелодрама"}})
	films = append(films, Film{"Барби", "/static/img/barbie.png", 6.7, []string{"комедия", "приключения", "фэнтези"}})
	films = append(films, Film{"Опенгеймер", "/static/img/oppenheimer.png", 8.5, []string{"биография", "драма", "история"}})
	films = append(films, Film{"Слуга народа", "/static/img/sluga_naroda.png", 0.7, []string{"комедия"}})
	films = append(films, Film{"Черная Роза", "/static/img/black_rose.png", 1.5, []string{"детектив", "триллер", "криминал"}})
	films = append(films, Film{"Бесславные ублюдки", "/static/img/inglourious_basterds.png", 8.0, []string{"боевик", "военный", "драма", "комедия"}})
	films = append(films, Film{"Бэтмен: Начало", "/static/img/batman_begins.png", 7.9, []string{"боевик", "фантастика", "драма", "приключения"}})
	films = append(films, Film{"Криминальное чтиво", "/static/img/pulp_fiction.png", 8.6, []string{"криминал", "драма"}})
	films = append(films, Film{"Терминатор", "/static/img/terminator.png", 8.0, []string{"боевик", "фантастика", "триллер"}})

	return films
}
