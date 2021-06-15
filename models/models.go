package models

import "time"

type Product struct {
	ID int 					`json:"id"`
	Name string				`json:"name"`
	Description string		`json:"description"`
	CreatedOn time.Time 	`json:"created_on"`
	ChangedOn time.Time 	`json:"changed_on"`
}