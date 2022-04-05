package repository

import (
	"fmt"

	vhs "github.com/mikerumy/vhservice"
)

type AuthorizationRepo struct {
	cfg vhs.DBConfig
}

func NewAuthorizationRepo(cfg vhs.DBConfig) *AuthorizationRepo {
	return &AuthorizationRepo{cfg: cfg}
}

func (r *AuthorizationRepo) POSTUser(user vhs.User) (int, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var id int

	query := fmt.Sprintf("INSERT INTO users (username, password_hash) values ($1, $2) RETURNING id")

	row := db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorizationRepo) GETUser(username, password string) (vhs.User, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var user vhs.User

	query := fmt.Sprintf("SELECT id FROM users WHERE username=$1 AND password_hash=$2")

	err := db.Get(&user, query, username, password)

	return user, err
}
