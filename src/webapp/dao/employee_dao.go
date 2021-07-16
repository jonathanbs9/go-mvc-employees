package dao

import (
	"fmt"
	"net/http"

	"github.com/jonathanbs9/go-mvc-employees/src/webapp/errors"
	"github.com/jonathanbs9/go-mvc-employees/src/webapp/model"
)

var (
	employees = map[int64]*model.Employee{
		100: {Id: 101, Name: "Jonathan Brull Schroeder", Email: "jonathans@avalith.net"},
	}
)

func GetEmployee(empId int64) (*model.Employee, *errors.AppError) {
	if employee := employees[empId]; employee != nil {
		return employee, nil
	}
	// errors.New(fmt.Sprintf("Employee %v NOT FOUND!", empId))
	return nil, &errors.AppError{
		Message:    fmt.Sprintf("Employee %v NOT FOUND!", empId),
		StatusCode: http.StatusNotFound,
		Status:     "not_found",
	}
}
