package dao

import (
	"rest/gin/Config"
	"rest/gin/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllEmployees(employee *[]models.Employee) (err error) {
	if err = Config.DB.Find(&employee).Error; err != nil {
		return err
	}
	return nil

}

func CreateEmployees(employee *models.Employee) (err error) {

	if err = Config.DB.Create(&employee).Error; err != nil {
		return err
	}

	return nil
}

func GetEmployeeById(employee *models.Employee, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(employee).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(employee *models.Employee, id string) (err error) {
	if err = Config.DB.Save(employee).Error; err != nil {
		return err
	}
	return nil
}
