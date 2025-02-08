package repository

import (
	"skillz_manager/model"
)

type IEmployeeRepository interface {
	FindByName(name string) (model.Employee, error)
	DeleteEmployee(employee model.Employee) (error, int)
}

type EmployeeRepository struct {
}

func (employeeRepository EmployeeRepository) FindByName(name string) (model.Employee, error) {
	//Implement code
	return model.Employee{}, nil
}

func (employeeRepository EmployeeRepository) DeleteEmployee(employee model.Employee) (error, int) {
	//Implement code
	return nil, -1
}
