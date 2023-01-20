package service

import (
	"errors"
	"rest/gin/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeService_GetAllEmployees_WithNoError(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockWorker := NewMockEmpDao(cntrl)

	svc := NewEmployeeService(mockWorker)

	mockWorker.EXPECT().GetAllEmployees(gomock.Any()).Return(nil)
	var employee []models.Employee
	err := svc.GetAllEmployees(&employee)

	assert.NoError(t, err)
}

func TestEmployeeService_GetEmployeeById_WithError(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockWorker := NewMockEmpDao(cntrl)

	svc := NewEmployeeService(mockWorker)

	expectedError := errors.New("record not found")

	mockWorker.EXPECT().GetEmployeeById(gomock.Any(), gomock.Any()).Return(errors.New("record not found"))
	var employee models.Employee

	resulterr := svc.GetEmployeeById(&employee, "34")
	assert.ErrorIs(t, resulterr, expectedError)

}
