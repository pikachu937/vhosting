package storage

import (
	"fmt"

	vhs "github.com/mikerumy/vhservice"
)

type AuthorizationStorage struct {
	cfg vhs.DBConfig
}

func NewAuthorizationStorage(cfg vhs.DBConfig) *AuthorizationStorage {
	return &AuthorizationStorage{cfg: cfg}
}

func (r *AuthorizationStorage) POSTUser(user vhs.User) (int, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var id int

	query := fmt.Sprintf("INSERT INTO users (username, password_hash) values ($1, $2) RETURNING id")

	row := db.QueryRow(query, user.Username, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorizationStorage) GETUser(username, password string) (vhs.User, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var user vhs.User

	query := fmt.Sprintf("SELECT id FROM users WHERE username=$1 AND password_hash=$2")

	err := db.Get(&user, query, username, password)

	return user, err
}
