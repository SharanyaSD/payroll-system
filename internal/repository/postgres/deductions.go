package repository

import (
	"fmt"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	"github.com/jmoiron/sqlx"
)

type DeductionsStore struct {
	Db *sqlx.DB
}

func NewDeductionsRepo(db *sqlx.DB) repository.DeductionsStorer {
	return &DeductionsStore{
		Db: db,
	}
}

func (ds *DeductionsStore) GetDeductionsByEmpoyeeID(ID string) (repository.Deductions, error) {
	var deductions repository.Deductions
	query := "SELECT * from deductions where id=$1"
	fmt.Println("SQL Query:", query)
	row := ds.Db.QueryRow(query, ID)
	err := row.Scan(
		&deductions.ID, &deductions.TDS, &deductions.PF, &deductions.Medical, &deductions.GrossDeduction,
	)
	if err != nil {
		return deductions, err
	}
	return deductions, nil
}

func (ds *DeductionsStore) InsertDeductions(deductions repository.Deductions) (repository.Deductions, error) {
	_, err := ds.Db.Exec("INSERT INTO deductions (id, tds, pf, medical, gross_deduction) VALUES ($1, $2, $3, $4, $5)",
		deductions.ID, deductions.TDS, deductions.PF, deductions.Medical, deductions.GrossDeduction)
	fmt.Println("Deductions added to ID ", deductions.ID, " - ", deductions.GrossDeduction)
	if err != nil {
		return repository.Deductions{}, err
	}
	return deductions, nil
}
