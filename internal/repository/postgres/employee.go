package repository

import (
	"fmt"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type EmployeeStore struct {
	Db *sqlx.DB
}

func NewEmployeeRepo(db *sqlx.DB) repository.EmployeeStorer {
	return &EmployeeStore{
		Db: db,
	}

}

func (e *EmployeeStore) GetEmployeeByID(ID string) (repository.Employee, error) {
	var emp repository.Employee
	query := "SELECT * FROM Employees WHERE id=$1"
	fmt.Println("SQL Query:", query, ID)
	row := e.Db.QueryRow(query, ID)
	err := row.Scan(
		&emp.ID, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Email, &emp.DateOfBirth,
		&emp.DateOfJoining, &emp.Designation, &emp.YearsOfExperience, &emp.ProofId, &emp.ResidentialAddress,
		&emp.HiredLocation, &emp.RoleId, &emp.WorkStatus, &emp.Salary, &emp.Password,
	)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

func (e *EmployeeStore) GetEmployeeByEmail(email string) (repository.Employee, error) {
	var emp repository.Employee
	query := "SELECT * FROM Employees WHERE email=$1"
	fmt.Println("SQL Query:", query, email)
	row := e.Db.QueryRow(query, email)
	err := row.Scan(
		&emp.ID, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Email, &emp.DateOfBirth,
		&emp.DateOfJoining, &emp.Designation, &emp.YearsOfExperience, &emp.ProofId, &emp.ResidentialAddress,
		&emp.HiredLocation, &emp.RoleId, &emp.WorkStatus, &emp.Salary, &emp.Password,
	)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

// func (e *EmployeeStore) CreateEmployee(emp repository.Employee) (repository.Employee, error) {
// 	err := e.Db.QueryRow("insert into Employees (id , first_name , middle_name , last_name ,email , date_of_birth , date_of_joining , designation , years_of_experience , proof_id , residential_address , hired_location  , work_status) values ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10, $11,$12,$13)", emp.ID, emp.FirstName, emp, emp.MiddleName, emp.LastName, emp.Email, emp.DateOfBirth, emp.DateOfJoining, emp.Designation, emp.YearsOfExperience, emp.ProofId, emp.ResidentialAddress, emp.HiredLocation, emp.WorkStatus).Scan(&emp)
// 	if err != nil {
// 		return repository.Employee{}, err
// 	}

// 	return emp, nil
// }

func (e *EmployeeStore) CreateEmployee(emp repository.Employee) (repository.Employee, error) {
	_, err := e.Db.Exec("INSERT INTO Employees (id, first_name, middle_name, last_name, email, date_of_birth, date_of_joining, designation, years_of_experience, proof_id, residential_address, hired_location, role_id, work_status, salary, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)",
		emp.ID, emp.FirstName, emp.MiddleName, emp.LastName, emp.Email, emp.DateOfBirth, emp.DateOfJoining, emp.Designation, emp.YearsOfExperience, emp.ProofId, emp.ResidentialAddress, emp.HiredLocation, emp.RoleId, emp.WorkStatus, emp.Salary, emp.Password)

	if err != nil {
		return repository.Employee{}, err
	}

	return emp, nil
}

func (e *EmployeeStore) GetAllEmployees() ([]repository.Employee, error) {
	rows, err := e.Db.Query("SELECT * FROM Employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emps []repository.Employee
	for rows.Next() {
		var emp repository.Employee
		if err := rows.Scan(&emp.ID, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Email, &emp.DateOfBirth, &emp.DateOfJoining, &emp.Designation, &emp.YearsOfExperience, &emp.ProofId, &emp.ResidentialAddress, &emp.HiredLocation, &emp.RoleId, &emp.WorkStatus, &emp.Salary, &emp.Password); err != nil {
			return nil, err
		}
		emps = append(emps, emp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return emps, nil
}

// func (e *EmployeeStore) UpdateEmployee(emp repository.Employee) (repository.Employee, error) {

// }

// 	return e.Db.QueryRow("update Employees set first_name='neha' where id=2",emp.ID).Scan(&emp)
// }

func (e *EmployeeStore) DeleteEmployee(ID string) (repository.Employee, error) {
	var emp repository.Employee
	_, err := e.Db.Exec("DELETE FROm Employees WHERE id = $1", ID)
	if err != nil {
		return emp, err
	}
	return emp, nil
}
