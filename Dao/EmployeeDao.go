package dao

import (
	"rest/gin/Config"
	"rest/gin/models"

	_ "github.com/go-sql-driver/mysql"
)
type EmployeeDao struct {
}

func NewEmployeeDao() *EmployeeDao {
	return &EmployeeDao{}
}

func (ed *EmployeeDao)GetAllEmployees(employee *[]models.Employee) (err error) {
	if err = Config.DB.Find(&employee).Error; err != nil {
		return err
	}
	return nil

}

func (ed *EmployeeDao)CreateEmployees(employee *models.Employee) (err error) {

	if err = Config.DB.Create(&employee).Error; err != nil {
		return err
	}

	return nil
}

func (ed *EmployeeDao)GetEmployeeById(employee *models.Employee, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(employee).Error; err != nil {
		return err
	}
	return nil
}

func (ed *EmployeeDao)UpdateEmployee(employee *models.Employee, id string) (err error) {
	if err = Config.DB.Save(employee).Error; err != nil {
		return err
	}
	return nil
}
