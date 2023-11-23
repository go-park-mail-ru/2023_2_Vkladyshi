package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/requests"
)

type IApi interface {
	SendResponse(w http.ResponseWriter, response requests.Response)
	LogoutSession(w http.ResponseWriter, r *http.Request)
	AuthAccept(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Comment(w http.ResponseWriter, r *http.Request)
	AddComment(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
}
