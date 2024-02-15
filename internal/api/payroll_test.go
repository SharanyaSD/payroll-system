package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/app/payroll/mocks"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/dto"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestCreatePayrollHandler(t *testing.T) {
	payrollSvc := mocks.NewPayrollService(t)
	payrollCreateHandler := CreatePayrollHandler(payrollSvc)

	test := []struct {
		name               string
		input              string
		setup              func(mock *mocks.PayrollService)
		expectedStatusCode int
	}{
		{
			name: "Success for creating payroll",
			input: `{
				"id": "S116"
			}`,
			setup: func(mockSvc *mocks.PayrollService) {
				mockSvc.On("CreatePayroll", mock.AnythingOfType("dto.CreatePayrollRequest")).Return(repository.Payroll{
					ID: "S115",
				}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			test.setup(payrollSvc)

			req, err := http.NewRequest("POST", "/createPayroll", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(payrollCreateHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetPayrollHandler(t *testing.T) {
	payrollSvc := mocks.NewPayrollService(t)
	payrollGetAll := GetPayrollHandler(payrollSvc)
	test := []struct {
		name               string
		setup              func(mock *mocks.PayrollService)
		expectedStatusCode int
	}{
		{
			name: "Success",
			setup: func(mockSvc *mocks.PayrollService) {
				mockSvc.On("GetPayroll").Return([]dto.Payroll{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			test.setup(payrollSvc)

			req, err := http.NewRequest("GET", "/getPayroll", nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(payrollGetAll)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected status code %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
