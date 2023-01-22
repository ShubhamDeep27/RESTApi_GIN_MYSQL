package Controller

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"rest/gin/models"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetEmployess_Success(t *testing.T) {
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
	empController := NewEmployeeController(mockWorker)

	mockWorker.EXPECT().GetAllEmployees().Return(employeeList, nil).Times(1)

	server := gin.Default()
	server.GET("Employee", empController.GetEmployess)
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/Employee", nil)

	server.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchEmployees(t, recorder.Body, employeeList)

}
func TestGetEmployess_InternalServerError(t *testing.T) {
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

	mockWorker := NewMockEmployeeService(cntrl)
	empController := NewEmployeeController(mockWorker)
	mockWorker.EXPECT().GetAllEmployees().Return(employeeList, errors.New("")).Times(1)

	server := gin.Default()
	server.GET("Employee", empController.GetEmployess)
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/Employee", nil)

	server.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusInternalServerError, recorder.Code)

}
