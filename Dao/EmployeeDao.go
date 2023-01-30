package dao

import (
	"log"
	Config "rest/gin/config"
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
	var employeelist []*models.Employee
	rows, err := Config.DB.Query("SELECT * FROM employee")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var employee models.Employee
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Mobile, &employee.Address, &employee.Salary, &employee.Age, &employee.Username, &employee.Password)
		if err != nil {
			return nil, err
		} else {
			employeelist = append(employeelist, &employee)
		}
	}
	return employeelist, nil
}

func (ed *EmployeeDaoImpl) CreateEmployees(emp *models.Employee) (err error) {

	_, err = Config.DB.Exec("INSERT INTO employee(id,name,mobile,address,salary,age,username,password)VALUES(?,?,?,?,?,?,?,?)", emp.ID, emp.Name, emp.Mobile, emp.Address, emp.Salary, emp.Age, emp.Username, emp.Password)
	if err != nil {
		return err
	}
	return nil

}

func (ed *EmployeeDaoImpl) GetEmployeeById(id string) (*models.Employee, error) {
	var employee models.Employee

	row := Config.DB.QueryRow("SELECT id,name,mobile,address,salary,age,username,password FROM employee WHERE id = ? LIMIT 1", id)

	err := row.Scan(&employee.ID, &employee.Name, &employee.Mobile, &employee.Address, &employee.Salary, &employee.Age, &employee.Username, &employee.Password)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (ed *EmployeeDaoImpl) UpdateEmployee(employee *models.Employee, id string) (err error) {
	_, err = Config.DB.Exec("UPDATE employee SET id = ?, name = ?, mobile = ?,address = ?,salary = ?,age = ?,username = ?,password = ? WHERE id = ?", &employee.ID, &employee.Name, &employee.Mobile, &employee.Address, &employee.Salary, &employee.Age, &employee.Username, &employee.Password, id)
	if err != nil {
		return err
	}
	return nil
}
