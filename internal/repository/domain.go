package repository

type EmployeeStorer interface {
	GetAllEmployees() ([]Employee, error)
	GetEmployeeByID(ID string) (Employee, error)
	CreateEmployee(emp Employee) (Employee, error)
	DeleteEmployee(ID string) (Employee, error)
	GetEmployeeByEmail(email string) (Employee, error)
}

type PayrollStorer interface {
	CreatePayroll(payroll Payroll) (Payroll, error)
	GetPayroll() ([]Payroll, error)
}

type EarningsStorer interface {
	GetEarningsByEmpoyeeID(ID string) (Earnings, error)
	InsertEarnings(earnings Earnings) (Earnings, error)
}

type DeductionsStorer interface {
	GetDeductionsByEmpoyeeID(ID string) (Deductions, error)
	InsertDeductions(deductions Deductions) (Deductions, error)
}

type Earnings struct {
	ID       string  `db:"id"`
	Basic    float64 `db:"basic"`
	HRA      float64 `db:"hra"`
	DA       float64 `db:"da"`
	SA       float64 `db:"sa"`
	CA       float64 `db:"ca"`
	Bonus    float64 `db:"bonus"`
	GrossPay float64 `db:"gross_pay"`
}

type Deductions struct {
	ID             string  `db:"id"`
	TDS            float64 `db:"tds"`
	PF             float64 `db:"pf"`
	Medical        float64 `db:"medical"`
	GrossDeduction float64 `db:"gross_deduction"`
	//AdditionalTaxes map[string]float64 // AdditionalTaxes should be initialized to avoid nil map panic
}
