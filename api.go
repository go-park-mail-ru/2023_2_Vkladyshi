package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"requests"
	"strconv"
)

var films = []Film.Film{}

type API struct {
	Core *Core.Core
}

func (a *API) Films(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
	w.Header().Set("Content-Type", "application/json")
	resp := Film.FilmsResponse{
		Page:  page,
		Total: uint64(len(films)),
		Films: films,
	}
	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

func (a *API) Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request Request.SigninRequest
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &request)
	user, found := a.Core.Users[request.Login]

	if !found || user.Password != request.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var answer requests.Answer
	answer.Status = http.StatusOK
	answer.Body = append(answer.Body, a.Core.CreateSession(&w, r, user.Login))
	bytes, _ := json.Marshal(answer)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (a *API) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request Request.SignupRequest
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &request)
	fmt.Println(request.Login, " ", request.Password, " ", request.Email)
	_, found := a.Core.Users[request.Login]

	if found {
		w.WriteHeader(http.StatusFound)
		return
	}

	a.Core.Users[request.Login] = User.User{Login: request.Login, Email: request.Email, Password: request.Password}
	var answer requests.Answer
	answer.Status = http.StatusOK
	bytes, _ := json.Marshal(answer)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
