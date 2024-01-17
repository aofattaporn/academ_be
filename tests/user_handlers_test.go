package handlers

import (
	"academ_be/handlers"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateResource_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Create a Gin router with the CreateResource handler
	router := gin.Default()
	router.POST("/create", handlers.CreateResource)

	// Mock a request with a valid JSON payload
	validJSON := `{"email": "test@example.com", "password": "testpass"}`
	req, err := http.NewRequest("POST", "/create", bytes.NewBufferString(validJSON))
	assert.NoError(t, err)

	// Set the request header
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	res := httptest.NewRecorder()

	// Serve the request to the router
	router.ServeHTTP(res, req)

	// Assert the HTTP status code and response body
	assert.Equal(t, http.StatusCreated, res.Code)

	expectedResponse := `{"status":201,"message":"SUCCESS","description":"USER_SIGNUP_SUCCESS"}`
	assert.Equal(t, expectedResponse, res.Body.String())
}
