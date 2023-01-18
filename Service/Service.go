package service

import (
	dao "rest/gin/Dao"
	models "rest/gin/models"
)

type EmployeeService struct {
	empdao *dao.EmployeeDao
}

func NewEmployeeService(dao *dao.EmployeeDao) *EmployeeService {
	return &EmployeeService{empdao: dao}
}

func (es *EmployeeService) GetAllEmployees(employee *[]models.Employee) (err error) {
	err = es.empdao.GetAllEmployees(employee)
	if err != nil {
		return err
	}
	return nil

}

func (es *EmployeeService) CreateEmployees(employee *models.Employee) (err error) {
	err = es.empdao.CreateEmployees(employee)
	if err != nil {
		return err

	}
	return nil

}

func (es *EmployeeService) GetEmployeeById(employee *models.Employee, id string) (err error) {
	err = es.empdao.GetEmployeeById(employee, id)
	if err != nil {
		return err
	}
	return nil

}

func (es *EmployeeService) UpdateEmployee(employee *models.Employee, id string) (err error) {
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
