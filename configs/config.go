package configs

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

type DbDsnCfg struct {
	User          string `yaml:"user"`
	DbName        string `yaml:"dbname"`
	Password      string `yaml:"password"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Sslmode       string `yaml:"sslmode"`
	MaxOpenConns  int    `yaml:"max_open_conns"`
	Timer         uint32 `yaml:"timer"`
	Films_db      string `yaml:"postgres"`
	Genres_db     string `yaml:"postgres"`
	Crew_db       string `yaml:"postgres"`
	Profession_db string `yaml:"postgres"`
}

type DbRedisCfg struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DbNumber int    `yaml:"db"`
	Timer    int    `yaml:"timer"`
}

func ReadCsrfRedisConfig() (*DbRedisCfg, error) {
	var path string
	flag.StringVar(&path, "config_path", "../../configs/db_csrf.yaml", "Путь к конфигу")

	csrfConfig := DbRedisCfg{}
	csrfFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(csrfFile, &csrfConfig)
	if err != nil {
		return nil, err
	}

	return &csrfConfig, nil
}

func ReadSessionRedisConfig() (*DbRedisCfg, error) {
	sessionConfig := DbRedisCfg{}
	sessionFile, err := os.ReadFile("../../configs/db_session.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(sessionFile, &sessionConfig)
	if err != nil {
		return nil, err
	}

	return &sessionConfig, nil
}

func ReadConfig() (*DbDsnCfg, error) {
	dsnConfig := DbDsnCfg{}
	dsnFile, err := os.ReadFile("../../configs/db_dsn.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(dsnFile, &dsnConfig)
	if err != nil {
		return nil, err
	}

	return &dsnConfig, nil
}

func ReadFilmConfig() (*DbDsnCfg, error) {
	dsnConfig := DbDsnCfg{}
	dsnFile, err := os.ReadFile("../../configs/db_film_dsn.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(dsnFile, &dsnConfig)
	if err != nil {
		return nil, err
	}

	return &dsnConfig, nil
}
