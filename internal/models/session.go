package models

type Session struct {
	Id           int    `db:"id"`
	Content      string `db:"content"`
	CreationDate string `db:"creation_date"`
}