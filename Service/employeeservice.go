package service

import (
	"errors"
	dao "rest/gin/Dao"
	models "rest/gin/models"
)

type EmployeeService interface {
	GetAllEmployees(employee *[]models.Employee) (err error)
	CreateEmployees(employee *models.Employee) (err error)
	GetEmployeeById(employee *models.Employee, id string) (err error)
	UpdateEmployee(employee *models.Employee, id string) (err error)
}

type EmployeeServiceImpl struct {
	empdao dao.EmployeeDao
}

func NewEmployeeServiceImpl(dao dao.EmployeeDao) EmployeeService {
	return &EmployeeServiceImpl{empdao: dao}
}

func (es *EmployeeServiceImpl) GetAllEmployees(employee *[]models.Employee) (err error) {

	err = es.empdao.GetAllEmployees(employee)
	if err != nil {
		return err
	}
	return nil

}

func (es *EmployeeServiceImpl) CreateEmployees(employee *models.Employee) (err error) {
	err = es.empdao.CreateEmployees(employee)
	if err != nil {
		return err

	}
	return nil

}

func (es *EmployeeServiceImpl) GetEmployeeById(employee *models.Employee, id string) (err error) {
	err = es.empdao.GetEmployeeById(employee, id)
	if err != nil {
		return errors.New("record not found")
	}
	return nil

}

func (es *EmployeeServiceImpl) UpdateEmployee(employee *models.Employee, id string) (err error) {
	var tempemployee models.Employee
	err = es.empdao.GetEmployeeById(&tempemployee, id)
	if err != nil {
		return err
	}
	err = es.empdao.UpdateEmployee(employee, id)
	if err != nil {
		return err
	}

	return nil

}
