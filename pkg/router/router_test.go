package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type statusResponse struct {
	status string `json:"status"`
}

func TestStatus(t *testing.T) {
	r := New()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	r.ServeHTTP(res, req)

	var result statusResponse
	json.NewDecoder(res.Body).Decode(&result)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}
