package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/app/emp/mocks"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/dto"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestLoginHandler(t *testing.T) {
	empSvc := mocks.NewService(t)
	empLoginSvc := LoginHandler(empSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for login",
			input: `{
						"email": "sharanyadatrange1@gmail.com",
						"password": "12345"   
					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("Login", "sharanyadatrange1@gmail.com", "12345").Return("token", nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail Invalid json",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(empSvc)

			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(empLoginSvc)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestCreateEmployeeHandler(t *testing.T) {
	empSvc := mocks.NewService(t)
	empCreateHandler := CreateEmployeeHandler(empSvc)

	test := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for creating employee",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary" : 200000.0,
				"password":"12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail Invalid JSON",
			input:              "",
			setup:              func(mock *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with missing required fields",
			input: `{
				"id":"",
				"first_name":"",
				"middle_name": "sachin",
				"last_name":"",
				"email":"",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary" : 200000.0,
				"password":"12345"

			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, errors.New("Error ")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with invalid salary",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary": -200000.0,
				"password": "12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, errors.New("Error ")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with invalid email",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "invalid",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "bangalore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary": 200000.0,
				"password": "12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, errors.New("Error ")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with future date of birth",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2025-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "bangalore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary": 200000.0,
				"password": "12345"
			}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with future date of joining",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2025-05-01T09:00:00Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary": 200000.0,
				"password": "12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, errors.New("Error ")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with invalid work status",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "on leave",
				"salary": 200000.0,
				"password": "12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, errors.New("Error ")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with empty proof ID",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary": 200000.0,
				"password": "12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, errors.New("Error ")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with invalid role ID",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": 1,
				"proof_id": "54236",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": ,
				"work_status": "active",
				"salary": 200000.0,
				"password": "12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail with negative years of experience",
			input: `{
				"id": "S116",
				"first_name": "sachin",
				"middle_name": "sachin",
				"last_name": "joshi",
				"email": "ssd174285@gmail.com",
				"date_of_birth": "2002-10-06T15:04:45Z",
				"date_of_joining": "2023-01-10T15:04:45Z",
				"designation": "Employee",
				"years_of_experience": -1,
				"proof_id": "54236",
				"residential_address": "banglore",
				"hired_location": "pune",
				"role_id": 1,
				"work_status": "active",
				"salary": 200000.0,
				"password": "12345"
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateEmployee", mock.AnythingOfType("dto.CreateEmployeeRequest")).Return(repository.Employee{
					ID: "S115",
				}, errors.New("Error ")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			test.setup(empSvc)

			req, err := http.NewRequest("POST", "/createEmployee", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(empCreateHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetEmployeeByIDHandler(t *testing.T) {
	empSvc := mocks.NewService(t)
	empGetByIdHandler := GetEmployeeByIDHandler(empSvc)

	test := []struct {
		name               string
		id                 string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for Get employee By ID ",
			id:   "S116",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetEmployeeByID", mock.Anything).Return(dto.Employee{}, nil).Once()

			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			test.setup(empSvc)

			req, err := http.NewRequest("GET", fmt.Sprintf(
				"/getEmployeeByID?id=%s",
				test.id,
			), bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(empGetByIdHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetAllEmployeesHandler(t *testing.T) {
	empSvc := mocks.NewService(t)
	empGetAllHandler := GetAllEmployeesHandler(empSvc)

	test := []struct {
		name               string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetAllEmployees").Return([]dto.Employee{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			test.setup(empSvc)

			req, err := http.NewRequest("GET", "/getAllEmployees", nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(empGetAllHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestDeleteEmployeeHandler(t *testing.T) {
	empSvc := mocks.NewService(t)
	empDeleteHandler := DeleteEmployeeHandler(empSvc)

	test := []struct {
		name               string
		id                 string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for deleting employee",
			id:   "S116",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteEmployee", "S116").Return(dto.Employee{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			test.setup(empSvc)

			req, err := http.NewRequest("GET", fmt.Sprintf(
				"/deleteEmployee?id=%s",
				test.id,
			), bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(empDeleteHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
