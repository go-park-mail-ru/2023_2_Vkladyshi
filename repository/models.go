package repository

import "database/sql"

type CrewItem struct {
	Id        int
	Name      string
	Birthdate string
	Photo     string
}

type FilmItem struct {
	Id          int
	Title       string
	Info        string
	Poster      string
	ReleaseDate string
	Country     string
	Mpaa        string
}

type UserItem struct {
	Id               int
	name             string
	Birthdate        string
	Photo            string
	Login            string
	Password         string
	RegistrationDate string
}

type ProfessionItem struct {
	Id    int
	Title string
}

type GenreItem struct {
	Id    int
	Title string
}

type CommentItem struct {
	IdUser  int
	IdFilm  int
	Rating  int
	Comment string
}

type PersonInFilmItem struct {
	IdFilm        int
	IdPerson      int
	IdProfession  int
	CharacterName sql.NullString
}
