package Controller

import (
	"bytes"
	"encoding/json"
	"os"
	"rest/gin/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func requireBodyMatchEmployee(t *testing.T, body *bytes.Buffer, emp models.Employee) {

	var gotemp models.Employee
	err := json.Unmarshal(body.Bytes(), &gotemp)
	require.NoError(t, err)
	require.Equal(t, emp, gotemp)
}

func requireBodyMatchEmployees(t *testing.T, body *bytes.Buffer, emp []*models.Employee) {

	var gotemp []*models.Employee

	err := json.Unmarshal(body.Bytes(), &gotemp)
	require.NoError(t, err)
	require.Equal(t, emp, gotemp)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
