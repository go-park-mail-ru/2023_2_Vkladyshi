package main

import "log/slog"

type Core struct {
	lg *slog.Logger
}

func AddFilms() []Film {
	films := []Film{}

	films = append(films, Film{"Леди Баг и Супер-Кот: Пробуждение силы", "/static/img/ladybag.png", 7.5, []string{"комедия", "приключения", "фэнтези", "мелодрама"}})
	films = append(films, Film{"Барби", "/static/img/barbie.png", 6.7, []string{"комедия", "приключения", "фэнтези"}})
	films = append(films, Film{"Опенгеймер", "/static/img/oppenheimer.png", 8.5, []string{"биография", "драма", "история"}})
	films = append(films, Film{"Слуга народа", "/static/img/sluga_naroda.png", 0.7, []string{"комедия"}})
	films = append(films, Film{"Черная Роза", "/static/img/black_rose.png", 1.5, []string{"детектив", "триллер", "криминал"}})
	films = append(films, Film{"Бесславные ублюдки", "/static/img/inglourious_basterds.png", 8.0, []string{"боевик", "военный", "драма", "комедия"}})
	films = append(films, Film{"Бэтмен: Начало", "/static/img/batman_begins.png", 7.9, []string{"боевик", "фантастика", "драма", "приключения"}})
	films = append(films, Film{"Криминальное чтиво", "/static/img/pulp_fiction.png", 8.6, []string{"криминал", "драма"}})
	films = append(films, Film{"Терминатор", "/static/img/terminator.png", 8.0, []string{"боевик", "фантастика", "триллер"}})

	return films
}
