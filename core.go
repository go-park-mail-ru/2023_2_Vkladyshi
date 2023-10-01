package main

import (
	"math/rand"
	"log/slog"
	"net/http"
	"time"
	"sync"
)

type Core struct {
	Sessions map[string]Session
	Users    map[string]User.User
	Lg       *slog.Logger
	Mutex    sync.RWMutex
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}
