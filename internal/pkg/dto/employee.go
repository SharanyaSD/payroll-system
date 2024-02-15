package dto

import "time"

type Employee struct {
	ID                 string    `json:"id"`
	FirstName          string    `json:"first_name"`
	MiddleName         string    `json:"middle_name"`
	LastName           string    `json:"last_name"`
	Email              string    `json:"email"`
	DateOfBirth        time.Time `json:"date_of_birth"`
	DateOfJoining      time.Time `json:"date_of_joining"`
	Designation        string    `json:"designation"`
	YearsOfExperience  int       `json:"years_of_experience"`
	ProofId            string    `json:"proof_id"`
	ResidentialAddress string    `json:"residential_address"`
	HiredLocation      string    `json:"hired_location"`
	RoleId             int       `json:"role_id"`
	WorkStatus         string    `json:"work_status"`
	Salary             float64   `json:"salary"`
	Password           string    `json:"-"`
}

type CreateEmployeeRequest struct {
	ID                 string    `json:"id"`
	FirstName          string    `json:"first_name"`
	MiddleName         string    `json:"middle_name"`
	LastName           string    `json:"last_name"`
	Email              string    `json:"email"`
	DateOfBirth        time.Time `json:"date_of_birth"`
	DateOfJoining      time.Time `json:"date_of_joining"`
	Designation        string    `json:"designation"`
	YearsOfExperience  int       `json:"years_of_experience"`
	ProofId            string    `json:"proof_id"`
	ResidentialAddress string    `json:"residential_address"`
	RoleId             int       `json:"role_id"`
	HiredLocation      string    `json:"hired_location"`

	WorkStatus string  `json:"work_status"`
	Salary     float64 `json:"salary"`
	Password   string  `json:"password"`
}

// type UpdateEmployeeRequest struct {

// }
