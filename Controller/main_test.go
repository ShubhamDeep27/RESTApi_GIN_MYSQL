package Controller

import (
	"bytes"
	"encoding/json"
	"io"

	"os"
	"rest/gin/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// func requireBodyMatchEmployee(t *testing.T, body *bytes.Buffer, emp models.Employee) {
// 	data, err := ioutil.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotemp models.Employee
// 	err = json.Unmarshal(data, &gotemp)
// 	require.NoError(t, err)
// 	require.Equal(t, emp, gotemp)
// }

func requireBodyMatchEmployees(t *testing.T, body *bytes.Buffer, emp []*models.Employee) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotemp []*models.Employee

	err = json.Unmarshal(data, &gotemp)
	require.NoError(t, err)
	require.Equal(t, emp, gotemp)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
