package repository

import "database/sql"

type TransactionsModel struct {
	ID              sql.NullInt64
	UserID          sql.NullInt64
	OrderID         sql.NullString
	TransactionDate sql.NullString
	Amount          sql.NullFloat64
	PaymentMethod   sql.NullString
	Status          sql.NullString
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	CountInHour     sql.NullInt64
}
