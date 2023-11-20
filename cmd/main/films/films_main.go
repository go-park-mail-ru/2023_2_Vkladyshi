package films_main

import (
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/delivery/films_delivery"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/usecase/films_usecase"
)

func main() {
	logFile, _ := os.Create("log.log")
	lg := slog.New(slog.NewJSONHandler(logFile, nil))

	config, err := configs.ReadConfig()
	if err != nil {
		lg.Error("read config error", "err", err.Error())
		return
	}

	core, err := films_usecase.GetCore(*config, lg)
	if err != nil {
		lg.Error("cant create core")
		return
	}
	api := films_delivery.GetApi(core, lg)

	api.ListenAndServe()
}
