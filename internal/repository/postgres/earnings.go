package repository

import (
	"fmt"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	"github.com/jmoiron/sqlx"
)

type EarningsStore struct {
	Db *sqlx.DB
}

func NewEarningsRepo(db *sqlx.DB) repository.EarningsStorer {
	return &EarningsStore{
		Db: db,
	}
}

func (es *EarningsStore) GetEarningsByEmpoyeeID(ID string) (repository.Earnings, error) {
	var earning repository.Earnings
	query := "SELECT * from earnings where id=$1"
	fmt.Println("SQL Query :", ID, query)
	row := es.Db.QueryRow(query, ID)
	err := row.Scan(
		&earning.Basic, &earning.HRA, &earning.DA, &earning.SA, &earning.CA,
		&earning.Bonus, &earning.GrossPay, &earning.ID,
	)
	if err != nil {
		return earning, err
	}
	return earning, nil
}

func (es *EarningsStore) InsertEarnings(earnings repository.Earnings) (repository.Earnings, error) {
	_, err := es.Db.Exec("INSERT INTO earnings (id, basic, hra, da, sa, ca, bonus, gross_pay) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		earnings.ID, earnings.Basic, earnings.HRA, earnings.DA, earnings.SA, earnings.CA, earnings.Bonus, earnings.GrossPay)
	fmt.Println("Earnings added to ID ", earnings.ID, " - ", earnings.GrossPay)

	if err != nil {
		return repository.Earnings{}, err
	}
	return earnings, nil
}
