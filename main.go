package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"gopkg.in/yaml.v2"
)

type DbDsnCfg struct {
	User     string `yaml:"user"`
	DbName   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Sslmode  string `yaml:"sslmode"`
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

	dsn := "user=" + dsnConfig.User + "dbname=" + dsnConfig.DbName + "password=" + dsnConfig.Password + "host=" + dsnConfig.Host +
		"port=" + dsnConfig.Port + "sslmode=" + dsnConfig.Sslmode
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)

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
		users:    make(map[string]User),
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
		lg: lg.With("module", "core"),
		Db: db,
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
