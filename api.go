package main

import (
	"encoding/json"
	"log/slog"
	"math"
	"net/http"
	"strconv"
)

type API struct {
	core *Core
	lg   *slog.Logger
}

type Film struct {
	Title    string   `json:"title"`
	ImageURL string   `json:"imagine_url"`
	Rating   float64  `json:"rating"`
	Genres   []string `json:"genres"`
}

type FilmsResponse struct {
	Page     uint64 `json:"current_page"`
	PageSize int    `json:"page_size"`
	Total    uint64 `json:"total"`
	Films    []Film `json:"films"`
}

func (a *API) Films(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
	} else {
		page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
		if err != nil {
			page = 1
		}
		pageSize, err := strconv.ParseUint(r.URL.Query().Get("page_size"), 10, 64)
		if err != nil {
			pageSize = 8
		}

		films := AddFilms()
		if uint64(cap(films)) < page*pageSize {
			page = uint64(math.Ceil(float64(uint64(cap(films)) / pageSize)))
		}
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
