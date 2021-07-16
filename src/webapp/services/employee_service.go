package services

import (
	"github.com/jonathanbs9/go-mvc-employees/src/webapp/dao"
	"github.com/jonathanbs9/go-mvc-employees/src/webapp/errors"
	"github.com/jonathanbs9/go-mvc-employees/src/webapp/model"
)

func GetEmployee(empId int64) (*model.Employee, *errors.AppError) {
	return dao.GetEmployee(empId)

}
