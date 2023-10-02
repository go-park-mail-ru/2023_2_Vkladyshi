package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logFile, _ := os.Create("log.log")

	core := Core{
		Sessions: make(map[string]Session),
		Users:    make(map[string]User),
	}

	api := API{
		core: &core,
		Lg: slog.New(slog.NewJSONHandler(logFile, nil))}
	}
	mx := http.NewServeMux()
	mx.HandleFunc("/signup", api.Signup)
	mx.HandleFunc("/signin", api.Signin)
	mx.HandleFunc("/films", api.Films)
	mx.HandleFunc("/logout", api.LogoutSession)
	http.ListenAndServe(":8080", mx)
}
