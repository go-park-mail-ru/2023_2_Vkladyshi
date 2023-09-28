package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mx := http.NewServeMux()

	logFile, _ := os.Create("log.log")
	core := Core{lg: slog.New(slog.NewJSONHandler(logFile, nil))}
	api := API{core: &core}

	http.ListenAndServe(":8080", mx)
}
