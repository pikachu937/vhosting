package storage

import (
	"database/sql"
	"errors"

	vhs "github.com/mikerumy/vhservice"
	auth "github.com/mikerumy/vhservice/pkg/service/methods"
)

type UserInterfaceStorage struct {
	cfg vhs.DBConfig
}

func NewUserInterfaceStorage(cfg vhs.DBConfig) *UserInterfaceStorage {
	return &UserInterfaceStorage{cfg: cfg}
}

func (r *UserInterfaceStorage) POSTUser(user vhs.User) (int, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var id int
	var row *sql.Row

	query := "INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id"

	row = db.QueryRow(query, user.Username, auth.GeneratePasswordHash(user.PasswordHash))

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfaceStorage) GETUser(id int) (*vhs.User, error) {
	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var user vhs.User

	query := "SELECT id, username, password_hash FROM users WHERE id=$1"

	if err := db.Get(&user, query, id); err != nil {
		return nil, errors.New("user not found in database")
	}

	return &user, nil
}

func (r *UserInterfaceStorage) GETAllUsers() (map[int]*vhs.User, error) {
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
		rows.Scan(&user.Id, &user.Username, &user.PasswordHash)
		users[user.Id] = &vhs.User{Id: user.Id, Username: user.Username, PasswordHash: user.PasswordHash}
	}

	if len(users) == 0 {
		return nil, errors.New("users not found in database")
	}

	return users, nil
}

func (r *UserInterfaceStorage) PUTUser(id int, user vhs.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var rows *sql.Rows
	var err error

	query := "UPDATE users SET username=$1, password_hash=$2 WHERE id=$3"

	if rows, err = db.Query(query, user.Username, user.PasswordHash, id); err != nil {
		return -1, err
	}
	defer rows.Close()

	return id, nil
}

func (r *UserInterfaceStorage) PATCHUser(id int, user vhs.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vhs.NewDBConnection(r.cfg)
	defer vhs.CloseDBConnection(db)

	var rows *sql.Rows
	var err error

	query := "UPDATE users SET username=CASE WHEN $1 <> '' THEN $1 ELSE username END, password_hash=CASE WHEN $2 <> '' THEN $2 ELSE password END WHERE id=$3"

	if rows, err = db.Query(query, user.Username, user.PasswordHash, id); err != nil {
		return -1, err
	}
	defer rows.Close()

	return id, nil
}

func (r *UserInterfaceStorage) DELETEUser(id int) (int, error) {
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

func (r *UserInterfaceStorage) checkUserInDB(id int) error {
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
		return errors.New("user not found in database")
	}

	return nil
}
