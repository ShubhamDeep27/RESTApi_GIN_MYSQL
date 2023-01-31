package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	dao "rest/gin/Dao"
	models "rest/gin/models"
	"strconv"
	"sync"
	"rest/gin/util"
)

type EmployeeService interface {
	GetAllEmployees() ([]*models.Employee, error)
	CreateEmployees(employee *models.CreateEmployee) (err error)
	GetEmployeeById(id string) (*models.Employee, error)
	UpdateEmployee(employee *models.Employee, id string) (err error)
	CreateEmployeesImaginary() ([]map[string]interface{}, error)
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

func (es *EmployeeServiceImpl) CreateEmployees(employee *models.CreateEmployee) error {

	err := util.EmployeeValidation(employee)
	if err != nil {
		return err
	}
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

var wg sync.WaitGroup

func (es *EmployeeServiceImpl) CreateEmployeesImaginary() ([]map[string]interface{}, error) {

	employee, err := es.GetAllEmployees()
	if err != nil {
		return nil, err
	}
	var data []map[string]interface{}
	for _, emp := range employee {
		wg.Add(1)

		go func(emp models.Employee) {
			resp, err := PostReq(emp)
			if err != nil {
				return
			}

			data = append(data, resp)
		}(*emp)

	}
	wg.Wait()
	// fmt.Println(data)
	return data, nil
}
func PostReq(emp models.Employee) (map[string]interface{}, error) {

	defer wg.Done()
	values := map[string]string{"name": emp.Name, "salary": strconv.Itoa(emp.Salary), "age": strconv.Itoa(emp.Age)}

	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewBuffer(json_data)

	resp, err := http.Post("https://dummy.restapiexample.com/api/v1/create", "application/json", body)
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}
