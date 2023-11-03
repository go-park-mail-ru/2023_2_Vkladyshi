package repository

import (
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/film"
)

type Repository interface {
	GetFilmsByGenre(genre string, start int, end int) ([]*film.FilmItem, error)
}
