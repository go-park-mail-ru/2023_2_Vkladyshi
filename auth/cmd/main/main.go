package main

import (
	pb "auth/auth"
	"auth/configs"
	delivery_auth "auth/delivery"
	server "auth/grpc_server"
	"auth/usecase"
	"log/slog"
	"os"

	"google.golang.org/grpc"
)

func main() {
	logFile, _ := os.Create("log.log")
	lg := slog.New(slog.NewJSONHandler(logFile, nil))

	config, err := configs.ReadConfig()
	if err != nil {
		lg.Error("read config error", "err", err.Error())
		return
	}

	configCsrf, err := configs.ReadCsrfRedisConfig()
	if err != nil {
		lg.Error("read config error", "err", err.Error())
		return
	}

	configSession, err := configs.ReadSessionRedisConfig()
	if err != nil {
		lg.Error("read config error", "err", err.Error())
		return
	}

	core, err := usecase.GetCore(config, *configCsrf, *configSession, lg)
	if err != nil {
		lg.Error("cant create core")
		return
	}
	api := delivery_auth.GetApi(core, lg)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server.Server{})

	api.ListenAndServe()
}
