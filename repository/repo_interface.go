package repository

import (
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/film"
)

type Repository interface {
	GetFilmsByGenre(genre string, start uint32, end uint32) ([]film.FilmItem, error)
	GetFilms(start uint32, end uint32) ([]film.FilmItem, error)
}
