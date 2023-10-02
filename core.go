package main

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Core struct {
	sessions map[string]Session
	users    map[string]User
	Mutex    sync.RWMutex
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (core *Core) CreateSession(w *http.ResponseWriter, r *http.Request, login string) (string, Session) {
	SID := RandStringRunes(32)

	core.Mutex.Lock()
	core.sessions[SID] = Session{
		Login:     login,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	core.Mutex.Unlock()
	return SID, core.sessions[SID]
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
