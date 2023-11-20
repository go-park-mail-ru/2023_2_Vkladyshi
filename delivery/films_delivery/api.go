package films_delivery

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/delivery/requests_responses"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/film"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/usecase/films_usecase"
)

type IApi interface {
	SendResponse(w http.ResponseWriter, response requests_responses.Response)
	Films(w http.ResponseWriter, r *http.Request)
	Film(w http.ResponseWriter, r *http.Request)
	Actor(w http.ResponseWriter, r *http.Request)
}

type API struct {
	core films_usecase.ICore
	lg   *slog.Logger
	mx   *http.ServeMux
}

func GetApi(c *films_usecase.Core, l *slog.Logger) *API {
	api := &API{
		core: c,
		lg:   l.With("module", "api"),
	}
	mx := http.NewServeMux()
	mx.HandleFunc("/api/v1/films", api.Films)
	mx.HandleFunc("/api/v1/film", api.Film)
	mx.HandleFunc("/api/v1/actor", api.Actor)

	api.mx = mx

	return api
}

func (a *API) ListenAndServe() {
	err := http.ListenAndServe(":8081", a.mx)
	if err != nil {
		a.lg.Error("ListenAndServe error", "err", err.Error())
	}
}

func (a *API) SendResponse(w http.ResponseWriter, response requests_responses.Response) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.lg.Error("failed to pack json", "err", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		a.lg.Error("failed to send response", "err", err.Error())
	}
}

func (a *API) Films(w http.ResponseWriter, r *http.Request) {
	response := requests_responses.Response{Status: http.StatusOK, Body: nil}

	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}

	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("page_size"), 10, 64)
	if err != nil {
		pageSize = 8
	}

	genreId, err := strconv.ParseUint(r.URL.Query().Get("collection_id"), 10, 64)
	if err != nil {
		genreId = 0
	}

	var films []film.FilmItem

	if genreId == 0 {
		films, err = a.core.GetFilms(uint64((page-1)*pageSize), pageSize)
	} else {
		films, err = a.core.GetFilmsByGenre(genreId, uint64((page-1)*pageSize), pageSize)
	}
	if err != nil {
		a.lg.Error("Films error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}

	genre, err := a.core.GetGenre(genreId)
	if err != nil {
		a.lg.Error("Films get genre error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	filmsResponse := requests_responses.FilmsResponse{
		Page:           page,
		PageSize:       pageSize,
		Total:          uint64(len(films)),
		CollectionName: genre,
		Films:          films,
	}
	response.Body = filmsResponse

	a.SendResponse(w, response)
}

func (a *API) Film(w http.ResponseWriter, r *http.Request) {
	response := requests_responses.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}

	filmId, err := strconv.ParseUint(r.URL.Query().Get("film_id"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	film, err := a.core.GetFilm(filmId)
	if err != nil {
		a.lg.Error("Film error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	if film.Title == "" {
		response.Status = http.StatusNotFound
		a.SendResponse(w, response)
		return
	}
	genres, err := a.core.GetFilmGenres(filmId)
	if err != nil {
		a.lg.Error("Film error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	rating, number, err := a.core.GetFilmRating(filmId)
	if err != nil {
		a.lg.Error("Film error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	directors, err := a.core.GetFilmDirectors(filmId)
	if err != nil {
		a.lg.Error("Film error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	scenarists, err := a.core.GetFilmScenarists(filmId)
	if err != nil {
		a.lg.Error("Film error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	characters, err := a.core.GetFilmCharacters(filmId)
	if err != nil {
		a.lg.Error("Film error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}

	filmResponse := requests_responses.FilmResponse{
		Film:       *film,
		Genres:     genres,
		Rating:     rating,
		Number:     number,
		Directors:  directors,
		Scenarists: scenarists,
		Characters: characters,
	}
	response.Body = filmResponse

	a.SendResponse(w, response)
}

func (a *API) Actor(w http.ResponseWriter, r *http.Request) {
	response := requests_responses.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}

	actorId, err := strconv.ParseUint(r.URL.Query().Get("actor_id"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	actor, err := a.core.GetActor(actorId)
	if err != nil {
		a.lg.Error("Actor error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	if actor == nil {
		response.Status = http.StatusNotFound
		a.SendResponse(w, response)
		return
	}
	career, err := a.core.GetActorsCareer(actorId)
	if err != nil {
		a.lg.Error("Actor error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}

	actorResponse := requests_responses.ActorResponse{
		Name:      actor.Name,
		Photo:     actor.Photo,
		BirthDate: actor.Birthdate,
		Country:   actor.Country,
		Info:      actor.Info,
		Career:    career,
	}

	response.Body = actorResponse
	a.SendResponse(w, response)
}
