package main

import (
	"log/slog"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Core struct {
	sessions    map[string]Session
	users       map[string]User
	collections map[string]string
	Mutex       sync.RWMutex
	lg          *slog.Logger
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (core *Core) CreateSession(login string) (string, Session, error) {
	SID := RandStringRunes(32)

	session := Session{
		Login:     login,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	core.Mutex.Lock()
	core.sessions[SID] = session
	core.Mutex.Unlock()

	return SID, session, nil
}

func (core *Core) KillSession(sid string) error {
	core.Mutex.Lock()
	delete(core.sessions, sid)
	core.Mutex.Unlock()
	return nil
}

func (core *Core) FindActiveSession(sid string) (bool, error) {
	core.Mutex.RLock()
	_, found := core.sessions[sid]
	core.Mutex.RUnlock()
	return found, nil
}

func (core *Core) CreateUserAccount(request SignupRequest) error {
	core.Mutex.Lock()
	core.users[request.Login] = User(request)
	core.Mutex.Unlock()
	return nil
}

func (core *Core) FindUserAccount(login string) (User, bool, error) {
	core.Mutex.RLock()
	user, found := core.users[login]
	core.Mutex.RUnlock()
	return user, found, nil
}

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}

func (core *Core) GetCollection(collectionId string) (string, bool, error) {
	core.Mutex.RLock()
	collectionName, found := core.collections[collectionId]
	core.Mutex.RUnlock()
	return collectionName, found, nil
}

func GetFilms() ([]Film, error) {
	films := []Film{}

	films = append(films, Film{"Леди Баг и Супер-Кот: Пробуждение силы", "../../icons/lady-poster.jpg", 7.5, []string{"Комедия", "Приключения", "Фэнтези", "Мелодрама", "Зарубежный"}})
	films = append(films, Film{"Барби", "../../icons/Barbie_2023_poster.jpeg", 6.7, []string{"Комедия", "Приключения", "Фэнтези", "Зарубежный"}})
	films = append(films, Film{"Опенгеймер", "../../icons/Op.jpg", 8.5, []string{"Биография", "Драма", "История", "Ужас", "Зарубежный"}})
	films = append(films, Film{"Слуга народа", "../../icons/Slave_nation.jpg", 0.7, []string{"Комедия", "Зарубежный"}})
	films = append(films, Film{"Черная Роза", "../../icons/Black_rose.jpg", 1.5, []string{"Детектив", "Триллер", "Криминал", "Российский"}})
	films = append(films, Film{"Бесславные ублюдки", "../../icons/bastards.jpg", 8.0, []string{"Боевик", "Военный", "Драма", "Комедия", "Зарубежный"}})
	films = append(films, Film{"Бэтмен: Начало", "../../icons/Batman_Begins.jpg", 7.9, []string{"Боевик", "Фантастика", "Драма", "Приключения", "Зарубежный"}})
	films = append(films, Film{"Криминальное чтиво", "../../icons/criminal.jpeg", 8.6, []string{"Криминал", "Драма", "Зарубежный"}})
	films = append(films, Film{"Терминатор", "/", 8.0, []string{"Боевик", "Фантастика", "Триллер", "Зарубежный"}})

	return films, nil
}

func SortFilms(collectionName string, films []Film) ([]Film, error) {
	sorted := make([]Film, 0, cap(films))

	for _, film := range films {
		for _, genre := range film.Genres {
			if strings.EqualFold(genre, collectionName) {
				sorted = append(sorted, film)
			}
		}
	}

	return sorted, nil
}
