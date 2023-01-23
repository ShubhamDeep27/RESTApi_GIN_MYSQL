package Controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"

	"rest/gin/models"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetEmployess(t *testing.T) {

	t.Run("TestGetEmployess_Success", func(t *testing.T) {

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
	})

	t.Run("TestGetEmployess_InternalServerError", func(t *testing.T) {
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

	})

}

func TestCreateEmployess(t *testing.T) {

	t.Run("TestCreateEmployess_Success", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		mockWorker.EXPECT().CreateEmployees(&employee).Return(nil)

		server := gin.Default()
		data, err := json.Marshal(employee)
		require.NoError(t, err)
		server.POST("Employee", empController.CreateEmployees)
		recorder := httptest.NewRecorder()

		request := httptest.NewRequest(http.MethodPost, "/Employee", bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)

	})
	t.Run("TestCreateEmployess_InternalServerError", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		mockWorker.EXPECT().CreateEmployees(gomock.Any()).Return(errors.New(""))

		server := gin.Default()
		data, err := json.Marshal(employee)
		require.NoError(t, err)
		server.POST("Employee", empController.CreateEmployees)
		recorder := httptest.NewRecorder()

		request := httptest.NewRequest(http.MethodPost, "/Employee", bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("TestCreateEmployess_BadRequest", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		mockWorker.EXPECT().CreateEmployees(&employee).Times(0).Return(nil)

		server := gin.Default()

		server.POST("Employee", empController.CreateEmployees)
		recorder := httptest.NewRecorder()

		request := httptest.NewRequest(http.MethodPost, "/Employee", nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetEmployeeById(t *testing.T) {

	t.Run("TestGetEmployeeById_Success", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		var id string = strconv.FormatUint(uint64(employee.ID), 10)
		mockWorker.EXPECT().GetEmployeeById(id).Return(&employee, nil)

		server := gin.Default()

		server.GET("Employee/:id", empController.GetEmployeeById)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/Employee/%s", id)
		request := httptest.NewRequest(http.MethodGet, url, nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
		requireBodyMatchEmployee(t, recorder.Body, employee)
	})
	t.Run("TestGetEmployeeById_InternalServerError", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		var id string = strconv.FormatUint(uint64(employee.ID), 10)

		mockWorker.EXPECT().GetEmployeeById(id).Return(&employee, errors.New(""))

		server := gin.Default()

		server.GET("Employee/:id", empController.GetEmployeeById)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/Employee/%s", id)
		request := httptest.NewRequest(http.MethodGet, url, nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)

	})
}

func TestUpdateEmployee(t *testing.T) {

	t.Run("TestUpdateEmployee_Success", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		var id string = strconv.FormatUint(uint64(employee.ID), 10)
		mockWorker.EXPECT().UpdateEmployee(&employee, id).Times(1).Return(nil)

		server := gin.Default()
		data, err := json.Marshal(employee)
		require.NoError(t, err)

		server.PUT("Employee/:id", empController.UpdateEmployee)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/Employee/%s", id)
		request := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
	t.Run("TestUpdateEmployess_InternalServerError", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		var id string = strconv.FormatUint(uint64(employee.ID), 10)
		mockWorker.EXPECT().UpdateEmployee(gomock.Any(), gomock.Any()).Return(errors.New(""))

		server := gin.Default()
		data, err := json.Marshal(employee)
		require.NoError(t, err)
		server.PUT("Employee/:id", empController.UpdateEmployee)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/Employee/%s", id)
		request := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)

	})
	t.Run("TestUpdateEmployess_BadRequest", func(t *testing.T) {
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

		mockWorker := NewMockEmployeeService(cntrl)
		empController := NewEmployeeController(mockWorker)
		var id string = strconv.FormatUint(uint64(employee.ID), 10)
		mockWorker.EXPECT().UpdateEmployee(&employee, id).Times(0).Return(nil)

		server := gin.Default()

		server.PUT("Employee/:id", empController.CreateEmployees)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/Employee/%s", id)
		request := httptest.NewRequest(http.MethodPut, url, nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusBadRequest, recorder.Code)

	})
}
