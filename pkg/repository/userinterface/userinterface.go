package repository

import (
	"errors"

	vhs "github.com/mikerumy/vhservice"
)

type UserInterfaceRepo struct {
	cfg vhs.DBConfig
}

func NewUserInterfaceRepo(cfg vhs.DBConfig) *UserInterfaceRepo {
	return &UserInterfaceRepo{cfg: cfg}
}

func (r *UserInterfaceRepo) POSTUser(user vhs.User) (int, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var id int

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"

	row := db.QueryRow(query, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfaceRepo) GETUser(id int) (*vhs.User, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var user vhs.User

	query := "SELECT id, username, password FROM users WHERE id=$1"

	if err := db.Get(&user, query, id); err != nil {
		return nil, errors.New("user not present in database")
	}

	return &user, nil
}

func (r *UserInterfaceRepo) PUTUser(id int, user vhs.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	query := "UPDATE users SET username=$1, password=$2 WHERE id=$3"

	_, err := db.Query(query, user.Username, user.Password, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfaceRepo) PATCHUser(id int, user vhs.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	query := "UPDATE users SET username=CASE WHEN $1 <> '' THEN $1 ELSE username END, password=CASE WHEN $2 <> '' THEN $2 ELSE password END WHERE id=$3"

	_, err := db.Query(query, user.Username, user.Password, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfaceRepo) DELETEUser(id int) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	query := "DELETE FROM users WHERE id=$1"
	_, err := db.Query(query, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfaceRepo) checkUserInDB(id int) error {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var idStr string

	query := "SELECT id FROM users WHERE id=$1"

	row, err := db.Query(query, id)
	if err != nil {
		return err
	}

	row.Next()
	if row.Scan(&idStr); idStr == "" {
		return errors.New("user not present in database")
	}

	return nil
}
