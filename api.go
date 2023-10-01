package main

import (
	"encoding/json"
	"io"
	"net/http"
	"requests"
	"strconv"
	"time"
)

type API struct {
	Core *Core.Core
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

	_, found := a.Core.Sessions[session.Value]

	if !found {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	delete(a.Core.Sessions, session.Value)
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (a *API) CreateSession(w *http.ResponseWriter, r *http.Request, email string) string {
	SID := Core.RandStringRunes(32)

	a.Core.Mutex.Lock()
	a.Core.Sessions[SID] = Core.Session{
		Email:     email,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	a.Core.Mutex.Unlock()

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Path:     "/",
		Expires:  a.Core.Sessions[SID].ExpiresAt,
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

	var request Request.SigninRequest
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

	a.Core.Mutex.RLock()
	user, found := a.Core.Users[request.Email]
	a.Core.Mutex.RUnlock()

	if !found || user.Password != request.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var answer requests.Response
	answer.Status = http.StatusOK
	answer.Body = a.CreateSession(&w, r, user.Email)

	bytes, err := json.Marshal(answer)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
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

	var request Request.SignupRequest
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

	a.Core.Mutex.RLock()
	_, found := a.Core.Users[request.Login]
	a.Core.Mutex.RUnlock()

	if found {
		w.WriteHeader(http.StatusConflict)
		return
	}

	a.Core.Mutex.Lock()
	a.Core.Users[request.Login] = User.User{Login: request.Login, Email: request.Email, Password: request.Password}
	a.Core.Mutex.Unlock()

	var answer requests.Response
	answer.Status = http.StatusOK
	bytes, err := json.Marshal(answer)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}


