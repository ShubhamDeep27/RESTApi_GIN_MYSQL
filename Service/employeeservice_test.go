package service

import (
	"errors"
	models "rest/gin/models"
	"strconv"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetEmployess_Success(t *testing.T) {

	employeeList := []*models.Employee{
		{ID: 2,
			Name:     "Ram",
			Mobile:   "78589535",
			Address:  "Delhi",
			Salary:   34566,
			Age:      23,
			Username: "ram",
			Password: "123",
		}, {
			ID:       3,
			Name:     "Raj",
			Mobile:   "7854589535",
			Address:  "UP",
			Salary:   36566,
			Age:      21,
			Username: "raj",
			Password: "1234",
		},
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)

	mockworker.EXPECT().GetAllEmployees().Times(1).Return(employeeList, nil)

	actualEmplist, err := empService.GetAllEmployees()
	assert.Equal(t, employeeList, actualEmplist)
	assert.Nil(t, err)
}
func TestGetEmployess_ErrorWhileGetting_AllEmployee(t *testing.T) {

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)

	mockworker.EXPECT().GetAllEmployees().Times(1).Return(nil, errors.New(""))

	actualEmplist, err := empService.GetAllEmployees()
	assert.Nil(t, actualEmplist)
	assert.Error(t, err)
}
func TestCreateEmployees_Success(t *testing.T) {

	employee := models.Employee{
		ID:       3,
		Name:     "Raj",
		Mobile:   "7854589535",
		Address:  "UP",
		Salary:   36566,
		Age:      21,
		Username: "raj",
		Password: "1234",
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)

	mockworker.EXPECT().CreateEmployees(&employee).Times(1).Return(nil)

	err := empService.CreateEmployees(&employee)
	assert.Nil(t, err)
}
func TestCreateEmployees_ErrorWhileCreating_NewEmployee(t *testing.T) {

	employee := models.Employee{
		ID:       3,
		Name:     "Raj",
		Mobile:   "7854589535",
		Address:  "UP",
		Salary:   36566,
		Age:      21,
		Username: "raj",
		Password: "1234",
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)

	mockworker.EXPECT().CreateEmployees(&employee).Times(1).Return(errors.New(""))

	err := empService.CreateEmployees(&employee)
	assert.Error(t, err)
}
func TestGetEmployessById_Success(t *testing.T) {

	var employee *models.Employee = &models.Employee{

		ID:       3,
		Name:     "Raj",
		Mobile:   "7854589535",
		Address:  "UP",
		Salary:   36566,
		Age:      21,
		Username: "raj",
		Password: "1234",
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)
	var id string = strconv.FormatUint(uint64(employee.ID), 10)
	mockworker.EXPECT().GetEmployeeById(id).Times(1).Return(employee, nil)

	actualEmp, err := empService.GetEmployeeById(id)
	assert.Equal(t, employee, actualEmp)
	assert.Nil(t, err)
}
func TestGetEmployessById_ErrorWhileGetting_Employee(t *testing.T) {
	var employee *models.Employee = &models.Employee{

		ID:       3,
		Name:     "Raj",
		Mobile:   "7854589535",
		Address:  "UP",
		Salary:   36566,
		Age:      21,
		Username: "raj",
		Password: "1234",
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)
	var id string = strconv.FormatUint(uint64(employee.ID), 10)
	mockworker.EXPECT().GetEmployeeById(id).Times(1).Return(nil, errors.New(""))

	actualEmp, err := empService.GetEmployeeById(id)
	assert.Nil(t, actualEmp)
	assert.Error(t, err)
}
func TestUpdateEmployee_Success(t *testing.T) {

	var employee *models.Employee = &models.Employee{

		ID:       3,
		Name:     "Raj",
		Mobile:   "7854589535",
		Address:  "UP",
		Salary:   36566,
		Age:      21,
		Username: "raj",
		Password: "1234",
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)
	var id string = strconv.FormatUint(uint64(employee.ID), 10)

	mockworker.EXPECT().GetEmployeeById(id).Times(1).Return(employee, nil)
	mockworker.EXPECT().UpdateEmployee(employee, id).Times(1).Return(nil)

	err := empService.UpdateEmployee(employee, id)
	assert.Nil(t, err)
}
func TestUpdateEmployee_ErrorWhileGetting_Employee(t *testing.T) {

	var employee *models.Employee = &models.Employee{

		ID:       3,
		Name:     "Raj",
		Mobile:   "7854589535",
		Address:  "UP",
		Salary:   36566,
		Age:      21,
		Username: "raj",
		Password: "1234",
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)
	var id string = strconv.FormatUint(uint64(employee.ID), 10)

	mockworker.EXPECT().GetEmployeeById(id).Times(1).Return(employee, errors.New(""))
	mockworker.EXPECT().UpdateEmployee(employee, id).Times(0).Return(nil)

	err := empService.UpdateEmployee(employee, id)
	assert.Error(t, err)
}
func TestUpdateEmployee_ErrorWhileUpdatingEmployee(t *testing.T) {

	var employee *models.Employee = &models.Employee{

		ID:       3,
		Name:     "Raj",
		Mobile:   "7854589535",
		Address:  "UP",
		Salary:   36566,
		Age:      21,
		Username: "raj",
		Password: "1234",
	}

	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockworker := NewMockEmployeeDao(cntrl)
	empService := NewEmployeeServiceImpl(mockworker)
	var id string = strconv.FormatUint(uint64(employee.ID), 10)

	mockworker.EXPECT().GetEmployeeById(id).Times(1).Return(employee, nil)
	mockworker.EXPECT().UpdateEmployee(employee, id).Times(1).Return(errors.New(""))

	err := empService.UpdateEmployee(employee, id)
	assert.Error(t, err)
}
