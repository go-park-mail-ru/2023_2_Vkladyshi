package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		a.core.lg.Debug(err.Error(), "can't find page")
	}

	w.Header().Set("Content-Type", "application/json")

	films := addFilms()

	response := FilmsResponse{
		Status: 200,
		Page:   page,
		Total:  uint64(len(films)),
		Films:  films,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		response.Status = 500
		a.core.lg.Debug(err.Error(), "json packing err")
	}

	w.Write(jsonResponse)
}
