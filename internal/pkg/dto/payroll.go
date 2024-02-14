package dto

import "time"

type Payroll struct {
	ID           string    `json:"id"`
	Salary       float64   `json:"salary"`
	NetPaySalary float64   `json:"net_pay_salary"`
	PayDate      time.Time `json:"pay_date"`
}

type CreatePayrollRequest struct {
	ID           string    `json:"id"`
	Salary       float64   `json:"salary"`
	NetPaySalary float64   `json:"net_pay_salary"`
	PayDate      time.Time `json:"pay_date"`
}
