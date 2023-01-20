package Controller

import (
	"rest/gin/models"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetEmployess(t *testing.T) {
	employeeList := []*models.Employee{
		{
			ID:       2,
			Name:     "Ram",
			Mobile:   "78589535",
			Address:  "Delhi",
			Salary:   34566,
			Age:      23,
			Username: "ram",
			Password: "123",
		},
		{
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

	mockWorker := NewMockEmployeeService(cntrl)

	svc := NewEmployeeController(mockWorker)

	mockWorker.EXPECT().GetAllEmployees(gomock.Any()).Return(nil)
	err := svc.GetAllEmployees(employeeList)

	assert.NoError(t, err)
}
