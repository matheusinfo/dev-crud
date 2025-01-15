package controllers

import (
	"bytes"
	"dev-crud/src/database"
	"dev-crud/src/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateDev(t *testing.T) {
	router := gin.Default()
	router.POST("/dev", CreateDev)

	dev := models.Dev{
		Username: "devUser",
		Email:    "dev@example.com",
	}

	jsonValue, _ := json.Marshal(dev)
	req, _ := http.NewRequest("POST", "/dev", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Dev
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, dev.Username, response.Username)
	assert.Equal(t, dev.Email, response.Email)
}

func TestGetDevs(t *testing.T) {
	router := gin.Default()
	router.GET("/dev", GetDevs)

	dev := models.Dev{
		Username: "devUser",
		Email:    "dev@example.com",
	}
	database.Instance.Create(&dev)

	req, _ := http.NewRequest("GET", "/dev", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Dev
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Greater(t, len(response), 0)
}

func TestGetDevById(t *testing.T) {
	router := gin.Default()
	router.GET("/dev/:id", GetDevById)

	dev := models.Dev{
		Username: "devUser",
		Email:    "dev@example.com",
	}
	database.Instance.Create(&dev)

	req, _ := http.NewRequest("GET", "/dev/1234", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Dev
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, dev.ID, response.ID)
	assert.Equal(t, dev.Username, response.Username)
	assert.Equal(t, dev.Email, response.Email)
}

func TestRemoveDev(t *testing.T) {
	router := gin.Default()
	router.DELETE("/dev/:id", RemoveDev)

	dev := models.Dev{
		Username: "devUser",
		Email:    "dev@example.com",
	}
	database.Instance.Create(&dev)

	req, _ := http.NewRequest("DELETE", "/dev/1234", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	
	var deletedDev models.Dev
	result := database.Instance.First(&deletedDev, "id = ?", dev.ID)
	assert.Error(t, result.Error)
}
