package models

import(
 "database/sql"
)
type Brand struct {
	Brand   sql.NullString `json:"value"`
}