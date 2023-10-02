package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type API struct {
	core *Core
	lg   *slog.Logger
}

type Film struct {
	Title    string
	ImageURL string `json:"imagine_url"`
	Rating   float64
	Genres   []string
}

type FilmsResponse struct {
	Page     uint64 `json:"current_page"`
	PageSize int    `json:"page_size"`
	Total    uint64 `json:"total"`
	Films    []Film `json:"films"`
}

func addFilms() []Film {
	films := []Film{}

	films = append(films, Film{"Леди Баг и Супер-Кот: Пробуждение силы", "/static/img/ladybag.png", 7.5, []string{"комедия", "приключения", "фэнтези", "мелодрама"}})
	films = append(films, Film{"Барби", "/static/img/barbie.png", 6.7, []string{"комедия", "приключения", "фэнтези"}})
	films = append(films, Film{"Опенгеймер", "/static/img/oppenheimer.png", 8.5, []string{"биография", "драма", "история"}})
	films = append(films, Film{"Слуга народа", "/static/img/sluga_naroda.png", 0.7, []string{"комедия"}})
	films = append(films, Film{"Черная Роза", "/static/img/black_rose.png", 1.5, []string{"детектив", "триллер", "криминал"}})
	films = append(films, Film{"Бесславные ублюдки", "/static/img/inglourious_basterds.png", 8.0, []string{"боевик", "военный", "драма", "комедия"}})
	films = append(films, Film{"Бэтмен: Начало", "/static/img/batman_begins.png", 7.9, []string{"боевик", "фантастика", "драма", "приключения"}})
	films = append(films, Film{"Криминальное чтиво", "/static/img/pulp_fiction.png", 8.6, []string{"криминал", "драма"}})
	films = append(films, Film{"Терминатор", "/static/img/terminator.png", 8.0, []string{"боевик", "фантастика", "триллер"}})

	return films
}

func (a *API) Films(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
	} else {
		page, _ := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
		pageSize, err := strconv.ParseUint(r.URL.Query().Get("page_size"), 10, 64)
		if err != nil {
			pageSize = 8
		}

		films := addFilms()

		filmsResponse := FilmsResponse{
			Page:     page,
			Total:    uint64(len(films)),
			Films:    films[pageSize*(page-1) : pageSize*page],
			PageSize: int(pageSize),
		}
		response.Body = filmsResponse
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.lg.Error("failed to pack json", "err", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
