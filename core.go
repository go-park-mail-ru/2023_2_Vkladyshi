package main

import (
	"log"
	"math/rand"
	"log/slog"
	"net/http"
	"time"
)

type Core struct {
	Sessions map[string]string
	Users    map[string]User.User
	lg *slog.Logger
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}

func (core *Core) CreateSession(w *http.ResponseWriter, r *http.Request, login string) string {
	SID := RandStringRunes(32)
	core.Sessions[SID] = login
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(*w, cookie)
	return SID
}
