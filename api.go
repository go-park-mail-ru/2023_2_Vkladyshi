package main

import (
	"net/http"
)

type API struct {
	core *Core
}

type Film struct {
	title    string
	imageURL string
	rating   float64
}

type FilmsResponse struct {
	Status int
	Page   uint64 `json:"current_page"`
	Total  uint64 `json:"total"`
	Films  []Film `json:"films"`
}

func addFilms() []Film {
	films := []Film{}

	films = append(films, Film{"Терминатор", "/static/img/terminator.png", 9.3})
	films = append(films, Film{"Барби", "/static/img/barbie.png", 8.2})
	films = append(films, Film{"Слуга народа", "/static/img/sluga_naroda.png", 0.7})
	films = append(films, Film{"Опенгеймер", "/static/img/oppenheimer.png", 8.7})
	films = append(films, Film{"Черная Роза", "/static/img/black_rose.png", 1.5})

	return films
}

func (a *API) Films(w http.ResponseWriter, r *http.Request) {

}
