package repository

import "database/sql"

type Users struct {
	ID    sql.NullInt64
	Name  sql.NullString
	Email sql.NullString
}
