package service

import (
	dao "rest/gin/dao"
	models "rest/gin/models"
)

type EmployeeService interface {
	GetAllEmployees() ([]*models.Employee, error)
	CreateEmployees(employee *models.Employee) (err error)
	GetEmployeeById(id string) (*models.Employee, error)
	UpdateEmployee(employee *models.Employee, id string) (err error)
}

type EmployeeServiceImpl struct {
	empdao dao.EmployeeDao
}

func NewEmployeeServiceImpl(dao dao.EmployeeDao) EmployeeService {
	return &EmployeeServiceImpl{empdao: dao}
}

func (es *EmployeeServiceImpl) GetAllEmployees() ([]*models.Employee, error) {

	employee, err := es.empdao.GetAllEmployees()
	if err != nil {
		return employee, err
	}
	return employee, nil

}

func (es *EmployeeServiceImpl) CreateEmployees(employee *models.Employee) (err error) {
	err = es.empdao.CreateEmployees(employee)
	if err != nil {
		return err

	}
	return nil

}

func (es *EmployeeServiceImpl) GetEmployeeById(id string) (*models.Employee, error) {

	employee, err := es.empdao.GetEmployeeById(id)
	if err != nil {
		return employee, err
	}
	return employee, nil

}

func (es *EmployeeServiceImpl) UpdateEmployee(employee *models.Employee, id string) error {
	_, err := es.empdao.GetEmployeeById(id)
	if err != nil {
		return err
	}
	err = es.empdao.UpdateEmployee(employee, id)
	if err != nil {
		return err
	}

	return nil

}
