// Code generated by MockGen. DO NOT EDIT.
// Source: repo_genre.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"
	requests "github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/requests"
	gomock "github.com/golang/mock/gomock"
)

// MockIGenreRepo is a mock of IGenreRepo interface.
type MockIGenreRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIGenreRepoMockRecorder
}

// MockIGenreRepoMockRecorder is the mock recorder for MockIGenreRepo.
type MockIGenreRepoMockRecorder struct {
	mock *MockIGenreRepo
}

// NewMockIGenreRepo creates a new mock instance.
func NewMockIGenreRepo(ctrl *gomock.Controller) *MockIGenreRepo {
	mock := &MockIGenreRepo{ctrl: ctrl}
	mock.recorder = &MockIGenreRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGenreRepo) EXPECT() *MockIGenreRepoMockRecorder {
	return m.recorder
}

// AddFilm mocks base method.
func (m *MockIGenreRepo) AddFilm(genres []uint64, filmId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilm", genres, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilm indicates an expected call of AddFilm.
func (mr *MockIGenreRepoMockRecorder) AddFilm(genres, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilm", reflect.TypeOf((*MockIGenreRepo)(nil).AddFilm), genres, filmId)
}

// GetFilmGenres mocks base method.
func (m *MockIGenreRepo) GetFilmGenres(filmId uint64) ([]models.GenreItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmGenres", filmId)
	ret0, _ := ret[0].([]models.GenreItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmGenres indicates an expected call of GetFilmGenres.
func (mr *MockIGenreRepoMockRecorder) GetFilmGenres(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmGenres", reflect.TypeOf((*MockIGenreRepo)(nil).GetFilmGenres), filmId)
}

// GetGenreById mocks base method.
func (m *MockIGenreRepo) GetGenreById(genreId uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenreById", genreId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenreById indicates an expected call of GetGenreById.
func (mr *MockIGenreRepoMockRecorder) GetGenreById(genreId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenreById", reflect.TypeOf((*MockIGenreRepo)(nil).GetGenreById), genreId)
}

// UsersStatistics mocks base method.
func (m *MockIGenreRepo) UsersStatistics(idUser uint64) ([]requests.UsersStatisticsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersStatistics", idUser)
	ret0, _ := ret[0].([]requests.UsersStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersStatistics indicates an expected call of UsersStatistics.
func (mr *MockIGenreRepoMockRecorder) UsersStatistics(idUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersStatistics", reflect.TypeOf((*MockIGenreRepo)(nil).UsersStatistics), idUser)
}
