package storage

import (
	"fmt"

	vh "github.com/mikerumy/vhosting"
)

type AuthorizationStorage struct {
	cfg vh.DBConfig
}

func NewAuthorizationStorage(cfg vh.DBConfig) *AuthorizationStorage {
	return &AuthorizationStorage{cfg: cfg}
}

func (r *AuthorizationStorage) POSTUser(user vh.User) (int, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	var id int

	query := fmt.Sprintf("INSERT INTO users (username, password_hash) values ($1, $2) RETURNING id")

	row := db.QueryRow(query, user.Username, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorizationStorage) GETUser(username, password string) (vh.User, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	var user vh.User

	query := fmt.Sprintf("SELECT id FROM users WHERE username=$1 AND password_hash=$2")

	err := db.Get(&user, query, username, password)

	return user, err
}
