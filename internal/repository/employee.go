package repository

import "time"

//import "github.com/google/uuid"

type Employee struct {
	ID                 string    `db:"id"`
	FirstName          string    `db:"first_name"`
	MiddleName         string    `db:"middle_name"`
	LastName           string    `db:"last_name"`
	Email              string    `db:"email"`
	DateOfBirth        time.Time `db:"date_of_birth"`
	DateOfJoining      time.Time `db:"date_of_joining"`
	Designation        string    `db:"designation"`
	YearsOfExperience  int       `db:"years_of_experience"`
	ProofId            string    ` db:"proof_id"`
	ResidentialAddress string    `db:"residential_address"`
	HiredLocation      string    ` db:"hired_location"`
	RoleId             int       ` db:"role_id"`
	WorkStatus         string    `db:"work_status"`
	Salary             float64   `db:"salary"`
	Password           string    `db:"password"`
}
