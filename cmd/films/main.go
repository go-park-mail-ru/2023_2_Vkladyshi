package main

import (
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/delivery"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/usecase"
)

func main() {
	logFile, _ := os.Create("film_log.log")
	lg := slog.New(slog.NewJSONHandler(logFile, nil))

	config, err := configs.ReadFilmConfig()
	if err != nil {
		lg.Error("read config error", "err", err.Error())
		return
	}

	core, err := usecase.GetCore(config, lg)
	if err != nil {
		lg.Error("cant create core")
		return
	}
	api := delivery.GetApi(core, lg)

	api.ListenAndServe()
}
