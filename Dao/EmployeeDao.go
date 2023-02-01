package dao

import (
	"rest/gin/Config"
	"rest/gin/models"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeDao interface {
	GetAllEmployees() ([]*models.Employee, error)
	CreateEmployees(employee *models.CreateEmployee) error
	GetEmployeeById(id string) (*models.Employee, error)
	UpdateEmployee(employee *models.CreateEmployee, id string) error
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

func (ed *EmployeeDaoImpl) CreateEmployees(employee *models.CreateEmployee) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), 8)
	if err != nil {
		return err
	}
	employee.Password = string(hashedPassword)
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

func (ed *EmployeeDaoImpl) UpdateEmployee(employee *models.CreateEmployee, id string) (err error) {
	empid, _ := strconv.ParseInt(id, 10, 32)
	employeeToEdit := models.Employee{
		ID:       empid,
		Name:     employee.Name,
		Mobile:   employee.Mobile,
		Address:  employee.Address,
		Age:      employee.Age,
		Salary:   employee.Salary,
		Username: employee.Username,
		Password: employee.Password,
	}

	if err = Config.DB.Save(&employeeToEdit).Error; err != nil {
		return err
	}
	return nil
}
