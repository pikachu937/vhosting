package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	vh "github.com/mikerumy/vhosting"
)

type UserInterfaceStorage struct {
	cfg vh.DBConfig
}

func NewUserInterfaceStorage(cfg vh.DBConfig) *UserInterfaceStorage {
	return &UserInterfaceStorage{cfg: cfg}
}

func (r *UserInterfaceStorage) POSTUser(user vh.User) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", vh.UsersTable, vh.Username, vh.PassHash,
		vh.IsActive, vh.IsSuperUser, vh.IsStaff, vh.FirstName, vh.LastName, vh.Email, vh.DateJoined, vh.LastLogin)
	val := "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	query := fmt.Sprintf(template, tbl, val)

	var userDateJoined time.Time = time.Now()
	var userLastLogin time.Time = userDateJoined
	var userPassHash string = vh.GeneratePasswordHash(user.PasswordHash)
	var userIsActive bool = false
	if user.IsActive == nil {
		userIsActive = true
	}
	_, err := db.Query(query, user.Username, userPassHash, userIsActive, user.IsSuperUser, user.IsStaff,
		user.FirstName, user.LastName, user.Email, userDateJoined, userLastLogin)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserInterfaceStorage) GETUser(id int) (*vh.User, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s", vh.Id, vh.Username, vh.PassHash,
		vh.IsActive, vh.IsSuperUser, vh.IsStaff, vh.FirstName, vh.LastName, vh.Email, vh.DateJoined, vh.LastLogin)
	tbl := vh.UsersTable
	cnd := fmt.Sprintf("%s=$1", vh.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var user vh.User
	err := db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserInterfaceStorage) GETAllUsers() (map[int]*vh.User, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.SELECT_COL_FROM_TBL
	col := "*"
	tbl := vh.UsersTable
	query := fmt.Sprintf(template, col, tbl)

	var rows *sql.Rows
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = map[int]*vh.User{}
	var user vh.User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.PasswordHash, &user.IsActive, &user.IsSuperUser,
			&user.IsStaff, &user.FirstName, &user.LastName, &user.Email, &user.DateJoined, &user.LastLogin)
		if err != nil {
			return nil, err
		}
		users[user.Id] = &vh.User{Id: user.Id, Username: user.Username, PasswordHash: user.PasswordHash,
			IsActive: user.IsActive, IsSuperUser: user.IsSuperUser, IsStaff: user.IsStaff, FirstName: user.FirstName,
			LastName: user.LastName, Email: user.Email, DateJoined: user.DateJoined, LastLogin: user.LastLogin}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users to get")
	}

	return users, nil
}

func (r *UserInterfaceStorage) PATCHUser(id int, user vh.User) error {
	err := r.checkUserExistence(id)
	if err != nil {
		return err
	}

	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := vh.UsersTable
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END, ", vh.Username, vh.Username) +
		fmt.Sprintf("%s=CASE WHEN $2 <> '' THEN $2 ELSE %s END, ", vh.PassHash, vh.PassHash) +
		fmt.Sprintf("%s=$3, ", vh.IsActive) +
		fmt.Sprintf("%s=$4, ", vh.IsSuperUser) +
		fmt.Sprintf("%s=$5, ", vh.IsStaff) +
		fmt.Sprintf("%s=CASE WHEN $6 <> '' THEN $6 ELSE %s END, ", vh.FirstName, vh.FirstName) +
		fmt.Sprintf("%s=CASE WHEN $7 <> '' THEN $7 ELSE %s END, ", vh.LastName, vh.LastName) +
		fmt.Sprintf("%s=CASE WHEN $8 <> '' THEN $8 ELSE %s END", vh.Email, vh.Email)
	cnd := fmt.Sprintf("%s=$9", vh.Id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	var userPassHash string = vh.GeneratePasswordHash(user.PasswordHash)
	var userIsActive bool = false
	if user.IsActive == nil {
		userIsActive = true
	}
	rows, err = db.Query(query, user.Username, userPassHash, userIsActive, user.IsSuperUser, user.IsStaff,
		user.FirstName, user.LastName, user.Email, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *UserInterfaceStorage) DELETEUser(id int) error {
	err := r.checkUserExistence(id)
	if err != nil {
		return err
	}

	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.DELETE_FROM_TBL_WHERE_CND
	tbl := vh.UsersTable
	cnd := fmt.Sprintf("%s=$1", vh.Id)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err = db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *UserInterfaceStorage) checkUserExistence(id int) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.SELECT_COL_FROM_TBL_WHERE_CND
	col := vh.Id
	tbl := vh.UsersTable
	cnd := fmt.Sprintf("%s=$1", vh.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var idInColumn string
	rows.Next()
	rows.Scan(&idInColumn)
	if err != nil {
		return err
	}

	if idInColumn == "" {
		return errors.New("user not found")
	}

	return nil
}
