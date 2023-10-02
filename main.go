package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mx := http.NewServeMux()

	logFile, _ := os.Create("log.log")
	lg := slog.New(slog.NewJSONHandler(logFile, nil))

	core := Core{lg: lg.With("module", "core")}
	api := API{core: &core, lg: lg.With("module", "api")}

	mx.HandleFunc("/api/v1/films", api.Films)

	http.ListenAndServe(":8080", mx)
}
