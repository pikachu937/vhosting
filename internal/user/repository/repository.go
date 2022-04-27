package repository

import (
	"database/sql"
	"fmt"
	"reflect"

	dbc "github.com/mikerumy/vhosting2/internal/constants/db"
	"github.com/mikerumy/vhosting2/internal/models"
	sq "github.com/mikerumy/vhosting2/pkg/constants/sql"
	"github.com/mikerumy/vhosting2/pkg/database"
)

type UserRepository struct {
	cfg models.Config
}

func NewUserRepository(cfg models.Config) *UserRepository {
	return &UserRepository{cfg: cfg}
}

func (r *UserRepository) CreateUser(usr models.User) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s, %s)", dbc.UsersTable, dbc.Username,
		dbc.PassHash, dbc.IsActive, dbc.IsSuperUser, dbc.IsStaff, dbc.FirstName,
		dbc.LastName, dbc.DateJoined, dbc.LastLogin)
	val := "($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, usr.Username, usr.PasswordHash, usr.IsActive, usr.IsSuperUser, usr.IsStaff,
		usr.FirstName, usr.LastName, usr.DateJoined, usr.LastLogin)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUser(id int) (*models.User, error) {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s", dbc.Id, dbc.Username,
		dbc.PassHash, dbc.IsActive, dbc.IsSuperUser, dbc.IsStaff, dbc.FirstName,
		dbc.LastName, dbc.DateJoined, dbc.LastLogin)
	tbl := dbc.UsersTable
	cnd := fmt.Sprintf("%s=$1", dbc.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var usr models.User
	err := db.Get(&usr, query, id)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *UserRepository) GetAllUsers() (map[int]*models.User, error) {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.SELECT_COL_FROM_TBL
	col := "*"
	tbl := dbc.UsersTable
	query := fmt.Sprintf(template, col, tbl)

	var rows *sql.Rows
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = map[int]*models.User{}
	var usr models.User
	for rows.Next() {
		err = rows.Scan(&usr.Id, &usr.Username, &usr.PasswordHash, &usr.IsActive, &usr.IsSuperUser,
			&usr.IsStaff, &usr.FirstName, &usr.LastName, &usr.DateJoined, &usr.LastLogin)
		if err != nil {
			return nil, err
		}
		users[usr.Id] = &models.User{Id: usr.Id, Username: usr.Username, PasswordHash: usr.PasswordHash,
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

func (r *UserRepository) PartiallyUpdateUser(id int, usr models.User) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := dbc.UsersTable
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END, ", dbc.Username, dbc.Username) +
		fmt.Sprintf("%s=CASE WHEN $2 <> '' THEN $2 ELSE %s END, ", dbc.PassHash, dbc.PassHash) +
		fmt.Sprintf("%s=$3, ", dbc.IsActive) +
		fmt.Sprintf("%s=$4, ", dbc.IsSuperUser) +
		fmt.Sprintf("%s=$5, ", dbc.IsStaff) +
		fmt.Sprintf("%s=CASE WHEN $6 <> '' THEN $6 ELSE %s END, ", dbc.FirstName, dbc.FirstName) +
		fmt.Sprintf("%s=CASE WHEN $7 <> '' THEN $7 ELSE %s END", dbc.LastName, dbc.LastName)
	cnd := fmt.Sprintf("%s=$8", dbc.Id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, usr.Username, usr.PasswordHash, usr.IsActive, usr.IsSuperUser,
		usr.IsStaff, usr.FirstName, usr.LastName, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.DELETE_FROM_TBL_WHERE_CND
	tbl := dbc.UsersTable
	cnd := fmt.Sprintf("%s=$1", dbc.Id)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *UserRepository) IsUserExists(idOrUsername interface{}) (bool, error) {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	var template, col, tbl, cnd, query string
	var rows *sql.Rows
	var err error

	if reflect.TypeOf(idOrUsername) == reflect.TypeOf(0) {
		template = sq.SELECT_COL_FROM_TBL_WHERE_CND
		col = dbc.Id
		tbl = dbc.UsersTable
		cnd = fmt.Sprintf("%s=$1", dbc.Id)
		query = fmt.Sprintf(template, col, tbl, cnd)
		rows, err = db.Query(query, idOrUsername.(int))
	} else {
		template = sq.SELECT_COL_FROM_TBL_WHERE_CND
		col = dbc.Username
		tbl = dbc.UsersTable
		cnd = fmt.Sprintf("%s=$1", dbc.Username)
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
