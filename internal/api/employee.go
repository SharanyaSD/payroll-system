package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/app/emp"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/dto"
)

func CreateEmployeeHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.CreateEmployeeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("%+v", req)

		if req.DateOfBirth.After(time.Now()) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Date of birth cannot be in the future"))
			return
		}

		employeeInfo, err := empSvc.CreateEmployee(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		responseJSON, err := json.Marshal(employeeInfo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}
}

func DeleteEmployeeHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		empID := r.URL.Query().Get("id")
		if empID == "" {
			http.Error(w, "Employee ID is required", http.StatusBadRequest)
			return
		}

		deletedEmp, err := empSvc.DeleteEmployee(empID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(deletedEmp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Employee deleted successfully"))
		w.Write(jsonData)
	}
}

func GetAllEmployeesHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := empSvc.GetAllEmployees()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	}
}

func GetEmployeeByIDHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if id == "" {
			http.Error(w, "No ID provided", http.StatusBadRequest)
			return
		}

		employee, err := empSvc.GetEmployeeByID(id)
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(err.Error()))
			// return
			http.Error(w, "Failed to get employee: "+err.Error(), http.StatusNotFound)
			return
		}

		responseJSON, err := json.Marshal(employee)
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(err.Error()))
			// return
			http.Error(w, "Failed to serialize employee to JSON: "+err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	}
}

func LoginHandler(empSvc emp.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials emp.Credentials
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		token, err := empSvc.Login(credentials.Email, credentials.Password)
		if err != nil {
			if errors.Is(err, emp.ErrEmailNotFound) {
				http.Error(w, "Email not found", http.StatusNotFound)
				return
			}

			http.Error(w, "Failed to login: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with token
		responseJSON, err := json.Marshal(map[string]string{"token": token})
		if err != nil {
			http.Error(w, "Failed to serialize login info to JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}
}
