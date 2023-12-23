// Code generated by MockGen. DO NOT EDIT.
// Source: core.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	slog "log/slog"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/models"
	requests "github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/requests"
	gomock "github.com/golang/mock/gomock"
)

// MockICore is a mock of ICore interface.
type MockICore struct {
	ctrl     *gomock.Controller
	recorder *MockICoreMockRecorder
}

// MockICoreMockRecorder is the mock recorder for MockICore.
type MockICoreMockRecorder struct {
	mock *MockICore
}

// NewMockICore creates a new mock instance.
func NewMockICore(ctrl *gomock.Controller) *MockICore {
	mock := &MockICore{ctrl: ctrl}
	mock.recorder = &MockICoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICore) EXPECT() *MockICoreMockRecorder {
	return m.recorder
}

// AddFilm mocks base method.
func (m *MockICore) AddFilm(film models.FilmItem, genres, actors []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilm", film, genres, actors)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilm indicates an expected call of AddFilm.
func (mr *MockICoreMockRecorder) AddFilm(film, genres, actors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilm", reflect.TypeOf((*MockICore)(nil).AddFilm), film, genres, actors)
}

// AddNearFilm mocks base method.
func (m *MockICore) AddNearFilm(ctx context.Context, active models.NearFilm, lg *slog.Logger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNearFilm", ctx, active, lg)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNearFilm indicates an expected call of AddNearFilm.
func (mr *MockICoreMockRecorder) AddNearFilm(ctx, active, lg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNearFilm", reflect.TypeOf((*MockICore)(nil).AddNearFilm), ctx, active, lg)
}

// AddRating mocks base method.
func (m *MockICore) AddRating(filmId, userId uint64, rating uint16) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRating", filmId, userId, rating)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRating indicates an expected call of AddRating.
func (mr *MockICoreMockRecorder) AddRating(filmId, userId, rating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRating", reflect.TypeOf((*MockICore)(nil).AddRating), filmId, userId, rating)
}

// DeleteRating mocks base method.
func (m *MockICore) DeleteRating(idUser, idFilm uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRating", idUser, idFilm)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRating indicates an expected call of DeleteRating.
func (mr *MockICoreMockRecorder) DeleteRating(idUser, idFilm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRating", reflect.TypeOf((*MockICore)(nil).DeleteRating), idUser, idFilm)
}

// FavoriteActors mocks base method.
func (m *MockICore) FavoriteActors(userId, start, end uint64) ([]models.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteActors", userId, start, end)
	ret0, _ := ret[0].([]models.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FavoriteActors indicates an expected call of FavoriteActors.
func (mr *MockICoreMockRecorder) FavoriteActors(userId, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteActors", reflect.TypeOf((*MockICore)(nil).FavoriteActors), userId, start, end)
}

// FavoriteActorsAdd mocks base method.
func (m *MockICore) FavoriteActorsAdd(userId, filmId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteActorsAdd", userId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// FavoriteActorsAdd indicates an expected call of FavoriteActorsAdd.
func (mr *MockICoreMockRecorder) FavoriteActorsAdd(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteActorsAdd", reflect.TypeOf((*MockICore)(nil).FavoriteActorsAdd), userId, filmId)
}

// FavoriteActorsRemove mocks base method.
func (m *MockICore) FavoriteActorsRemove(userId, filmId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteActorsRemove", userId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// FavoriteActorsRemove indicates an expected call of FavoriteActorsRemove.
func (mr *MockICoreMockRecorder) FavoriteActorsRemove(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteActorsRemove", reflect.TypeOf((*MockICore)(nil).FavoriteActorsRemove), userId, filmId)
}

// FavoriteFilms mocks base method.
func (m *MockICore) FavoriteFilms(userId, start, end uint64) ([]models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteFilms", userId, start, end)
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FavoriteFilms indicates an expected call of FavoriteFilms.
func (mr *MockICoreMockRecorder) FavoriteFilms(userId, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteFilms", reflect.TypeOf((*MockICore)(nil).FavoriteFilms), userId, start, end)
}

// FavoriteFilmsAdd mocks base method.
func (m *MockICore) FavoriteFilmsAdd(userId, filmId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteFilmsAdd", userId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// FavoriteFilmsAdd indicates an expected call of FavoriteFilmsAdd.
func (mr *MockICoreMockRecorder) FavoriteFilmsAdd(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteFilmsAdd", reflect.TypeOf((*MockICore)(nil).FavoriteFilmsAdd), userId, filmId)
}

// FavoriteFilmsRemove mocks base method.
func (m *MockICore) FavoriteFilmsRemove(userId, filmId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteFilmsRemove", userId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// FavoriteFilmsRemove indicates an expected call of FavoriteFilmsRemove.
func (mr *MockICoreMockRecorder) FavoriteFilmsRemove(userId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteFilmsRemove", reflect.TypeOf((*MockICore)(nil).FavoriteFilmsRemove), userId, filmId)
}

// FindActor mocks base method.
func (m *MockICore) FindActor(name, birthDate string, films, career []string, country string, first, limit uint64) ([]models.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActor", name, birthDate, films, career, country, first, limit)
	ret0, _ := ret[0].([]models.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActor indicates an expected call of FindActor.
func (mr *MockICoreMockRecorder) FindActor(name, birthDate, films, career, country, first, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActor", reflect.TypeOf((*MockICore)(nil).FindActor), name, birthDate, films, career, country, first, limit)
}

// FindFilm mocks base method.
func (m *MockICore) FindFilm(title, dateFrom, dateTo string, ratingFrom, ratingTo float32, mpaa string, genres []uint32, actors []string, first, limit uint64) ([]models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilm", title, dateFrom, dateTo, ratingFrom, ratingTo, mpaa, genres, actors, first, limit)
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFilm indicates an expected call of FindFilm.
func (mr *MockICoreMockRecorder) FindFilm(title, dateFrom, dateTo, ratingFrom, ratingTo, mpaa, genres, actors, first, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilm", reflect.TypeOf((*MockICore)(nil).FindFilm), title, dateFrom, dateTo, ratingFrom, ratingTo, mpaa, genres, actors, first, limit)
}

// GetActorInfo mocks base method.
func (m *MockICore) GetActorInfo(actorId uint64) (*requests.ActorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorInfo", actorId)
	ret0, _ := ret[0].(*requests.ActorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorInfo indicates an expected call of GetActorInfo.
func (mr *MockICoreMockRecorder) GetActorInfo(actorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorInfo", reflect.TypeOf((*MockICore)(nil).GetActorInfo), actorId)
}

// GetActorsCareer mocks base method.
func (m *MockICore) GetActorsCareer(actorId uint64) ([]models.ProfessionItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsCareer", actorId)
	ret0, _ := ret[0].([]models.ProfessionItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorsCareer indicates an expected call of GetActorsCareer.
func (mr *MockICoreMockRecorder) GetActorsCareer(actorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsCareer", reflect.TypeOf((*MockICore)(nil).GetActorsCareer), actorId)
}

// GetCalendar mocks base method.
func (m *MockICore) GetCalendar() (*requests.CalendarResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCalendar")
	ret0, _ := ret[0].(*requests.CalendarResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCalendar indicates an expected call of GetCalendar.
func (mr *MockICoreMockRecorder) GetCalendar() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCalendar", reflect.TypeOf((*MockICore)(nil).GetCalendar))
}

// GetFilmInfo mocks base method.
func (m *MockICore) GetFilmInfo(filmId uint64) (*requests.FilmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmInfo", filmId)
	ret0, _ := ret[0].(*requests.FilmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmInfo indicates an expected call of GetFilmInfo.
func (mr *MockICoreMockRecorder) GetFilmInfo(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmInfo", reflect.TypeOf((*MockICore)(nil).GetFilmInfo), filmId)
}

// GetFilmsAndGenreTitle mocks base method.
func (m *MockICore) GetFilmsAndGenreTitle(genreId, start, end uint64) ([]models.FilmItem, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsAndGenreTitle", genreId, start, end)
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFilmsAndGenreTitle indicates an expected call of GetFilmsAndGenreTitle.
func (mr *MockICoreMockRecorder) GetFilmsAndGenreTitle(genreId, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsAndGenreTitle", reflect.TypeOf((*MockICore)(nil).GetFilmsAndGenreTitle), genreId, start, end)
}

// GetGenre mocks base method.
func (m *MockICore) GetGenre(genreId uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenre", genreId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenre indicates an expected call of GetGenre.
func (mr *MockICoreMockRecorder) GetGenre(genreId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenre", reflect.TypeOf((*MockICore)(nil).GetGenre), genreId)
}

// GetNearFilms mocks base method.
func (m *MockICore) GetNearFilms(ctx context.Context, userId uint64, lg *slog.Logger) ([]models.NearFilm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearFilms", ctx, userId, lg)
	ret0, _ := ret[0].([]models.NearFilm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearFilms indicates an expected call of GetNearFilms.
func (mr *MockICoreMockRecorder) GetNearFilms(ctx, userId, lg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearFilms", reflect.TypeOf((*MockICore)(nil).GetNearFilms), ctx, userId, lg)
}

// GetUserId mocks base method.
func (m *MockICore) GetUserId(ctx context.Context, sid string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserId", ctx, sid)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserId indicates an expected call of GetUserId.
func (mr *MockICoreMockRecorder) GetUserId(ctx, sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserId", reflect.TypeOf((*MockICore)(nil).GetUserId), ctx, sid)
}

// Trends mocks base method.
func (m *MockICore) Trends() ([]models.FilmItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trends")
	ret0, _ := ret[0].([]models.FilmItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Trends indicates an expected call of Trends.
func (mr *MockICoreMockRecorder) Trends() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trends", reflect.TypeOf((*MockICore)(nil).Trends))
}

// UsersStatistics mocks base method.
func (m *MockICore) UsersStatistics(idUser uint64) ([]requests.UsersStatisticsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersStatistics", idUser)
	ret0, _ := ret[0].([]requests.UsersStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersStatistics indicates an expected call of UsersStatistics.
func (mr *MockICoreMockRecorder) UsersStatistics(idUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersStatistics", reflect.TypeOf((*MockICore)(nil).UsersStatistics), idUser)
}
