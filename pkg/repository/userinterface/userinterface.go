package repository

import (
	"database/sql"
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
	var row *sql.Row

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"

	row = db.QueryRow(query, user.Username, user.Password)

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

func (r *UserInterfaceRepo) GETAllUsers() (map[int]*vhs.User, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	users := make(map[int]*vhs.User)
	var rows *sql.Rows
	var err error

	query := "SELECT * FROM users"

	if rows, err = db.Query(query); err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user vhs.User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users[user.Id] = &vhs.User{Id: user.Id, Username: user.Username, Password: user.Password}
	}

	return users, nil
}

func (r *UserInterfaceRepo) PUTUser(id int, user vhs.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var rows *sql.Rows
	var err error

	query := "UPDATE users SET username=$1, password=$2 WHERE id=$3"

	if rows, err = db.Query(query, user.Username, user.Password, id); err != nil {
		return -1, err
	}
	defer rows.Close()

	return id, nil
}

func (r *UserInterfaceRepo) PATCHUser(id int, user vhs.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var rows *sql.Rows
	var err error

	query := "UPDATE users SET username=CASE WHEN $1 <> '' THEN $1 ELSE username END, password=CASE WHEN $2 <> '' THEN $2 ELSE password END WHERE id=$3"

	if rows, err = db.Query(query, user.Username, user.Password, id); err != nil {
		return -1, err
	}
	defer rows.Close()

	return id, nil
}

func (r *UserInterfaceRepo) DELETEUser(id int) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var rows *sql.Rows
	var err error

	query := "DELETE FROM users WHERE id=$1"

	if rows, err = db.Query(query, id); err != nil {
		return -1, err
	}
	defer rows.Close()

	return id, nil
}

func (r *UserInterfaceRepo) checkUserInDB(id int) error {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var idStr string
	var rows *sql.Rows
	var err error

	query := "SELECT id FROM users WHERE id=$1"

	if rows, err = db.Query(query, id); err != nil {
		return err
	}
	defer rows.Close()

	rows.Next()
	if rows.Scan(&idStr); idStr == "" {
		return errors.New("user not present in database")
	}

	return nil
}
