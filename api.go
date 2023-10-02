package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type API struct {
	core *Core
	Lg   *slog.Logger
}

type Session struct {
	Email     string
	ExpiresAt time.Time
}

func (a *API) LogoutSession(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, found := a.core.Sessions[session.Value]

	if !found {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	delete(a.core.Sessions, session.Value)
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (a *API) CreateSession(w *http.ResponseWriter, r *http.Request, email string) string {
	SID := RandStringRunes(32)

	a.core.Mutex.Lock()
	a.core.Sessions[SID] = Session{
		Email:     email,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	a.core.Mutex.Unlock()

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Path:     "/",
		Expires:  a.core.Sessions[SID].ExpiresAt,
		HttpOnly: true,
	}
	http.SetCookie(*w, cookie)
	return SID
}

func (a *API) Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request SigninRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a.core.Mutex.RLock()
	user, found := a.core.Users[request.Email]
	a.core.Mutex.RUnlock()

	if !found || user.Password != request.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var response Response
	response.Status = http.StatusOK
	response.Body = a.CreateSession(&w, r, user.Email)

	bytes, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (a *API) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request SignupRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a.core.Mutex.RLock()
	_, found := a.core.Users[request.Login]
	a.core.Mutex.RUnlock()

	if found {
		w.WriteHeader(http.StatusConflict)
		return
	}

	a.core.Mutex.Lock()
	a.core.Users[request.Login] = User{Login: request.Login, Email: request.Email, Password: request.Password}
	a.core.Mutex.Unlock()

	var response Response
	response.Status = http.StatusOK
	bytes, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
