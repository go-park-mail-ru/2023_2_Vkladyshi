package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mx := http.NewServeMux()

	logFile, _ := os.Create("log.log")
	core := Core.Core{Lg: log.New(logFile, "", log.LstdFlags)}
	api := Api.API{Core: &core}
	api.Core.Sessions = make(map[string]string)
	api.Core.Users = make(map[string]User.User)
	mx.HandleFunc("/signup", api.Signup)
	mx.HandleFunc("/signin", api.Signin)
	mx.HandleFunc("/films", api.Films)
	http.ListenAndServe(":8080", mx)
}
