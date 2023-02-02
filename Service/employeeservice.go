package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"rest/gin/dao"
	"rest/gin/models"
	"rest/gin/util"
	"strconv"
	"sync"
)

type EmployeeService interface {
	GetAllEmployees() ([]*models.Employee, error)
	CreateEmployees(employee *models.CreateEmployee) models.ErrorResponse
	GetEmployeeById(id string) (*models.Employee, error)
	UpdateEmployee(employee *models.CreateEmployee, id string) models.ErrorResponse
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

func (es *EmployeeServiceImpl) CreateEmployees(employee *models.CreateEmployee) models.ErrorResponse {

	err := util.EmployeeValidation(employee)
	if err != (models.ErrorResponse{}) {
		return err
	}
	errs := es.empdao.CreateEmployees(employee)
	if errs != nil {
		return models.ErrorResponse{ErrorCode: "503", ErrorMsg: util.GET_ERROR_MSGS["503"]}

	}
	return models.ErrorResponse{}

}

func (es *EmployeeServiceImpl) GetEmployeeById(id string) (*models.Employee, error) {

	employee, err := es.empdao.GetEmployeeById(id)
	if err != nil {
		return employee, err
	}
	return employee, nil

}

func (es *EmployeeServiceImpl) UpdateEmployee(employee *models.CreateEmployee, id string) models.ErrorResponse {
	_, err := es.empdao.GetEmployeeById(id)
	if err != nil {
		return models.ErrorResponse{ErrorCode: "502", ErrorMsg: util.GET_ERROR_MSGS["501"]}
	}
	errVal := util.EmployeeValidation(employee)
	if errVal != (models.ErrorResponse{}) {
		return errVal
	}

	err = es.empdao.UpdateEmployee(employee, id)
	if err != nil {
		return models.ErrorResponse{ErrorCode: "507", ErrorMsg: util.GET_ERROR_MSGS["507"]}
	}

	return models.ErrorResponse{}

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
