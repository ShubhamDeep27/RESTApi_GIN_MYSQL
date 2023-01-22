package dao

import (
	"rest/gin/Config"
	"rest/gin/models"

	_ "github.com/go-sql-driver/mysql"
)

type EmployeeDao interface {
	GetAllEmployees() ([]*models.Employee, error)
	CreateEmployees(employee *models.Employee) error
	GetEmployeeById(id string) (*models.Employee, error)
	UpdateEmployee(employee *models.Employee, id string) error
}

type EmployeeDaoImpl struct {
}

func NewEmployeeDaoImpl() *EmployeeDaoImpl {
	return &EmployeeDaoImpl{}
}

func (ed *EmployeeDaoImpl) GetAllEmployees() ([]*models.Employee, error) {
	var employee []*models.Employee
	if err := Config.DB.Find(&employee).Error; err != nil {
		return employee, err
	}
	return employee, nil

}

func (ed *EmployeeDaoImpl) CreateEmployees(employee *models.Employee) (err error) {

	if err = Config.DB.Create(&employee).Error; err != nil {
		return err
	}

	return nil
}

func (ed *EmployeeDaoImpl) GetEmployeeById(id string) (*models.Employee, error) {
	var employee *models.Employee
	if err := Config.DB.Where("id = ?", id).First(&employee).Error; err != nil {
		return employee, err
	}
	return employee, nil
}

func (ed *EmployeeDaoImpl) UpdateEmployee(employee *models.Employee, id string) (err error) {
	if err = Config.DB.Save(employee).Error; err != nil {
		return err
	}
	return nil
}
