package dao

import (
	"cudo/test/config"
	"cudo/test/repository"
)

func GetFreqCheck() (results []repository.TransactionsModel, err error) {
	db := config.DB
	query := "select " +
		"     tr1.user_id," +
		"     tr1.transaction_date," +
		"     count(tr2.id) as count_in_hour" +
		" from transactions as tr1" +
		" join transactions as tr2 on tr1.user_id = tr2.user_id and" +
		"     tr2.transaction_date between tr2.transaction_date - INTERVAL '1 hour' and tr1.transaction_date" +
		" group by tr1.user_id, tr1.transaction_date" +
		" having count(tr2.id) > 5 "

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var temp repository.TransactionsModel
		errs := rows.Scan(&temp.UserID, &temp.TransactionDate, &temp.CountInHour)
		if errs != nil {
			return nil, errs
		}

		results = append(results, temp)
	}
	return
}
