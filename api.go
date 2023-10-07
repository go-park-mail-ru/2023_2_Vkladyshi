package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"time"
)

type API struct {
	core *Core
	lg   *slog.Logger
}

type Session struct {
	Login     string
	ExpiresAt time.Time
}

type Film struct {
	Title    string   `json:"title"`
	ImageURL string   `json:"poster_href"`
	Rating   float64  `json:"rating"`
	Genres   []string `json:"genres"`
}

type FilmsResponse struct {
	Page           uint64 `json:"current_page"`
	PageSize       uint64 `json:"page_size"`
	CollectionName string `json:"collection_name"`
	Total          uint64 `json:"total"`
	Films          []Film `json:"films"`
}

func (a *API) SendResponse(w http.ResponseWriter, response Response) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.lg.Error("failed to pack json", "err", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (a *API) Films(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}

	if r.Method != http.MethodGet {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}

	collectionId := r.URL.Query().Get("collection_id")
	if collectionId == "" {
		collectionId = "new"
	}

	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("page_size"), 10, 64)
	if err != nil {
		pageSize = 8
	}

	collectionName, found, _ := a.core.GetCollection(collectionId)
	if !found {
		collectionName = "Новинки"
	}

	films, _ := GetFilms()
	if collectionName != "Новинки" {
		films, _ = SortFilms(collectionName, films)
	}

	if uint64(len(films)) < page*pageSize {
		page = uint64(math.Ceil(float64(len(films)) / float64(pageSize)))
	}
	if pageSize > uint64(len(films))-(page-1)*pageSize {
		pageSize = uint64(len(films)) - (page-1)*pageSize
	}
	filmsResponse := FilmsResponse{
		Page:           page,
		PageSize:       pageSize,
		Total:          uint64(len(films)),
		CollectionName: collectionName,
		Films:          films[pageSize*(page-1) : pageSize*page],
	}
	response.Body = filmsResponse

	a.SendResponse(w, response)
}

func (a *API) LogoutSession(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		response.Status = http.StatusUnauthorized
		a.SendResponse(w, response)
		return
	}

	found, _ := a.core.FindActiveSession(session.Value)
	if !found {
		response.Status = http.StatusUnauthorized
		a.SendResponse(w, response)
		return
	} else {
		a.core.KillSession(session.Value)
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
	}

	a.SendResponse(w, response)
}

func (a *API) AuthAccept(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}
	var authorized bool

	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		authorized, _ = a.core.FindActiveSession(session.Value)
	}

	if !authorized {
		response.Status = http.StatusUnauthorized
		a.SendResponse(w, response)
		return
	}

	a.SendResponse(w, response)
}

func (a *API) Signin(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodPost {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
	var request SigninRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	if err = json.Unmarshal(body, &request); err != nil {
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	user, found, _ := a.core.FindUserAccount(request.Login)
	if !found || user.Password != request.Password {
		response.Status = http.StatusUnauthorized
		a.SendResponse(w, response)
		return
	} else {
		sid, session, _ := a.core.CreateSession(user.Login)
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    sid,
			Path:     "/",
			Expires:  session.ExpiresAt,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	a.SendResponse(w, response)
}

func (a *API) Signup(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodPost {
		response.Status = http.StatusMethodNotAllowed
		a.SendResponse(w, response)
		return
	}
	var request SignupRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		response.Status = http.StatusBadRequest
		a.SendResponse(w, response)
		return
	}

	_, found, _ := a.core.FindUserAccount(request.Login)
	if found {
		response.Status = http.StatusConflict
		a.SendResponse(w, response)
		return
	} else {
		a.core.CreateUserAccount(request)
	}

	a.SendResponse(w, response)
}
