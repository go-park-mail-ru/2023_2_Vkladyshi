package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/comment"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/film"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/profile"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"photo", "login"})

	testUser := profile.UserItem{
		Photo: "url1",
		Login: "l1",
	}
	expect := []*profile.UserItem{&testUser}

	for _, item := range expect {
		rows = rows.AddRow(item.Login, item.Photo)
	}

	mock.ExpectQuery("SELECT login, photo FROM profiles WHERE").WithArgs(expect[0].Login, expect[0].Password).WillReturnRows(rows)

	repo := &profile.RepoPostgre{
		DB: db,
	}

	user, foundAccount, err := repo.GetUser(expect[0].Login, expect[0].Password)
	if err != nil {
		t.Errorf("GetUser error: %s", err)
	}
	if !foundAccount {
		t.Errorf("user not found")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(user, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], user)
		return
	}

	mock.
		ExpectQuery("SELECT login, photo FROM profiles WHERE").
		WithArgs(expect[0].Login, expect[0].Password).
		WillReturnError(fmt.Errorf("db_error"))

	_, found, err := repo.GetUser(expect[0].Login, expect[0].Password)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if found {
		t.Errorf("expected not found")
	}
}

func TestFindUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"login"})

	testUser := profile.UserItem{
		Login: "l1",
	}
	expect := []*profile.UserItem{&testUser}

	for _, item := range expect {
		rows = rows.AddRow(item.Login)
	}

	mock.ExpectQuery("SELECT login FROM profiles WHERE").WithArgs(expect[0].Login).WillReturnRows(rows)

	repo := &profile.RepoPostgre{
		DB: db,
	}

	foundAccount, err := repo.FindUser(expect[0].Login)
	if err != nil {
		t.Errorf("GetUser error: %s", err)
	}
	if !foundAccount {
		t.Errorf("user not found")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	mock.
		ExpectQuery("SELECT login FROM profiles WHERE").
		WithArgs(expect[0].Login).
		WillReturnError(fmt.Errorf("db_error"))

	found, err := repo.FindUser(expect[0].Login)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if found {
		t.Errorf("expected not found")
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"login"})

	testUser := profile.UserItem{
		Login:     "l1",
		Password:  "p1",
		Birthdate: "2003-10-08",
		Name:      "n1",
		Email:     "e1",
	}
	expect := []*profile.UserItem{&testUser}

	for _, item := range expect {
		rows = rows.AddRow(item.Login)
	}

	mock.
		ExpectExec("INSERT INTO profiles").
		WithArgs(testUser.Name, testUser.Birthdate, testUser.Login, testUser.Password, testUser.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &profile.RepoPostgre{
		DB: db,
	}

	err = repo.CreateUser(testUser.Login, testUser.Password, testUser.Name, testUser.Birthdate, testUser.Email)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	mock.
		ExpectExec("INSERT INTO profiles").
		WithArgs(testUser.Name, testUser.Birthdate, testUser.Login, testUser.Password, testUser.Email).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.CreateUser(testUser.Login, testUser.Password, testUser.Name, testUser.Birthdate, testUser.Email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}

func TestGetFilmsByGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Id", "Title", "Poster"})

	expect := []film.FilmItem{
		{Id: 1, Title: "t1", Poster: "url1"},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Title, item.Poster)
	}

	mock.ExpectQuery("SELECT film.id, film.title, poster FROM film JOIN").WithArgs("g1", 1, 2).WillReturnRows(rows)

	repo := &film.RepoPostgre{
		DB: db,
	}

	films, err := repo.GetFilmsByGenre("g1", 1, 2)
	if err != nil {
		t.Errorf("GetFilmsByGenre error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(films, expect) {
		t.Errorf("results not match, want %v, have %v", expect, films)
		return
	}

	mock.
		ExpectQuery("SELECT film.id, film.title, poster FROM film JOIN").
		WithArgs("g3", 1, 2).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetFilmsByGenre("g3", 1, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetFilms(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Id", "Title", "Poster"})

	expect := []film.FilmItem{
		{Id: 1, Title: "t1", Poster: "url1"},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Title, item.Poster)
	}

	mock.ExpectQuery("SELECT film.id, film.title, poster FROM film").WithArgs(1, 2).WillReturnRows(rows)

	repo := &film.RepoPostgre{
		DB: db,
	}

	films, err := repo.GetFilms(1, 2)
	if err != nil {
		t.Errorf("GetFilms error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(films, expect) {
		t.Errorf("results not match, want %v, have %v", expect, films)
		return
	}

	mock.
		ExpectQuery("SELECT film.id, film.title, poster FROM film").
		WithArgs(1, 2).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetFilms(1, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetFilm(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Id", "Title", "Info", "Poster", "ReleaseDate", "Country", "Mpaa"})

	expect := []film.FilmItem{
		{Id: 1, Title: "t1", Info: "i1", Poster: "url1", ReleaseDate: "date1", Country: "c1", Mpaa: "12"},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Title, item.Info, item.Poster, item.ReleaseDate, item.Country, item.Mpaa)
	}

	mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)

	repo := &film.RepoPostgre{
		DB: db,
	}

	films, err := repo.GetFilm(1)
	if err != nil {
		t.Errorf("GetFilm error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(films, &expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect, films)
		return
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetFilm(1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetFilmRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Average"})

	expect := 4.2

	rows = rows.AddRow(expect)

	mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)

	repo := &comment.RepoPostgre{
		DB: db,
	}

	rating, err := repo.GetFilmRating(1)
	if err != nil {
		t.Errorf("GetFilm error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if rating != expect {
		t.Errorf("results not match, want %v, have %v", expect, rating)
		return
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetFilmRating(1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetFilmComments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Login", "Rating", "Comment"})

	expect := []comment.CommentItem{
		{Username: "l1", Rating: 4, Comment: "c1"},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.Username, item.Rating, item.Comment)
	}

	mock.ExpectQuery("SELECT").WithArgs(1, 0, 5).WillReturnRows(rows)

	repo := &comment.RepoPostgre{
		DB: db,
	}

	comments, err := repo.GetFilmComments(1, 0, 5)
	if err != nil {
		t.Errorf("GetFilm error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(comments, expect) {
		t.Errorf("results not match, want %v, have %v", expect, comments)
		return
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(1, 0, 5).
		WillReturnError(fmt.Errorf("db_error"))

	comments, err = repo.GetFilmComments(1, 0, 5)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if comments != nil {
		t.Errorf("get comments error, comments should be nil")
	}
}
