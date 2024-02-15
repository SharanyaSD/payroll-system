package payroll

import (
	"errors"
	"fmt"
	"time"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/app/emp"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/dto"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/email"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
)

// dependencies
type payrollService struct {
	payrollRepo     repository.PayrollStorer
	employeeService emp.Service
	emailService    email.Service
}

// initialize services
func NewPayrollService(payrollRepo repository.PayrollStorer, employeeService emp.Service, emailService email.Service) PayrollService {
	return &payrollService{
		payrollRepo:     payrollRepo,
		employeeService: employeeService,
		emailService:    emailService,
	}
}

type PayrollService interface {
	CreatePayroll(payrollDetails dto.CreatePayrollRequest) (repository.Payroll, error)
	GetPayroll() ([]dto.Payroll, error)

	//GetPayrollByID(payroll_id string) (dto.Payroll, error)
}

func (ps *payrollService) CreatePayroll(payrollDetails dto.CreatePayrollRequest) (repository.Payroll, error) {
	employee, err := ps.employeeService.GetEmployeeByID(payrollDetails.ID)
	if err != nil {
		return repository.Payroll{}, err
	}

	//check date of salary
	payDate := time.Now()
	if payDate.Day() < 5 {
		return repository.Payroll{}, fmt.Errorf("CANNOT BE GENERATED BEFORE 5TH")
	}

	salary := employee.Salary

	earnings, err := ps.employeeService.GetEarningsByEmpoyeeID(payrollDetails.ID)
	if err != nil {
		return repository.Payroll{}, err
	}

	deductions, err := ps.employeeService.GetDeductionsByEmpoyeeID(payrollDetails.ID)
	if err != nil {
		return repository.Payroll{}, err
	}

	netPay := calculateNetPay(salary, earnings, deductions)

	payroll := repository.Payroll{
		ID:           payrollDetails.ID,
		Salary:       salary,
		NetPaySalary: netPay,
		PayDate:      payDate,
	}

	_, err = ps.payrollRepo.CreatePayroll(payroll)
	if err != nil {
		return repository.Payroll{}, err
	}

	//email msg
	emailContent := fmt.Sprintf("Dear %s,\n Your payroll for the month has been generated. \n\n Salary: %.2f\nNet Pay: %.2f\n\nBest regards,\nPayroll Team", employee.FirstName, payroll.Salary, payroll.NetPaySalary)

	//send email
	err = ps.emailService.SendEmail("sharanyadatrange1@gmail.com", "Payroll for the month", emailContent)
	if err != nil {
		return repository.Payroll{}, err
	}

	return payroll, nil

}

func calculateNetPay(salary float64, earnings repository.Earnings, deductions repository.Deductions) float64 {
	netPay := (salary + earnings.GrossPay) - deductions.GrossDeduction
	return netPay
}

func (ps *payrollService) GetPayroll() ([]dto.Payroll, error) {
	//date of checking payroll
	now := time.Now()

	//extracting current year and month from Date()
	currentYear, currentMonth, _ := now.Date()

	// creates a new time.Time instance - year, month, day, hour, min, sec, nanosec, location
	payDate := time.Date(currentYear, currentMonth, 5, 0, 0, 0, 0, now.Location())

	if now.Before(payDate) {
		payDate = payDate.AddDate(0, -1, 0)
	}

	if currentYear < payDate.Year() || (currentYear == payDate.Year() && currentMonth < payDate.Month()) {
		return nil, errors.New("PAY daate is in future")
	}

	// payrolls, err := ps.payrollRepo.GetPayroll()
	payrolls, err := ps.payrollRepo.GetPayroll()
	if err != nil {
		return []dto.Payroll{}, err
	}
	var dtoPayrolls []dto.Payroll
	for _, payroll := range payrolls {
		if payroll.PayDate.After(payDate) {
			// In go - magic date     Mon Jan 2 15:04:05 2006 -0700 MST looks like 1, 1, 2, 3, 4, 5, 6, 7
			fmt.Printf("Salary Date: %s\n", payroll.PayDate.Format("January 2, 2006"))
			dtoPayroll := dto.Payroll{
				ID:           payroll.ID,
				Salary:       payroll.Salary,
				NetPaySalary: payroll.NetPaySalary,
				PayDate:      payroll.PayDate,
			}

			dtoPayrolls = append(dtoPayrolls, dtoPayroll)
		}
	}
	return dtoPayrolls, nil
}
