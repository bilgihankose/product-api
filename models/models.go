package models

import "time"

type Product struct {
	ID int
	Name string
	Description string
	CreatedOn time.Time
	ChangedOn time.Time
}