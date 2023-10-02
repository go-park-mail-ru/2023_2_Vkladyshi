package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type API struct {
	core *Core
	lg   *slog.Logger
}

type Session struct {
	Email     string
	ExpiresAt time.Time
}

func (a *API) LogoutSession(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		response.Status = http.StatusUnauthorized
	}

	found := a.core.FindActiveSession(session.Value)
	if !found {
		response.Status = http.StatusUnauthorized
	} else {
		a.core.KillSession(session.Value)
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
	}

	answer, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(answer)
}

func (a *API) Signin(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodPost {
		response.Status = http.StatusMethodNotAllowed
	} else {
		var request SigninRequest

		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Status = http.StatusBadRequest
		}

		if err = json.Unmarshal(body, &request); err != nil {
			response.Status = http.StatusBadRequest
		}

		user, found := a.core.FindUserAccount(request.Login)
		if !found || user.Password != request.Password {
			response.Status = http.StatusUnauthorized
		} else {
			sid, session := a.core.CreateSession(&w, r, user.Email)
			cookie := &http.Cookie{
				Name:     "session_id",
				Value:    sid,
				Path:     "/",
				Expires:  session.ExpiresAt,
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		}
	}

	answer, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(answer)
}

func (a *API) Signup(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: http.StatusOK, Body: nil}
	if r.Method != http.MethodPost {
		response.Status = http.StatusMethodNotAllowed
	} else {
		var request SignupRequest

		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Status = http.StatusBadRequest
		}

		err = json.Unmarshal(body, &request)
		if err != nil {
			response.Status = http.StatusBadRequest
		}

		_, found := a.core.FindUserAccount(request.Login)
		if found {
			response.Status = http.StatusConflict
		} else {
			a.core.CreateUserAccount(request)
		}
	}

	answer, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(answer)
}
