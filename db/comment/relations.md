Таблица users_comment:
	Таблица для хранение комментов к фильму. Содержит оценку от пользователя и комментарий, а также id пользователя и id фильма. 
	Оценка конкретного пользователя хранится как integer, общая же оценка для фильма будет считаться 
	отдельно и она уже будет float.
	{ id_user, id_film } -> { comment, rating }