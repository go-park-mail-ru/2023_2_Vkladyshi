package usecase

import (
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/crew"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/film"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/genre"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/profession"
)

type FilmsResponse struct {
	Page           uint64          `json:"current_page"`
	PageSize       uint64          `json:"page_size"`
	CollectionName string          `json:"collection_name"`
	Total          uint64          `json:"total"`
	Films          []film.FilmItem `json:"films"`
}

type FilmResponse struct {
	Film       film.FilmItem     `json:"film"`
	Genres     []genre.GenreItem `json:"genre"`
	Rating     float64           `json:"rating"`
	Number     uint64            `json:"number"`
	Directors  []crew.CrewItem   `json:"directors"`
	Scenarists []crew.CrewItem   `json:"scenarists"`
	Characters []crew.Character  `json:"actors"`
}

type ActorResponse struct {
	Name      string                      `json:"name"`
	Photo     string                      `json:"poster_href"`
	Career    []profession.ProfessionItem `json:"career"`
	BirthDate string                      `json:"birthday"`
	Country   string                      `json:"country"`
	Info      string                      `json:"info_text"`
}
