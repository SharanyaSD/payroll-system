package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/app/payroll"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/dto"
)

func CreatePayrollHandler(payrollSvc payroll.PayrollService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreatePayrollRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("%+v", req)

		payrollInfo, err := payrollSvc.CreatePayroll(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		responseJSON, err := json.Marshal(payrollInfo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		//return
	})
}

func GetPayrollHandler(payrollSvc payroll.PayrollService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := payrollSvc.GetPayroll()
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
		//return
	}
}
