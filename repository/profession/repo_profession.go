package profession

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
)

type IProfessionRepo interface {
	GetActorsProfessions(actorId uint64) ([]ProfessionItem, error)
}

type RepoPostgre struct {
	DB *sql.DB
}

func GetProfessionRepo(config configs.DbDsnCfg, lg *slog.Logger) *RepoPostgre {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.DbName, config.Password, config.Host, config.Port, config.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		lg.Error("sql open error", "err", err.Error())
		return nil
	}
	err = db.Ping()
	if err != nil {
		lg.Error("sql ping error", "err", err.Error())
		return nil
	}
	db.SetMaxOpenConns(config.MaxOpenConns)

	postgreDb := RepoPostgre{DB: db}

	go postgreDb.pingDb(config.Timer, lg)
	return &postgreDb
}

func (repo *RepoPostgre) pingDb(timer uint32, lg *slog.Logger) {
	err := repo.DB.Ping()
	if err != nil {
		lg.Error("Repo Profession db ping error", "err", err.Error())
	}

	time.Sleep(time.Duration(timer) * time.Second)
}

func (repo *RepoPostgre) GetActorsProfessions(actorId uint64) ([]ProfessionItem, error) {
	professions := []ProfessionItem{}

	rows, err := repo.DB.Query(
		"SELECT DISTINCT title FROM profession "+
			"JOIN person_in_film ON profession.id = person_in_film.id_profession "+
			"WHERE id_person = $1", actorId)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := ProfessionItem{}
		err := rows.Scan(&post.Title)
		if err != nil {
			return nil, err
		}
		professions = append(professions, post)
	}

	return professions, nil
}