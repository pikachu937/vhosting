package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	vhs "github.com/mikerumy/vhservice"
)

type UserInterfacePostgres struct {
	db *sqlx.DB
}

func NewUserInterfacePostgres(db *sqlx.DB) *UserInterfacePostgres {
	return &UserInterfacePostgres{db: db}
}

func (r *UserInterfacePostgres) POSTUser(user vhs.User) (int, error) {

	var id int

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"

	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfacePostgres) GETUser(id int) (*vhs.User, error) {
	var user vhs.User

	query := "SELECT id, username, password FROM users WHERE id=$1"

	err := r.db.Get(&user, query, id)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *UserInterfacePostgres) PUTUser(id int, user vhs.User) (int, error) {
	query := "UPDATE users SET username=$1, password=$2 WHERE id=$3"

	_, err := r.db.Query(query, user.Username, user.Password, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfacePostgres) PATCHUser(id int, user vhs.User) (int, error) {
	query := "UPDATE users SET username=CASE WHEN $1 <> '' THEN $1 ELSE username END, password=CASE WHEN $2 <> '' THEN $2 ELSE password END WHERE id=$3"

	_, err := r.db.Query(query, user.Username, user.Password, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfacePostgres) DELETEUser(id int) (int, error) {
	var name string

	query := "SELECT username FROM users WHERE id=$1"

	row, err := r.db.Query(query, id)
	if err != nil {
		return -1, err
	}

	row.Next()
	if row.Scan(&name); name == "" {
		return -1, errors.New("user not found")
	}

	query = "DELETE FROM users WHERE id=$1"
	row, err = r.db.Query(query, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
