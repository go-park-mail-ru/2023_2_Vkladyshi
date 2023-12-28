Таблица crew:
	Содержит в себе немного информации про людей, которые принимали участие в создании фильмов
	(как актёры, актёры озвучки для мультфильмов, так и режиссеры, сценаристы). 
	Фото хранятся в виде url'а. name - ФИО, а не только имя. 
	{ id } -> { birth_date, photo }
	{ name } -> { birth_date, photo }

Таблица profession:
	Таблица, в которой хранятся все названия профессий, которые учавствуют в создании фильмов.
	{ id } -> { title }
	
Таблица film:
	Таблица для хранения фильмов. записываем название фильма, описание, постер, дату выхода,
	страну выпуска и возрастной рейтинг.
	poster - постер к фильму в виде url'a. mpaa - возрастной рейтинг.
	{ id } -> { info, poster, release_date, country, mpaa }
	{ title } -> { info, poster, release_date, country, mpaa }
	
Таблица person_in_film:
	Связующая таблица между професссиями, работниками и фильмами. Также есть поле для имени персонажа, 
	если человек является актёром/актёром озвучки. 
	Это поле может быть NULL, если человек не играл какого-либо пероснажа(например, сценаристы и режиссеры).
	{ id_person, id_film, id_profession } -> { character_name }
	
Таблица genre:
	Таблица, в которой хранятся все названия жанров.
	{ id } -> { title }
	
Таблица films_genre:
	Связующая таблица между фильмами и жанрами. 
	Она нужна, так как у одного фильма может быть сразу несколько жанров.
	
Таблица profiles: 
	Таблица пользователей. 
	В ней хранятся логин, захэшированный пароль, дата регистрации и немного информации о пользователе.
	Фото пользователя хранится в виде url'a. 
	{ id } -> { name, photo, birth_date, password, registration_date }
	{ login } -> { name, photo, birth_date, password, registration_date }
	{ email } -> { name, photo, birth_date, password, registration_date }
	
Таблица users_comment:
	Связующая таблица между фильмами и пользователями. Содержит оценку от пользователя и комментарий. 
	Оценка конкретного пользователя хранится как integer, общая же оценка для фильма будет считаться 
	отдельно и она уже будет float.
	{ id_user, id_film } -> { comment, rating }

Таблица users_favorite_actor:
	Таблица избранных актёров пользователя.

Таблица users_favorite_film:
	Таблица избранных фильмов пользователя.


В каждой таблице кортежи содержат только одно значение каждого из атрибутов, все неключевые атрибуты зависят от потенциальных ключей, но не зависят от неключевых. 