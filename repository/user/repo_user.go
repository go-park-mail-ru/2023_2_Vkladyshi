package profile

import "database/sql"

type IUserRepo interface {
	GetUser(login string, password string) (*UserItem, bool, error)
	FindUser(login string) (bool, error)
	CreateUser(login string, password string, name string, birthDate string, email string) error
}

type RepoPostgre struct {
	DB *sql.DB
}

func (repo *RepoPostgre) GetUser(login string, password string) (*UserItem, bool, error) {
	post := &UserItem{}

	err := repo.DB.QueryRow(
		"SELECT login, photo FROM profiles"+
			"WHERE login = $1 AND password = $2", login, password).Scan(&post.Login, &post.Photo)
	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	return post, true, nil
}

func (repo *RepoPostgre) FindUser(login string) (bool, error) {
	post := &UserItem{}

	err := repo.DB.QueryRow(
		"SELECT login FROM profiles"+
			"WHERE login = $1", login).Scan(&post.Login)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *RepoPostgre) CreateUser(login string, password string, name string, birthDate string, email string) error {
	_, err := repo.DB.Exec(
		"INSERT INTO profiles(name, birth_date, photo, login, password, email, registration_date)"+
			"VALUES($1, $2, '../../user_avatars/default.jpg', $3, $4, $5, CURRENT_TIMESTAMP)",
		name, birthDate, login, password, email)
	if err != nil {
		return err
	}

	return nil
}
