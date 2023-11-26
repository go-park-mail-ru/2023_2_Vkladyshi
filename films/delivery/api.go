package delivery

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/films/usecase"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/requests"
)

type API struct {
	core usecase.ICore
	lg   *slog.Logger
	mx   *http.ServeMux
}

func GetApi(c *usecase.Core, l *slog.Logger) *API {
	api := &API{
		core: c,
		lg:   l.With("module", "api"),
	}
	mx := http.NewServeMux()
	mx.HandleFunc("/api/v1/films", api.Films)
	mx.HandleFunc("/api/v1/film", api.Film)
	mx.HandleFunc("/api/v1/actor", api.Actor)
	mx.HandleFunc("/api/v1/favorite/films", api.FavoriteFilms)
	mx.HandleFunc("/api/v1/favorite/film/add", api.FavoriteFilmsAdd)
	mx.HandleFunc("/api/v1/favorite/film/remove", api.FavoriteFilmsRemove)
	mx.HandleFunc("/api/v1/favorite/actors", api.FavoriteActors)
	mx.HandleFunc("/api/v1/favorite/actor/add", api.FavoriteActorsAdd)
	mx.HandleFunc("/api/v1/favorite/actor/remove", api.FavoriteActorsRemove)

	api.mx = mx

	return api
}

func (a *API) ListenAndServe() {
	err := http.ListenAndServe(":8081", a.mx)
	if err != nil {
		a.lg.Error("ListenAndServe error", "err", err.Error())
	}
}

func (a *API) SendResponse(w http.ResponseWriter, response requests.Response) {
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
	response := requests.Response{Status: http.StatusOK, Body: nil}

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

	var films []models.FilmItem

	films, genre, err := a.core.GetFilmsAndGenreTitle(genreId, uint64((page-1)*pageSize), pageSize)
	if err != nil {
		a.lg.Error("get films error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}

	filmsResponse := requests.FilmsResponse{
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
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}

	filmId, err := strconv.ParseUint(r.URL.Query().Get("film_id"), 10, 64)
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			response.Status = http.StatusNotFound
			a.SendResponse(w, response)
			return
		}
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	film, err := a.core.GetFilmInfo(filmId)
	if err != nil {
		a.lg.Error("Film error", "err", err.Error())
		response.Status = http.StatusInternalServerError
		a.SendResponse(w, response)
		return
	}
	if film.Film.Title == "" {
		response.Status = http.StatusNotFound
		a.SendResponse(w, response)
		return
	}

	response.Body = film

	a.SendResponse(w, response)
}

func (a *API) Actor(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}

	actorId, err := strconv.ParseUint(r.URL.Query().Get("actor_id"), 10, 64)
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			response.Status = http.StatusNotFound
			a.SendResponse(w, response)
			return
		}
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	actor, err := a.core.GetActorInfo(actorId)
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

	response.Body = actor
	a.SendResponse(w, response)
}

func (a *API) FindFilm(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}

	/*title := r.URL.Query().Get("title")
	dateFrom := r.URL.Query().Get("actor_id")
	dateTo := r.URL.Query().Get("actor_id")
	ratingFrom := r.URL.Query().Get("actor_id")
	ratingTo := r.URL.Query().Get("actor_id")
	mpaa := r.URL.Query().Get("actor_id")
	genres := r.URL.Query().Get("actor_id")
	actors := r.URL.Query().Get("actor_id")

	a.core.SearchFilm(title, dateFrom, dateTo, ratingFrom, ratingTo)*/
}

func (a *API) FavoriteFilmsAdd(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
}

func (a *API) FavoriteFilmsRemove(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
}

func (a *API) FavoriteFilms(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
}

func (a *API) FavoriteActorsAdd(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
}

func (a *API) FavoriteActorsRemove(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
}

func (a *API) FavoriteActors(w http.ResponseWriter, r *http.Request) {
	response := requests.Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
}
