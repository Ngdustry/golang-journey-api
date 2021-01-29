package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	testUtil "golang-journey-api/pkg/utils/testing"

	"github.com/stretchr/testify/assert"
)

var (
	env testUtil.TestEnvironment
)

type statusResponse struct {
	Status string `json:"status"`
}

func TestMain(m *testing.M) {
	env = testUtil.InitTestEnvironment()

	code := m.Run()

	os.Exit(code)
}

func TestStatus(t *testing.T) {
	env.Router.HandleFunc("/status", getStatus).Methods("GET")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/status", nil)
	env.Router.ServeHTTP(res, req)

	var result statusResponse
	json.NewDecoder(res.Body).Decode(&result)

	assert.NotNil(t, result.Status)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", result.Status)

	testUtil.ResetTestRouter(&env)
}
