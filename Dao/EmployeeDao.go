package dao

import (
	"rest/gin/Config"
	"rest/gin/models"

	_ "github.com/go-sql-driver/mysql"
)

type EmployeeDao interface {
	GetAllEmployees(employee *[]models.Employee) (err error)
	CreateEmployees(employee *models.Employee) (err error)
	GetEmployeeById(employee *models.Employee, id string) (err error)
	UpdateEmployee(employee *models.Employee, id string) (err error)
}

type EmployeeDaoImpl struct {
}

func NewEmployeeDaoImpl() *EmployeeDaoImpl {
	return &EmployeeDaoImpl{}
}

func (ed *EmployeeDaoImpl) GetAllEmployees(employee *[]models.Employee) (err error) {
	if err = Config.DB.Find(&employee).Error; err != nil {
		return err
	}
	return nil

}

func (ed *EmployeeDaoImpl) CreateEmployees(employee *models.Employee) (err error) {

	if err = Config.DB.Create(&employee).Error; err != nil {
		return err
	}

	return nil
}

func (ed *EmployeeDaoImpl) GetEmployeeById(employee *models.Employee, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(employee).Error; err != nil {
		return err
	}
	return nil
}

func (ed *EmployeeDaoImpl) UpdateEmployee(employee *models.Employee, id string) (err error) {
	if err = Config.DB.Save(employee).Error; err != nil {
		return err
	}
	return nil
}
