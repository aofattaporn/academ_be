package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTester(c *gin.Context) {

	handleSuccess(c, http.StatusOK, SUCCESS, "Get Test Api Success", nil)
}
