package Models

import (
	"rest/gin/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllEmployees(employee *[]Employee) (err error) {
	if err = Config.DB.Find(&employee).Error; err != nil {
		return err
	}
	return nil

}

func CreateEmployees(employee *Employee) (err error) {

	if err = Config.DB.Create(&employee).Error; err != nil {
		return err
	}

	return nil
}

func GetEmployeeById(employee *Employee, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(employee).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(employee *Employee, id string) (err error) {
	if err = Config.DB.Save(employee).Error; err != nil {
		return err
	}
	return nil
}
