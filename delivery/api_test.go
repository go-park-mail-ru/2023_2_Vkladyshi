package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/delivery/requests_responses"
)

func TestSignupGet(t *testing.T) {
	h := httptest.NewRequest(http.MethodGet, "/signup", nil)
	w := httptest.NewRecorder()

	api := API{}
	api.Signup(w, h)
	var response requests_responses.Response

	body, _ := io.ReadAll(w.Body)
	err := json.Unmarshal(body, &response)
	if err != nil {
		t.Error("cant unmarshal jsone")
	}

	if response.Status != http.StatusMethodNotAllowed {
		t.Errorf("got incorrect status")
	}
}

func TestSigninGet(t *testing.T) {
	h := httptest.NewRequest(http.MethodGet, "/signin", nil)
	w := httptest.NewRecorder()

	api := API{}
	api.Signin(w, h)
	var response requests_responses.Response

	body, _ := io.ReadAll(w.Body)
	err := json.Unmarshal(body, &response)
	if err != nil {
		t.Error("cant unmarshal jsone")
	}

	if response.Status != http.StatusMethodNotAllowed {
		t.Errorf("got incorrect status")
	}
}
