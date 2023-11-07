package profession

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/configs"
)

type IProfessionRepo interface {
	PingDb() error
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
	return &postgreDb
}

func (repo *RepoPostgre) PingDb() error {
	err := repo.DB.Ping()
	if err != nil {
		return err
	}
	return nil
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
