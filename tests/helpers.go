package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// PerformRequest simulates an HTTP request for testing purposes.
func PerformRequest(handler http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	var jsonBody []byte
	if body != nil {
		jsonBody, _ = json.Marshal(body)
	}
	req := httptest.NewRequest(method, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

// ParseJSON parses a JSON response body into the specified struct.
func ParseJSON(rec *httptest.ResponseRecorder, target interface{}) error {
	return json.NewDecoder(rec.Body).Decode(target)
}
