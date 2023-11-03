package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"gopkg.in/yaml.v2"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/film"
	profile "github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/user"
)

type DbDsnCfg struct {
	User         string `yaml:"user"`
	DbName       string `yaml:"dbname"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Sslmode      string `yaml:"sslmode"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func getPostgres() (*sql.DB, error) {
	dsnConfig := DbDsnCfg{}
	dsnFile, err := os.ReadFile("./configs/db_dsn.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(dsnFile, dsnConfig)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		dsnConfig.User, dsnConfig.DbName, dsnConfig.Password, dsnConfig.Host, dsnConfig.Port, dsnConfig.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dsnConfig.MaxOpenConns)

	return db, nil
}

func main() {
	logFile, _ := os.Create("log.log")
	lg := slog.New(slog.NewJSONHandler(logFile, nil))

	db, err := getPostgres()
	if err != nil {
		lg.Error("failed to connect to db", "err", err.Error())
	}
	core := Core{
		sessions: make(map[string]Session),
		collections: map[string]string{
			"new":       "Новинки",
			"action":    "Боевик",
			"comedy":    "Комедия",
			"ru":        "Российский",
			"eu":        "Зарубежный",
			"war":       "Военный",
			"kids":      "Детский",
			"detective": "Детектив",
			"drama":     "Драма",
			"crime":     "Криминал",
			"melodrama": "Мелодрама",
			"horror":    "Ужас",
		},
		lg:    lg.With("module", "core"),
		Films: &film.RepoPostgre{DB: db},
		Users: &profile.RepoPostgre{DB: db},
	}
	api := API{core: &core, lg: lg.With("module", "api")}

	mx := http.NewServeMux()
	mx.HandleFunc("/signup", api.Signup)
	mx.HandleFunc("/signin", api.Signin)
	mx.HandleFunc("/logout", api.LogoutSession)
	mx.HandleFunc("/authcheck", api.AuthAccept)
	mx.HandleFunc("/api/v1/films", api.Films)
	err = http.ListenAndServe(":8080", mx)
	if err != nil {
		api.lg.Error("ListenAndServe error", "err", err.Error())
	}
}
