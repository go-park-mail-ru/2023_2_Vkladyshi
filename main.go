package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logFile, _ := os.Create("log.log")

	core := Core{
		sessions: make(map[string]Session),
		users:    make(map[string]User),
	}

	api := API{
		core: &core,
		Lg: slog.New(slog.NewJSONHandler(logFile, nil))},
	}
	mx := http.NewServeMux()
	mx.HandleFunc("/signup", api.Signup)
	mx.HandleFunc("/signin", api.Signin)
	mx.HandleFunc("/logout", api.LogoutSession)
	http.ListenAndServe(":8080", mx)
}
