package adapters_test

import (
	"net/http"
	"testing"

	"github.com/andreascandle/FlexiResponseGo/adapters"
	"github.com/andreascandle/FlexiResponseGo/core"
	"github.com/andreascandle/FlexiResponseGo/tests"
	"github.com/stretchr/testify/assert"
)

func TestHTTPSuccessResponse(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adapters.HTTPSuccessResponse(w, r, "Success message", map[string]string{"key": "value"})
	})
	rec := tests.PerformRequest(handler, "GET", "/", nil)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp core.StandardResponse
	err := tests.ParseJSON(rec, &resp)
	assert.NoError(t, err)

	// Assert response fields
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, "Success message", resp.Message)

	// Assert Data field as map[string]interface{}
	expectedData := map[string]interface{}{"key": "value"}
	assert.Equal(t, expectedData, resp.Data)
}

func TestHTTPErrorResponse(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adapters.HTTPErrorResponse(w, r, http.StatusBadRequest, "Error message", "Detail")
	})
	rec := tests.PerformRequest(handler, "POST", "/error", nil)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp core.StandardResponse
	err := tests.ParseJSON(rec, &resp)
	assert.NoError(t, err)
	assert.Equal(t, "error", resp.Status)
	assert.Equal(t, "Error message", resp.Message)
	assert.Equal(t, "Detail", resp.Error)
}
