package main

import "github.com/mikerumy/vhosting/pkg/config"

type NonCatVideo struct {
	Id             int
	CodeMP         string
	StartDatetime  string
	DurationRecord int
}

type Repo struct {
	cfg *config.Config
}
