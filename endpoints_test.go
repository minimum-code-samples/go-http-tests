// +build alltests rest

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	assert.NoError(t, err)
	resp := httptest.NewRecorder()
	Router().ServeHTTP(resp, req)

	// Check the response code
	assert.Equal(t, 200, resp.Code, "200 OK response")
	// Check the response body
	assert.Equal(t, []byte("Hello world"), resp.Body.Bytes())
}
