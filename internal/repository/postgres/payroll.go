package repository

import (
	"fmt"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	"github.com/jmoiron/sqlx"
)

type PayrollStore struct {
	Db *sqlx.DB
}

func NewPayrollRepo(db *sqlx.DB) repository.PayrollStorer {
	return &PayrollStore{
		Db: db,
	}
}

func (pr *PayrollStore) CreatePayroll(payroll repository.Payroll) (repository.Payroll, error) {

	_, err := pr.Db.Exec("INSERT INTO payroll (ID, salary, net_pay_salary, pay_date) VALUES ($1, $2, $3, $4)",
		payroll.ID, payroll.Salary, payroll.NetPaySalary, payroll.PayDate)
	fmt.Println("ID passed ", payroll.ID)
	if err != nil {
		return repository.Payroll{}, err
	}
	return payroll, nil
}

func (pr *PayrollStore) GetPayroll() ([]repository.Payroll, error) {

	var payrolls []repository.Payroll
	err := pr.Db.Select(&payrolls, "SELECT * FROM payroll")
	if err != nil {
		return nil, err
	}
	return payrolls, nil
}

// func (pr *PayrollStore) GetPayrollByID(ID string) (repository.Payroll, error) {
// 	var payroll repository.Payroll
// 	query := "SELECT * FROM Payroll WHERE id=$1"
// 	fmt.Println("SQL Query:", query)
// 	row := pr.Db.QueryRow(query, ID)
// 	err := row.Scan(
// 		&payroll.ID, &payroll.Salary, &payroll.NetPaySalary, &payroll.PayDate,
// 	)
// 	if err != nil {
// 		return payroll, err
// 	}
// 	return payroll, nil

// }

func (pr *PayrollStore) InsertEarnings(earnings repository.Earnings) (repository.Earnings, error) {
	_, err := pr.Db.Exec("INSERT INTO earnings (id, basic, hra, da, sa, ca, bonus, gross_pay) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		earnings.ID, earnings.Basic, earnings.HRA, earnings.DA, earnings.SA, earnings.CA, earnings.Bonus, earnings.GrossPay)
	fmt.Println("Earnings added ", earnings.ID)
	if err != nil {
		return repository.Earnings{}, err
	}
	return earnings, nil
}

func (pr *PayrollStore) InsertDeductions(deductions repository.Deductions) (repository.Deductions, error) {
	_, err := pr.Db.Exec("INSERT INTO deductions (id, tds, pf, medical, gross_deduction) VALUES ($1, $2, $3, $4, $5)",
		deductions.ID, deductions.TDS, deductions.PF, deductions.Medical, deductions.GrossDeduction)
	fmt.Println("Earnings added ", deductions.ID)
	if err != nil {
		return repository.Deductions{}, err
	}
	return deductions, nil
}
