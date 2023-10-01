package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logFile, _ := os.Create("log.log")

	core := Core.Core{
		Lg:       slog.New(slog.NewJSONHandler(logFile, nil)),
		Sessions: make(map[string]Core.Session),
		Users:    make(map[string]User.User),
	}

	api := Api.API{Core: &core}
	mx := http.NewServeMux()
	mx.HandleFunc("/signup", api.Signup)
	mx.HandleFunc("/signin", api.Signin)
	mx.HandleFunc("/films", api.Films)
	mx.HandleFunc("/logout", api.LogoutSession)
	http.ListenAndServe(":8080", mx)
}

