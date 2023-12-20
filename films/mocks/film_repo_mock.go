// Code generated by MockGen. DO NOT EDIT.
// Source: repo_film.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIFilmsRepo is a mock of IFilmsRepo interface.
type MockIFilmsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIFilmsRepoMockRecorder
}

// MockIFilmsRepoMockRecorder is the mock recorder for MockIFilmsRepo.
type MockIFilmsRepoMockRecorder struct {
	mock *MockIFilmsRepo
}

// NewMockIFilmsRepo creates a new mock instance.
func NewMockIFilmsRepo(ctrl *gomock.Controller) *MockIFilmsRepo {
	mock := &MockIFilmsRepo{ctrl: ctrl}
	mock.recorder = &MockIFilmsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFilmsRepo) EXPECT() *MockIFilmsRepoMockRecorder {
	return m.recorder
}

// AddFavoriteFilm mocks base method.
func (m *MockIFilmsRepo) AddFavoriteFilm(userId, filmId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoriteFilm", userId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoriteFilm indicates an expected call of AddFavoriteFilm.
func (mr *MockIFilmsRepoMockRecorder) AddFavoriteFilm(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoriteFilm", reflect.TypeOf((*MockIFilmsRepo)(nil).AddFavoriteFilm), userId, filmId)
}

// AddFilm mocks base method.
func (m *MockIFilmsRepo) AddFilm(film models.FilmItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilm", film)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilm indicates an expected call of AddFilm.
func (mr *MockIFilmsRepoMockRecorder) AddFilm(film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilm", reflect.TypeOf((*MockIFilmsRepo)(nil).AddFilm), film)
}

// AddRating mocks base method.
func (m *MockIFilmsRepo) AddRating(filmId, userId uint64, rating uint16) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRating", filmId, userId, rating)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRating indicates an expected call of AddRating.
func (mr *MockIFilmsRepoMockRecorder) AddRating(filmId, userId, rating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRating", reflect.TypeOf((*MockIFilmsRepo)(nil).AddRating), filmId, userId, rating)
}

// CheckFilm mocks base method.
func (m *MockIFilmsRepo) CheckFilm(userId, filmId uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFilm", userId, filmId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckFilm indicates an expected call of CheckFilm.
func (mr *MockIFilmsRepoMockRecorder) CheckFilm(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFilm", reflect.TypeOf((*MockIFilmsRepo)(nil).CheckFilm), userId, filmId)
}

// FindFilm mocks base method.
func (m *MockIFilmsRepo) FindFilm(title, dateFrom, dateTo string, ratingFrom, ratingTo float32, mpaa string, genres []uint32, actors []string) ([]models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilm", title, dateFrom, dateTo, ratingFrom, ratingTo, mpaa, genres, actors)
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFilm indicates an expected call of FindFilm.
func (mr *MockIFilmsRepoMockRecorder) FindFilm(title, dateFrom, dateTo, ratingFrom, ratingTo, mpaa, genres, actors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilm", reflect.TypeOf((*MockIFilmsRepo)(nil).FindFilm), title, dateFrom, dateTo, ratingFrom, ratingTo, mpaa, genres, actors)
}

// GetFavoriteFilms mocks base method.
func (m *MockIFilmsRepo) GetFavoriteFilms(userId, start, end uint64) ([]models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteFilms", userId, start, end)
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteFilms indicates an expected call of GetFavoriteFilms.
func (mr *MockIFilmsRepoMockRecorder) GetFavoriteFilms(userId, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteFilms", reflect.TypeOf((*MockIFilmsRepo)(nil).GetFavoriteFilms), userId, start, end)
}

// GetFilm mocks base method.
func (m *MockIFilmsRepo) GetFilm(filmId uint64) (*models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilm", filmId)
	ret0, _ := ret[0].(*models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilm indicates an expected call of GetFilm.
func (mr *MockIFilmsRepoMockRecorder) GetFilm(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilm", reflect.TypeOf((*MockIFilmsRepo)(nil).GetFilm), filmId)
}

// GetFilmId mocks base method.
func (m *MockIFilmsRepo) GetFilmId(title string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmId", title)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmId indicates an expected call of GetFilmId.
func (mr *MockIFilmsRepoMockRecorder) GetFilmId(title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmId", reflect.TypeOf((*MockIFilmsRepo)(nil).GetFilmId), title)
}

// GetFilmRating mocks base method.
func (m *MockIFilmsRepo) GetFilmRating(filmId uint64) (float64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmRating", filmId)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFilmRating indicates an expected call of GetFilmRating.
func (mr *MockIFilmsRepoMockRecorder) GetFilmRating(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmRating", reflect.TypeOf((*MockIFilmsRepo)(nil).GetFilmRating), filmId)
}

// GetFilms mocks base method.
func (m *MockIFilmsRepo) GetFilms(start, end uint64) ([]models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilms", start, end)
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilms indicates an expected call of GetFilms.
func (mr *MockIFilmsRepoMockRecorder) GetFilms(start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilms", reflect.TypeOf((*MockIFilmsRepo)(nil).GetFilms), start, end)
}

// GetFilmsByGenre mocks base method.
func (m *MockIFilmsRepo) GetFilmsByGenre(genre, start, end uint64) ([]models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsByGenre", genre, start, end)
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmsByGenre indicates an expected call of GetFilmsByGenre.
func (mr *MockIFilmsRepoMockRecorder) GetFilmsByGenre(genre, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsByGenre", reflect.TypeOf((*MockIFilmsRepo)(nil).GetFilmsByGenre), genre, start, end)
}

// HasUsersRating mocks base method.
func (m *MockIFilmsRepo) HasUsersRating(userId, filmId uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUsersRating", userId, filmId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasUsersRating indicates an expected call of HasUsersRating.
func (mr *MockIFilmsRepoMockRecorder) HasUsersRating(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUsersRating", reflect.TypeOf((*MockIFilmsRepo)(nil).HasUsersRating), userId, filmId)
}

// RemoveFavoriteFilm mocks base method.
func (m *MockIFilmsRepo) RemoveFavoriteFilm(userId, filmId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFavoriteFilm", userId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFavoriteFilm indicates an expected call of RemoveFavoriteFilm.
func (mr *MockIFilmsRepoMockRecorder) RemoveFavoriteFilm(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFavoriteFilm", reflect.TypeOf((*MockIFilmsRepo)(nil).RemoveFavoriteFilm), userId, filmId)
}
