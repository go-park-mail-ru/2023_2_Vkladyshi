package comment

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetFilmComments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Login", "Rating", "Comment", "Photo"})

	expect := []CommentItem{
		{Username: "l1", Rating: 4, Comment: "c1", Photo: "p1"},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.Username, item.Rating, item.Comment, item.Photo)
	}

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT profile.login, rating, comment, profile.photo FROM users_comment JOIN profile ON users_comment.id_user = profile.id WHERE id_film = $1 OFFSET $2 LIMIT $3")).
		WithArgs(1, 0, 5).
		WillReturnRows(rows)

	repo := &RepoPostgre{
		db: db,
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

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT profile.login, rating, comment, profile.photo FROM users_comment JOIN profile ON users_comment.id_user = profile.id WHERE id_film = $1 OFFSET $2 LIMIT $3")).
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

func TestAddComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	testComment := CommentItem{
		Username: "l1",
		IdFilm:   1,
		Rating:   1,
		Comment:  "c1",
	}

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO users_comment(id_film, rating, comment, id_user) SELECT $1, $2, $3, profile.id FROM profile WHERE login = $4")).
		WithArgs(testComment.IdFilm, testComment.Rating, testComment.Comment, testComment.Username).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &RepoPostgre{
		db: db,
	}

	err = repo.AddComment(testComment.IdFilm, testComment.Username, testComment.Rating, testComment.Comment)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO users_comment(id_film, rating, comment, id_user) SELECT $1, $2, $3, profile.id FROM profile WHERE login = $4")).
		WithArgs(testComment.IdFilm, testComment.Rating, testComment.Comment, testComment.Username).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.AddComment(testComment.IdFilm, testComment.Username, testComment.Rating, testComment.Comment)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
