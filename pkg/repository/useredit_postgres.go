package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	vhs "github.com/mikerumy/vhservice"
)

type UserEditPostgres struct {
	db *sqlx.DB
}

func NewUserEditPostgres(db *sqlx.DB) *UserEditPostgres {
	return &UserEditPostgres{db: db}
}

func (r *UserEditPostgres) CreateUser(user vhs.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, email) VALUES ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Email)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserEditPostgres) GetUser(id int) (*vhs.User, error) {
	var user vhs.User

	query := fmt.Sprintf("SELECT id, name, email FROM %s WHERE id=$1", usersTable)

	err := r.db.Get(&user, query, id)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *UserEditPostgres) UpdateUser(id int, user vhs.User) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET name=$1, email=$2 WHERE id=$3", usersTable)

	_, err := r.db.Query(query, user.Name, user.Email, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserEditPostgres) PartiallyUpdateUser(id int, user vhs.User) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET name=CASE WHEN $1 <> '' THEN $1 ELSE name END, email=CASE WHEN $2 <> '' THEN $2 ELSE email END WHERE id=$3", usersTable)

	_, err := r.db.Query(query, user.Name, user.Email, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserEditPostgres) DeleteUser(id int) (int, error) {
	var name string

	query := fmt.Sprintf("SELECT name FROM %s WHERE id=$1", usersTable)

	row, err := r.db.Query(query, id)
	if err != nil {
		return -1, err
	}

	row.Next()
	if row.Scan(&name); name == "" {
		return -1, errors.New("user not found")
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	row, err = r.db.Query(query, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
