package models

import "time"

type CreateTransaction struct {
	Category string
	UserID   string
	Name     string
	Cost     float32
}

type Transaction struct {
	ID       string
	UserID   string
	Category string
	Name     string
	Cost     float32
	Date     string
}

type TimeFrame struct {
	StartDate time.Time
	EndDate   time.Time
}

type CreateTimeFrame struct {
	StartDate string
	EndDate   string
}
