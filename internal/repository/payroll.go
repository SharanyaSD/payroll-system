package repository

import "time"

type Payroll struct {
	ID           string    `db:"id"`
	Salary       float64   `db:"salary"`
	NetPaySalary float64   `db:"net_pay_salary"`
	PayDate      time.Time `db:"pay_date"`
}
