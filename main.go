package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mx := http.NewServeMux()

	coreLogFile, _ := os.Create("core_log.log")
	apiLogFile, _ := os.Create("api_log.log")
	core := Core{lg: slog.New(slog.NewJSONHandler(coreLogFile, nil))}
	api := API{core: &core, lg: slog.New(slog.NewJSONHandler(apiLogFile, nil))}

	mx.HandleFunc("/api/v1/films", api.Films)

	http.ListenAndServe(":8080", mx)
}
