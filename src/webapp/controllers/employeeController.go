package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jonathanbs9/go-mvc-employees/src/webapp/errors"
	"github.com/jonathanbs9/go-mvc-employees/src/webapp/services"
)

func GetEmployee(response http.ResponseWriter, req *http.Request) {

	empId, err := strconv.ParseInt(req.URL.Query().Get("employee_id"), 10, 64)
	if err != nil {
		apiErr := &errors.AppError{
			Message:    "employee_id must be number (int)",
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	emp, apiErr := services.GetEmployee(empId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write([]byte(jsonValue))
		return
	}

	jsonVal, _ := json.Marshal(emp)
	response.WriteHeader(http.StatusOK)
	response.Write(jsonVal)

}
