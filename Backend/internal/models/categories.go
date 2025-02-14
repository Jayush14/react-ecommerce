package models

import "database/sql"
type Category struct {
	Category sql.NullString `json:"value"`
}