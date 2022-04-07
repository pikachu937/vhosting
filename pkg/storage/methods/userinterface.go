package storage

import (
	"database/sql"
	"errors"
	"fmt"

	vh "github.com/mikerumy/vhosting"
)

type UserInterfaceStorage struct {
	cfg vh.DBConfig
}

func NewUserInterfaceStorage(cfg vh.DBConfig) *UserInterfaceStorage {
	return &UserInterfaceStorage{cfg: cfg}
}

func (r *UserInterfaceStorage) POSTUser(user vh.User) (int, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	sql := vh.INSERT_TBL_COL_VAL_RET
	tbl := vh.UsersTable
	col := fmt.Sprintf("(%s, %s)", vh.Username, vh.PassHash)
	val := "($1, $2)"
	ret := vh.Id
	query := fmt.Sprintf(sql, tbl, col, val, ret)

	userPassHash := vh.GeneratePasswordHash(user.PasswordHash)
	row := db.QueryRow(query, user.Username, userPassHash)

	var id int

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UserInterfaceStorage) GETUser(id int) (*vh.User, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	sql := vh.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s", vh.Id, vh.Username, vh.PassHash)
	tbl := vh.UsersTable
	cnd := fmt.Sprintf("%s=$1", vh.Id)
	query := fmt.Sprintf(sql, col, tbl, cnd)

	var user vh.User

	if err := db.Get(&user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserInterfaceStorage) GETAllUsers() (map[int]*vh.User, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	sql := vh.SELECT_COL_FROM_TBL
	col := "*"
	tbl := vh.UsersTable
	query := fmt.Sprintf(sql, col, tbl)

	users := make(map[int]*vh.User)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user vh.User
		if err := rows.Scan(&user.Id, &user.Username, &user.PasswordHash); err != nil {
			return nil, err
		}
		users[user.Id] = &vh.User{Id: user.Id, Username: user.Username, PasswordHash: user.PasswordHash}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserInterfaceStorage) PUTUser(id int, user vh.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	var rows *sql.Rows
	var err error

	sql := vh.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := vh.UsersTable
	val := fmt.Sprintf("%s=$1, %s=$2", vh.Username, vh.PassHash)
	cnd := fmt.Sprintf("%s=$3", vh.Id)
	query := fmt.Sprintf(sql, tbl, val, cnd)

	if rows, err = db.Query(query, user.Username, user.PasswordHash, id); err != nil {
		return -1, err
	}
	defer rows.Close()

	return id, nil
}

func (r *UserInterfaceStorage) PATCHUser(id int, user vh.User) (int, error) {
	if err := r.checkUserInDB(id); err != nil {
		return -1, err
	}

	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	var rows *sql.Rows
	var err error

	// query := "UPDATE users SET username=CASE WHEN $1 <> '' THEN $1 ELSE username END, password_hash=CASE WHEN $2 <> '' THEN $2 ELSE password END WHERE id=$3"
	query := "UPDATE users SET username=IF $1 <> '' THEN $1 ELSE username END IF, password_hash=IF $2 <> '' THEN $2 ELSE password END IF WHERE id=$3"

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

	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

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
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

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
