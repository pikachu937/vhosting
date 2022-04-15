package storage

import (
	"database/sql"
	"fmt"
	"reflect"

	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/internal/config"
	user "github.com/mikerumy/vhosting/internal/user"
)

type UserInterfaceStorage struct {
	cfg config.DBConfig
}

func NewUserInterfaceStorage(cfg config.DBConfig) *UserInterfaceStorage {
	return &UserInterfaceStorage{cfg: cfg}
}

func (r *UserInterfaceStorage) CheckUserExistence(idOrUsername interface{}) (bool, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	var template, col, tbl, cnd, query string
	var rows *sql.Rows
	var err error

	if reflect.TypeOf(idOrUsername) == reflect.TypeOf(0) {
		template = vh.SELECT_COL_FROM_TBL_WHERE_CND
		col = vh.Id
		tbl = vh.UsersTable
		cnd = fmt.Sprintf("%s=$1", vh.Id)
		query = fmt.Sprintf(template, col, tbl, cnd)
		rows, err = db.Query(query, idOrUsername.(int))
	} else {
		template = vh.SELECT_COL_FROM_TBL_WHERE_CND
		col = vh.Username
		tbl = vh.UsersTable
		cnd = fmt.Sprintf("%s=$1", vh.Username)
		query = fmt.Sprintf(template, col, tbl, cnd)
		rows, err = db.Query(query, idOrUsername.(string))
	}
	if err != nil {
		return false, err
	}
	defer rows.Close()

	rowIsPresent := rows.Next()
	if !rowIsPresent {
		return false, nil
	}

	return true, nil
}

func (r *UserInterfaceStorage) POSTUser(usr user.User) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s, %s)", vh.UsersTable, vh.Username, vh.PassHash,
		vh.IsActive, vh.IsSuperUser, vh.IsStaff, vh.FirstName, vh.LastName, vh.DateJoined, vh.LastLogin)
	val := "($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, usr.Username, usr.PasswordHash, usr.IsActive, usr.IsSuperUser, usr.IsStaff,
		usr.FirstName, usr.LastName, usr.DateJoined, usr.LastLogin)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserInterfaceStorage) GETUser(id int) (*user.User, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s", vh.Id, vh.Username, vh.PassHash,
		vh.IsActive, vh.IsSuperUser, vh.IsStaff, vh.FirstName, vh.LastName, vh.DateJoined, vh.LastLogin)
	tbl := vh.UsersTable
	cnd := fmt.Sprintf("%s=$1", vh.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var usr user.User
	err := db.Get(&usr, query, id)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *UserInterfaceStorage) GETAllUsers() (map[int]*user.User, error) {
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

	var users = map[int]*user.User{}
	var usr user.User
	for rows.Next() {
		err = rows.Scan(&usr.Id, &usr.Username, &usr.PasswordHash, &usr.IsActive, &usr.IsSuperUser,
			&usr.IsStaff, &usr.FirstName, &usr.LastName, &usr.DateJoined, &usr.LastLogin)
		if err != nil {
			return nil, err
		}
		users[usr.Id] = &user.User{Id: usr.Id, Username: usr.Username, PasswordHash: usr.PasswordHash,
			IsActive: usr.IsActive, IsSuperUser: usr.IsSuperUser, IsStaff: usr.IsStaff,
			FirstName: usr.FirstName, LastName: usr.LastName, DateJoined: usr.DateJoined,
			LastLogin: usr.LastLogin}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func (r *UserInterfaceStorage) PATCHUser(id int, usr user.User) error {
	exist, err := r.CheckUserExistence(id)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}

	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := vh.UsersTable
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END, ", vh.Username, vh.Username) +
		fmt.Sprintf("%s=CASE WHEN $2 <> '' THEN $2 ELSE %s END, ", vh.PassHash, vh.PassHash) +
		fmt.Sprintf("%s=CASE WHEN $3 <> false THEN $3 ELSE %s END, ", vh.IsActive, vh.IsActive) +
		fmt.Sprintf("%s=CASE WHEN $4 <> true THEN $4 ELSE %s END, ", vh.IsSuperUser, vh.IsSuperUser) +
		fmt.Sprintf("%s=CASE WHEN $5 <> true THEN $5 ELSE %s END, ", vh.IsStaff, vh.IsStaff) +
		fmt.Sprintf("%s=CASE WHEN $6 <> '' THEN $6 ELSE %s END, ", vh.FirstName, vh.FirstName) +
		fmt.Sprintf("%s=CASE WHEN $7 <> '' THEN $7 ELSE %s END, ", vh.LastName, vh.LastName)
	cnd := fmt.Sprintf("%s=$8", vh.Id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err = db.Query(query, usr.Username, usr.PasswordHash, usr.IsActive, usr.IsSuperUser,
		usr.IsStaff, usr.FirstName, usr.LastName, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *UserInterfaceStorage) DELETEUser(id int) error {
	exist, err := r.CheckUserExistence(id)
	if err != nil {
		return err
	}
	if !exist {
		return nil
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
