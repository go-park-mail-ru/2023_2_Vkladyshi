package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/comment"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/profile"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/repository/session"
)

type ICore interface {
	CreateSession(ctx context.Context, login string) (string, session.Session, error)
	KillSession(ctx context.Context, sid string) error
	FindActiveSession(ctx context.Context, sid string) (bool, error)
	CreateUserAccount(login string, password string, name string, birthDate string, email string) error
	FindUserAccount(login string, password string) (*profile.UserItem, bool, error)
	FindUserByLogin(login string) (bool, error)
	GetFilmComments(filmId uint64, first uint64, limit uint64) ([]comment.CommentItem, error)
	AddComment(filmId uint64, userLogin string, rating uint16, text string) error
	GetUserName(ctx context.Context, sid string) (string, error)
	GetUserProfile(login string) (*profile.UserItem, error)
	CheckCsrfToken(ctx context.Context, token string) (bool, error)
	CreateCsrfToken(ctx context.Context) (string, error)
	EditProfile(prevLogin string, login string, password string, email string, birthDate string, photo string) error
	FindUsersComment(login string, filmId uint64) (bool, error)
}
