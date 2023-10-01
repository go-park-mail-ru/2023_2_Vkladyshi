package main

type SignupRequest struct {
	Login    string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Login    string `json:"username"`
	Password string `json:"password"`
}

type Answer struct {
	Status int
	Body   []interface{}
}
