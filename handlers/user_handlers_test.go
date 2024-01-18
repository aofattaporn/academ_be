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

type MockUserService struct {
	CreateUserFunc func(user *models.User) error
}

func (m *MockUserService) CreateUser(user *models.User) error {
	return m.CreateUserFunc(user)
}

func TestCreateResource(t *testing.T) {
	r := gin.Default()
	r.POST("/api/resource", handlers.CreateResource)

	user := models.User{}

	reqBody, err := json.Marshal(user)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/resource", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	mockUserService := &MockUserService{
		CreateUserFunc: func(user *models.User) error {
			return nil
		},
	}

	handlers.SetUserService(mockUserService)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
