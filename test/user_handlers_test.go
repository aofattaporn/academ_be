package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"academ_be/handlers"
	"academ_be/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockUserService is a mock implementation of the UserService interface for testing.
type MockUserService struct {
	CreateUserFunc func(user *models.User) error
}

func (m *MockUserService) CreateUser(user *models.User) error {
	return m.CreateUserFunc(user)
}

func TestCreateResource(t *testing.T) {
	// Create a test router with the handler
	r := gin.Default()
	r.POST("/api/resource", handlers.CreateResource) // Use the original package name

	// Prepare a sample user for the test
	user := models.User{
		// populate fields as needed
	}

	// Prepare request body
	reqBody, err := json.Marshal(user)
	assert.NoError(t, err)

	// Set up the request
	req, err := http.NewRequest("POST", "/api/resource", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create a mock UserService for testing
	mockUserService := &MockUserService{
		CreateUserFunc: func(user *models.User) error {
			// Mock implementation for CreateUser function, e.g., return nil for success
			return nil
		},
	}

	// Set the mock UserService for testing
	handlers.SetUserService(mockUserService) // Use the alias 'handlers'

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verify the response
	assert.Equal(t, http.StatusCreated, w.Code)
}
